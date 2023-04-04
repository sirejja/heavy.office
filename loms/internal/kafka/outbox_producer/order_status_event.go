package outbox_producer

import (
	"fmt"
	"log"
	"time"

	"github.com/Shopify/sarama"
)

type OrderStatusMsg struct {
	ID     int64
	Status string
}

func (p *Producer) sendOrderOrderStatusEvent(topic string, orderMsg OrderStatusMsg) error {
	op := "Producer.sendOrderOrderStatusEvent"

	msg := &sarama.ProducerMessage{
		Topic:     topic,
		Partition: -1,
		Value:     sarama.StringEncoder(fmt.Sprintf(`{"id":%d, status: %s}`, orderMsg.ID, orderMsg.Status)),
		Key:       sarama.StringEncoder(fmt.Sprint(orderMsg.ID)),
		Timestamp: time.Now(),
		Headers: []sarama.RecordHeader{
			{
				Key:   []byte("status-changed-header"),
				Value: []byte(orderMsg.Status),
			},
		},
	}

	partition, offset, err := p.producer.SendMessage(msg)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	log.Printf("order id: %d, partition: %d, offset: %d", orderMsg.ID, partition, offset)

	return nil
}
