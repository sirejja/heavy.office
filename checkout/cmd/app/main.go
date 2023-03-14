package main

import (
	"context"
	"log"
	"net"
	v1 "route256/checkout/internal/api/v1"
	"route256/checkout/internal/clients/loms"
	"route256/checkout/internal/clients/products"
	"route256/checkout/internal/config"
	"route256/checkout/internal/repositories/carts_products_repo"
	"route256/checkout/internal/repositories/carts_repo"
	"route256/checkout/internal/services/cart"
	desc "route256/checkout/pkg/v1/api"
	"route256/libs/interceptors"
	"route256/libs/transactor"
	"time"

	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

const port = ":8080"

func main() {
	cfg := config.New()
	if err := cfg.Init(); err != nil {
		log.Fatal("config init", err)
	}

	postgresConfig, err := pgxpool.ParseConfig(cfg.Storage.PostgresDSN)
	if err != nil {
		log.Fatal("can not parse postgres DSN", err)
	}
	postgresConfig.MaxConnIdleTime = time.Minute
	postgresConfig.MaxConnLifetime = time.Hour
	postgresConfig.MinConns = 1
	postgresConfig.MaxConns = 2

	pool, err := pgxpool.ConnectConfig(context.Background(), postgresConfig)
	if err != nil {
		log.Fatal("Unable to connect to database", err)
	}
	defer pool.Close()
	transactionManager := transactor.New(pool)

	cartsRepo := carts_repo.New(transactionManager)
	cartsProductsRepo := carts_products_repo.New(transactionManager)

	lomsConn, err := grpc.Dial(cfg.Services.Loms.URL, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("failed to connect to lomsClient", err)
	}
	defer lomsConn.Close()

	lomsClient, err := loms.New(lomsConn)
	if err != nil {
		log.Fatal("failed to create lomsClient", err)
	}

	productsServiceConn, err := grpc.Dial(cfg.Services.Products.URL, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("failed to connect to productsClient", err)
	}
	defer productsServiceConn.Close()

	productsClient, err := products.New(productsServiceConn, cfg.Services.Products.Token)
	if err != nil {
		log.Fatal("failed to create productsClient", err)
	}

	cartProcessor := cart.New(lomsClient, productsClient, cartsRepo, cartsProductsRepo, transactionManager)

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to connect to listen %v: %v", port, err)
	}

	server := grpc.NewServer(
		grpc.UnaryInterceptor(
			grpcMiddleware.ChainUnaryServer(
				interceptors.LoggingInterceptor,
			),
		),
	)
	reflection.Register(server)
	desc.RegisterCheckoutServer(server, v1.New(cartProcessor))

	log.Println("grpc server listening at", port)
	if err = server.Serve(lis); err != nil {
		log.Fatal("failed to serve", err)
	}
}
