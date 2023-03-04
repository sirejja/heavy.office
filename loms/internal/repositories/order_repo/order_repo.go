package order_repo

import (
	"route256/loms/internal/models"
)

type IOrderRepo interface {
	ListOrder(orderID int64) (*models.Order, error)
	PayedOrder(orderID int64) error
	CancelOrder(orderID int64) error
}

type OrderRepo struct {
}

var _ IOrderRepo = (*OrderRepo)(nil)

func New() OrderRepo {
	return OrderRepo{}
}
