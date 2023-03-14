package order_repo

import (
	"context"
	"route256/libs/transactor"
	"route256/loms/internal/repositories/schema"
)

type IOrderRepo interface {
	GetOrder(ctx context.Context, filter *GetOrderFilter) (*schema.OrdersSchema, error)
	PayedOrder(ctx context.Context, orderID int64) (uint64, error)
	CancelOrder(ctx context.Context, orderID int64) (uint64, error)
	CreateOrder(ctx context.Context, ins *CreateOrderIns) (uint64, error)
	ListOrder(ctx context.Context, filter *ListOrderFilter) ([]*schema.WarehouseOrdersList, error)
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
