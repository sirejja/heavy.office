package order_repo

import (
	"context"
	"route256/libs/transactor"
	"route256/loms/internal/models"
)

type IOrderRepo interface {
	GetOrderDetails(ctx context.Context, orderID uint64) (*models.OrderDetails, error)
	UpdateOrderStatus(ctx context.Context, orderID int64, status models.OrderStatus) (uint64, error)
	CreateOrder(ctx context.Context, userID int64, status models.OrderStatus) (uint64, error)
	ListOrder(ctx context.Context, orderID uint64) ([]models.ListOrder, error)
	ListOrderStacked(ctx context.Context, orderID uint64) ([]models.StackedOrder, error)

	GetOrdersForCancel(ctx context.Context) ([]int64, error)
}

type OrderRepo struct {
	db   *transactor.TransactionManager
	name string
}

var _ IOrderRepo = (*OrderRepo)(nil)

func New(pool *transactor.TransactionManager) *OrderRepo {
	return &OrderRepo{
		db:   pool,
		name: "orders",
	}
}
