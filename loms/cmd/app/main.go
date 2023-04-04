package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"route256/libs/interceptors"
	kafka "route256/libs/kafka/producer"
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
	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

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

	pool, err := pgxpool.ConnectConfig(ctx, postgresConfig)
	if err != nil {
		log.Fatal("Unable to connect to database", err)
	}
	defer pool.Close()
	transactionManager := transactor.New(pool)

	warehouseRepo := warehouse_repo.New(transactionManager)
	ordersRepo := order_repo.New(transactionManager)
	warehouseOrdersRepo := warehouse_orders_repo.New(transactionManager)
	outboxRepo := outbox_repo.New(transactionManager)

	producer, err := kafka.NewSyncProducer(cfg.Kafka.Brokers)
	if err != nil {
		log.Fatal("Unable to connect to kafka", err)
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

	lis, err := net.Listen("tcp", cfg.Web.Port)
	if err != nil {
		log.Fatalf("failed to connect to listen %v: %v", cfg.Web.Port, err)
	}

	server := grpc.NewServer(
		grpc.UnaryInterceptor(
			grpcMiddleware.ChainUnaryServer(
				interceptors.LoggingInterceptor,
			),
		),
	)
	reflection.Register(server)
	desc.RegisterLomsServer(server, v1.New(warehouseProcessor, ordersProcessor))

	// cronjob
	log.Println("CronJob starting...")
	cancelOrdersJob := cancel_orders_cron.New(ctx, ordersProcessor, ordersRepo, cfg.CancelOrdersCronPeriod)
	outboxCron := outbox.New(ctx, outbox_producer.New(producer), outboxRepo, cfg.OutboxCronPeriod)

	cronJob := cronjob.New(cancelOrdersJob, outboxCron)
	cronJob.Start()
	log.Println("CronJob started")

	log.Println("grpc server listening at", cfg.Web.Port)
	if err = server.Serve(lis); err != nil {
		log.Fatal("failed to serve", err)
	}
}
