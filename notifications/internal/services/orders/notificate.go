package orders

import (
	"log"

	"github.com/Shopify/sarama"
)

func (o *Order) NotificateOrderCreated(message *sarama.ConsumerMessage) error {
	log.Printf("Message claimed: value = %s, timestamp = %v, topic = %s", string(message.Value), message.Timestamp, message.Topic)
	return nil
}
