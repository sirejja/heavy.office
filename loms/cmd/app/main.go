package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"route256/libs/interceptors"
	kafka "route256/libs/kafka/producer"
	"route256/libs/logger"
	"route256/libs/tracer"
	"route256/libs/transactor"
	v1 "route256/loms/internal/api/v1"
	"route256/loms/internal/config"
	"route256/loms/internal/cronjob"
	"route256/loms/internal/kafka/outbox_producer"
	"route256/loms/internal/repositories/order_repo"
	"route256/loms/internal/repositories/outbox_repo"
	"route256/loms/internal/repositories/warehouse_orders_repo"
	"route256/loms/internal/repositories/warehouse_repo"
	cancel_orders_cron "route256/loms/internal/services/cron/cancel_orders"
	"route256/loms/internal/services/cron/outbox"
	"route256/loms/internal/services/orders"
	"route256/loms/internal/services/warehouse"
	desc "route256/loms/pkg/v1/api"
	"time"

	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpcPrometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	cfg := config.New()
	if err := cfg.Init(); err != nil {
		panic("config init")
	}

	logger.Init(cfg.Env == config.EnvDev)
	tracer.InitGlobalTracer(cfg.ServiceName)

	postgresConfig, err := pgxpool.ParseConfig(cfg.Storage.PostgresDSN)
	if err != nil {
		logger.Fatal("can not parse postgres DSN", zap.Error(err))
	}
	postgresConfig.MaxConnIdleTime = time.Minute
	postgresConfig.MaxConnLifetime = time.Hour
	postgresConfig.MinConns = 1
	postgresConfig.MaxConns = 2

	pool, err := pgxpool.ConnectConfig(ctx, postgresConfig)
	if err != nil {
		logger.Fatal("Unable to connect to database", zap.Error(err))
	}
	defer pool.Close()
	transactionManager := transactor.New(pool)

	warehouseRepo := warehouse_repo.New(transactionManager)
	ordersRepo := order_repo.New(transactionManager)
	warehouseOrdersRepo := warehouse_orders_repo.New(transactionManager)
	outboxRepo := outbox_repo.New(transactionManager)

	producer, err := kafka.NewSyncProducer(cfg.Kafka.Brokers)
	if err != nil {
		logger.Fatal("Unable to connect to kafka", zap.Error(err))
	}
	defer producer.Close()

	ordersProcessor := orders.New(
		ordersRepo,
		warehouseRepo,
		warehouseOrdersRepo,
		outboxRepo,
		transactionManager,
		cfg,
	)

	warehouseProcessor := warehouse.New(warehouseRepo, transactionManager)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.WebPort))
	if err != nil {
		logger.Fatal(fmt.Sprintf("failed to connect to listen %v", cfg.WebPort), zap.Error(err))
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
	desc.RegisterLomsServer(server, v1.New(warehouseProcessor, ordersProcessor))

	// prometheus
	grpcPrometheus.Register(server)
	http.Handle("/metrics", promhttp.Handler())
	go func() {
		if err := http.ListenAndServe(fmt.Sprintf(":%d", cfg.MetricsPort), nil); err != nil {
			logger.Fatal("failed to serve", zap.Error(err))
		}
	}()

	// cronjob
	logger.Info("CronJob starting...")
	cancelOrdersJob := cancel_orders_cron.New(ctx, ordersProcessor, ordersRepo, cfg.CancelOrdersCronPeriod)
	outboxCron := outbox.New(ctx, outbox_producer.New(producer), outboxRepo, cfg.OutboxCronPeriod)

	cronJob := cronjob.New(cancelOrdersJob, outboxCron)
	cronJob.Start()
	logger.Info("CronJob started")

	logger.Info("grpc server listening at", zap.Int("port", cfg.WebPort))
	if err = server.Serve(lis); err != nil {
		logger.Fatal("failed to serve", zap.Error(err))
	}
}
