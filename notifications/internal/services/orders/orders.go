package orders

import "github.com/Shopify/sarama"

type IOrdersService interface {
	NotificateOrderCreated(message *sarama.ConsumerMessage) error
}

type Order struct {
}

var _ IOrdersService = (*Order)(nil)

func New() IOrdersService {
	return &Order{}
}
