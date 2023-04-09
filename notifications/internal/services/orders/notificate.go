package orders

import (
	"route256/libs/logger"

	"github.com/Shopify/sarama"
	"go.uber.org/zap"
)

func (o *Order) NotificateOrderCreated(message *sarama.ConsumerMessage) error {
	logger.Info("Message claimed order_status changed",
		zap.ByteString("value", message.Value),
		zap.Time("timestamp", message.Timestamp),
		zap.String("topic", message.Topic),
	)
	return nil
}
