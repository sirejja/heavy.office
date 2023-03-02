package orders

import (
	"context"
	"route256/loms/internal/models"
	"route256/loms/internal/repo/order_repo"
	"route256/loms/internal/repo/warehouse_repo"
)

type OrdersProcessor interface {
	CreateOrder(ctx context.Context, user int64, items []models.Item) (*uint64, error)
	CancelOrder(ctx context.Context, orderID int64) error
	ListOrder(ctx context.Context, orderID int64) (*order_repo.Order, error)
	PayedOrder(ctx context.Context, orderID int64) error
}

type Order struct {
	ordersRepo    order_repo.IOrderRepo
	warehouseRepo warehouse_repo.IWarehouseRepo
}

func New(ordersRepo order_repo.IOrderRepo, warehouseRepo warehouse_repo.IWarehouseRepo) OrdersProcessor {
	return &Order{
		ordersRepo:    ordersRepo,
		warehouseRepo: warehouseRepo,
	}
}
