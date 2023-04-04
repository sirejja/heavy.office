package kafka

import (
	"route256/libs/kafka/consumer_group"
	"route256/notifications/internal/config"
	"route256/notifications/internal/services/orders"

	"github.com/Shopify/sarama"
)

func NewConsumerGroup(cfg *config.ConfigStruct, ordersService orders.IOrdersService) (consumer_group.Consumer, *sarama.Config) {
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
	topicHandlerMapping := map[string]func(msg *sarama.ConsumerMessage) error{
		cfg.Kafka.Topics.OrderStatus.Topic: ordersService.NotificateOrderCreated,
	}

	return consumer_group.New(topicHandlerMapping), kafkaCfg
}
