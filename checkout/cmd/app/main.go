package main

import (
	"log"
	"net"
	v1 "route256/checkout/internal/api/v1"
	"route256/checkout/internal/clients/loms"
	"route256/checkout/internal/clients/products"
	"route256/checkout/internal/config"
	"route256/checkout/internal/interceptors"
	"route256/checkout/internal/services/cart"
	desc "route256/checkout/pkg/grpc/server"

	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const port = ":8080"

func main() {
	cfg := config.New()
	if err := cfg.Init(); err != nil {
		log.Fatal("config init", err)
	}

	lomsClient, err := loms.New(cfg.Services.Loms.URL)
	if err != nil {
		log.Fatal("failed to connect to lomsClient", err)
	}
	productsClient, err := products.New(cfg.Services.Products.URL, cfg.Services.Products.Token)
	if err != nil {
		log.Fatal("failed to connect to productsClient", err)
	}

	cartProcessor := cart.New(lomsClient, productsClient)

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
	desc.RegisterCheckoutServer(server, v1.New(*cartProcessor))

	log.Println("grpc server listening at", port)
	if err = server.Serve(lis); err != nil {
		log.Fatal("failed to serve", err)
	}
}
