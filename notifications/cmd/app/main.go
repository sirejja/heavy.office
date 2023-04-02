package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"route256/notifications/internal/config"
	"route256/notifications/internal/kafka/consumer_group"
	"route256/notifications/internal/services/orders"
	"sync"
	"syscall"

	"github.com/Shopify/sarama"
)

func main() {
	keepRunning := true
	log.Println("Starting a new Sarama consumer")

	cfg := config.New()
	if err := cfg.Init(); err != nil {
		log.Fatal("config init", err)
	}
	/**
	 * Construct a new Sarama configuration.
	 * The Kafka cluster version has to be defined before the consumer/producer is initialized.
	 */
	kafkaCfg := sarama.NewConfig()
	kafkaCfg.Version = sarama.MaxVersion
	kafkaCfg.Consumer.Offsets.Initial = sarama.OffsetNewest

	switch cfg.Kafka.BalanceStrategy {
	case "sticky":
		kafkaCfg.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{sarama.BalanceStrategySticky}
	case "roundrobin":
		kafkaCfg.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{sarama.BalanceStrategyRoundRobin}
	case "range":
		kafkaCfg.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{sarama.BalanceStrategyRange}
	}

	consumer := consumer_group.New(orders.New())

	ctx, cancel := context.WithCancel(context.Background())
	client, err := sarama.NewConsumerGroup(cfg.Kafka.Brokers, cfg.Kafka.GroupName, kafkaCfg)
	if err != nil {
		log.Panicf("Error creating consumer group client: %v", err)
	}

	consumptionIsPaused := false

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			// `Consume` should be called inside an infinite loop, when a
			// server-side rebalance happens, the consumer session will need to be
			// recreated to get the new claims
			if err := client.Consume(ctx, []string{cfg.Kafka.Topics.OrderStatus.Topic}, &consumer); err != nil {
				log.Panicf("Error from consumer: %v", err)
			}
			// check if context was cancelled, signaling that the consumer should stop
			if ctx.Err() != nil {
				return
			}
		}
	}()

	<-consumer.Ready() // Await till the consumer has been set up
	log.Println("Sarama consumer up and running!...")

	sigusr1 := make(chan os.Signal, 1)
	signal.Notify(sigusr1, syscall.SIGUSR1)

	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)

	for keepRunning {
		select {
		case <-ctx.Done():
			log.Println("terminating: context cancelled")
			keepRunning = false
		case <-sigterm:
			log.Println("terminating: via signal")
			keepRunning = false
		case <-sigusr1:
			toggleConsumptionFlow(client, &consumptionIsPaused)
		}
	}

	cancel()
	wg.Wait()
	if err = client.Close(); err != nil {
		log.Panicf("Error closing client: %v", err)
	}
}

func toggleConsumptionFlow(client sarama.ConsumerGroup, isPaused *bool) {
	if *isPaused {
		client.ResumeAll()
		log.Println("Resuming consumption")
	} else {
		client.PauseAll()
		log.Println("Pausing consumption")
	}

	*isPaused = !*isPaused
}
