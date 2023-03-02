package main

import (
	"log"
	"net"
	v1 "route256/loms/internal/api/v1"
	"route256/loms/internal/config"
	"route256/loms/internal/interceptors"
	"route256/loms/internal/repo/order_repo"
	"route256/loms/internal/repo/warehouse_repo"
	"route256/loms/internal/services/orders"
	"route256/loms/internal/services/warehouse"
	desc "route256/loms/pkg/grpc/server"

	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const port = ":8081"

func main() {
	cfg := config.New()
	if err := cfg.Init(); err != nil {
		log.Fatal("config init", err)
	}

	warehouseRepo := warehouse_repo.New()
	ordersRepo := order_repo.New()

	ordersProcessor := orders.New(ordersRepo, warehouseRepo)
	warehouseProcessor := warehouse.New(warehouseRepo)

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
	desc.RegisterLomsServer(server, v1.New(*warehouseProcessor, ordersProcessor))

	log.Println("grpc server listening at", port)
	if err = server.Serve(lis); err != nil {
		log.Fatal("failed to serve", err)
	}
}
