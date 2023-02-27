package order_repo

import (
	"route256/loms/internal/models"
)

type IOrderRepo interface {
	ListOrder(orderID int64) (*Order, error)
	PayedOrder(orderID int64) error
	CancelOrder(orderID int64) error
}

type Order struct {
	User   uint64
	Status models.OrderStatus
	Items  []models.Item
}

func New() IOrderRepo {
	return &Order{}
}

func (o *Order) ListOrder(orderID int64) (*Order, error) {
	order := Order{User: 111, Status: models.NewOrderStatus, Items: []models.Item{{SKU: 111, Count: 10}}}
	return &order, nil
}

func (o *Order) PayedOrder(orderID int64) error {
	return nil
}

func (o *Order) CancelOrder(orderID int64) error {
	return nil
}
