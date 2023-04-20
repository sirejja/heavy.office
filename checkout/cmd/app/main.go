package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	v1 "route256/checkout/internal/api/v1"
	"route256/checkout/internal/clients/grpc/loms"
	"route256/checkout/internal/clients/grpc/products"
	"route256/checkout/internal/config"
	"route256/checkout/internal/repositories/carts_products_repo"
	"route256/checkout/internal/repositories/carts_repo"
	"route256/checkout/internal/services/cart"
	desc "route256/checkout/pkg/v1/api"
	"route256/libs/interceptors"
	"route256/libs/logger"
	"route256/libs/tracer"
	"route256/libs/transactor"
	"time"

	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpcPrometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/opentracing/opentracing-go"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
	"golang.org/x/time/rate"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"route256/libs/cache/inmemory"
)

func main() {

	cfg := config.New()
	if err := cfg.Init(); err != nil {
		panic(fmt.Sprintf("config init %v", err))
	}

	logger.Init(cfg.Env == config.EnvDev)
	tracer.InitGlobalTracer(cfg.ServiceName)

	// db init
	postgresConfig, err := pgxpool.ParseConfig(cfg.Storage.PostgresDSN)
	if err != nil {
		logger.Fatal("can not parse postgres DSN", zap.Error(err))
	}
	postgresConfig.MaxConnIdleTime = time.Minute
	postgresConfig.MaxConnLifetime = time.Hour
	postgresConfig.MinConns = 1
	postgresConfig.MaxConns = 2

	pool, err := pgxpool.ConnectConfig(context.Background(), postgresConfig)
	if err != nil {
		logger.Fatal("Unable to connect to database", zap.Error(err))
	}
	defer pool.Close()
	transactionManager := transactor.New(pool)

	cartsRepo := carts_repo.New(transactionManager)
	cartsProductsRepo := carts_products_repo.New(transactionManager)

	// clients init
	lomsConn, err := grpc.Dial(
		cfg.Services.Loms.URL,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(otgrpc.OpenTracingClientInterceptor(opentracing.GlobalTracer())),
	)
	if err != nil {
		logger.Fatal("failed to connect to lomsClient", zap.Error(err))
	}
	defer lomsConn.Close()

	lomsClient, err := loms.New(lomsConn)
	if err != nil {
		logger.Fatal("failed to create lomsClient", zap.Error(err))
	}

	productsServiceConn, err := grpc.Dial(
		cfg.Services.Products.URL,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(otgrpc.OpenTracingClientInterceptor(opentracing.GlobalTracer())),
	)
	if err != nil {
		logger.Fatal("failed to connect to productsClient", zap.Error(err))
	}
	defer productsServiceConn.Close()

	cache := inmemory.New(30, time.Hour)
	inmemory.InitMetricsCache(cfg.ServiceName)

	productsLimiter := rate.NewLimiter(rate.Every(time.Second/100), 10)
	productsClient, err := products.New(productsServiceConn, cfg.Services.Products.Token, cache)
	if err != nil {
		logger.Fatal("failed to create productsClient", zap.Error(err))
	}

	// business logic init
	cartProcessor := cart.New(lomsClient, productsClient, cartsRepo, cartsProductsRepo, transactionManager, productsLimiter)

	// sever init
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.WebPort))
	if err != nil {
		logger.Fatal(fmt.Sprintf("failed to connect to listen %d", cfg.WebPort), zap.Error(err))
	}

	interceptors.InitMetricsInterceptor(cfg.ServiceName)
	server := grpc.NewServer(
		grpc.UnaryInterceptor(
			grpcMiddleware.ChainUnaryServer(
				interceptors.LoggingInterceptor,
				interceptors.TracingInterceptor,
				interceptors.MetricsInterceptor,
				grpcPrometheus.UnaryServerInterceptor,
			),
		),
	)

	reflection.Register(server)
	desc.RegisterCheckoutServer(server, v1.New(cartProcessor))

	// prometheus
	grpcPrometheus.Register(server)
	http.Handle("/metrics", promhttp.Handler())
	go func() {
		if err = http.ListenAndServe(fmt.Sprintf(":%d", cfg.MetricsPort), nil); err != nil {
			logger.Fatal("failed to serve", zap.Error(err))
		}
	}()

	logger.Info("grpc server listening at", zap.Int("port", cfg.WebPort))
	if err = server.Serve(lis); err != nil {
		logger.Fatal("failed to serve", zap.Error(err))
	}
}
