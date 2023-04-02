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
	SendOrderOrderStatusEvent(id int64, status models.OrderStatus) error
}

//type Handler func(id string)

func NewOrderSender(producer sarama.SyncProducer, topic string) IOrderSender {
	s := &orderSender{
		producer: producer,
		topic:    topic,
	}
	return s
}

func (s *orderSender) SendOrderOrderStatusEvent(id int64, status models.OrderStatus) error {
	op := "orderSender.SendOrderOrderStatusEvent"

	msg := &sarama.ProducerMessage{
		Topic:     s.topic,
		Partition: -1,
		Value:     sarama.StringEncoder(fmt.Sprintf(`{"id":%d, status: %s}`, id, status.ToString())),
		Key:       sarama.StringEncoder(fmt.Sprint(id)),
		Timestamp: time.Now(),
		Headers: []sarama.RecordHeader{
			{
				Key:   []byte("status-changed-header"),
				Value: []byte(status.ToString()),
			},
		},
	}

	partition, offset, err := s.producer.SendMessage(msg)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	log.Printf("order id: %d, partition: %d, offset: %d", id, partition, offset)
	return nil
}
