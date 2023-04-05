package outbox_producer

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/Shopify/sarama"
)

type Producer struct {
	producer sarama.SyncProducer
}

type IProducerHandler interface {
	ResolveProducerHandler(topic string, msg []byte) error
}

func New(producer sarama.SyncProducer) IProducerHandler {
	return &Producer{
		producer: producer,
	}
}

func (p *Producer) ResolveProducerHandler(topic string, msg []byte) error {
	op := "Producer.ResolveProducerHandler"
	var err error

	switch topic {
	case "order_status":
		data := OrderStatusMsg{}
		if err = json.Unmarshal(msg, &data); err != nil {
			return fmt.Errorf("%s: %w", op, err)
		}
		err = p.sendOrderOrderStatusEvent(topic, data)
	default:
		return fmt.Errorf("%s: %w", op, errors.New("Unrecognized task"))
	}

	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
