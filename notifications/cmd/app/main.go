package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"route256/libs/logger"
	"route256/libs/tracer"
	"route256/notifications/internal/config"
	"route256/notifications/internal/kafka"
	"route256/notifications/internal/services/orders"
	"sync"
	"syscall"

	"github.com/Shopify/sarama"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
)

func main() {
	keepRunning := true

	cfg := config.New()
	if err := cfg.Init(); err != nil {
		panic("config init")
	}

	logger.Init(cfg.Env == config.EnvDev)
	tracer.InitGlobalTracer(cfg.ServiceName)

	logger.Info("Starting a new Sarama consumer")
	consumer, kafkaCfg := kafka.NewConsumerGroup(cfg, orders.New())

	ctx, cancel := context.WithCancel(context.Background())
	client, err := sarama.NewConsumerGroup(cfg.Kafka.Brokers, cfg.Kafka.GroupName, kafkaCfg)
	if err != nil {
		logger.Fatal("Error creating consumer group client", zap.Error(err))
	}

	consumptionIsPaused := false

	// prometheus
	http.Handle("/metrics", promhttp.Handler())
	go func() {
		if err := http.ListenAndServe(fmt.Sprintf(":%d", cfg.MetricsPort), nil); err != nil {
			logger.Fatal("failed to serve", zap.Error(err))
		}
	}()

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			if err := client.Consume(ctx, []string{cfg.Kafka.Topics.OrderStatus.Topic}, &consumer); err != nil {
				logger.Fatal("Error from consumer", zap.Error(err))
			}
			if ctx.Err() != nil {
				return
			}
		}
	}()

	<-consumer.Ready() // Await till the consumer has been set up
	logger.Info("Sarama consumer up and running!...")

	sigusr1 := make(chan os.Signal, 1)
	signal.Notify(sigusr1, syscall.SIGUSR1)

	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)

	for keepRunning {
		select {
		case <-ctx.Done():
			logger.Info("terminating: context cancelled")
			keepRunning = false
		case <-sigterm:
			logger.Info("terminating: via signal")
			keepRunning = false
		case <-sigusr1:
			consumer.ToggleConsumptionFlow(client, &consumptionIsPaused)
		}
	}

	cancel()
	wg.Wait()
	if err = client.Close(); err != nil {
		logger.Error("Error closing client: %v", zap.Error(err))
	}
}
