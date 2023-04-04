package order_sender

import (
	"fmt"
	"log"
	"route256/loms/internal/models"
	"time"

	"github.com/Shopify/sarama"
)

type orderSender struct {
	producer sarama.SyncProducer
	topic    string
}

type IOrderSender interface {
	SendOrderOrderStatusEvent(orderMsg OrderStatusMsg) error
}

func NewOrderSender(producer sarama.SyncProducer, topic string) IOrderSender {
	s := &orderSender{
		producer: producer,
		topic:    topic,
	}
	return s
}

type OrderStatusMsg struct {
	ID     int64
	Status models.OrderStatus
}

func (s *orderSender) SendOrderOrderStatusEvent(orderMsg OrderStatusMsg) error {
	op := "orderSender.SendOrderOrderStatusEvent"

	msg := &sarama.ProducerMessage{
		Topic:     s.topic,
		Partition: -1,
		Value:     sarama.StringEncoder(fmt.Sprintf(`{"id":%d, status: %s}`, orderMsg.ID, orderMsg.Status.ToString())),
		Key:       sarama.StringEncoder(fmt.Sprint(orderMsg.ID)),
		Timestamp: time.Now(),
		Headers: []sarama.RecordHeader{
			{
				Key:   []byte("status-changed-header"),
				Value: []byte(orderMsg.Status.ToString()),
			},
		},
	}

	partition, offset, err := s.producer.SendMessage(msg)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	log.Printf("order id: %d, partition: %d, offset: %d", orderMsg.ID, partition, offset)
	return nil
}
