package orders

import (
	"context"
	"route256/libs/transactor"
	"route256/loms/internal/models"
	"route256/loms/internal/repositories/order_repo"
	"route256/loms/internal/repositories/warehouse_orders_repo"
	"route256/loms/internal/repositories/warehouse_repo"
)

type IOrdersService interface {
	CreateOrder(ctx context.Context, user int64, items []models.Item) (int64, error)
	CancelOrder(ctx context.Context, orderID int64) error
	ListOrder(ctx context.Context, orderID int64) (*models.Order, error)
	PayedOrder(ctx context.Context, orderID int64) error
}

type Order struct {
	ordersRepo          order_repo.IOrderRepo
	warehouseRepo       warehouse_repo.IWarehouseRepo
	warehouseOrdersRepo warehouse_orders_repo.IWarehouseOrdersRepo
	txManager           transactor.ITransactor
}

func New(
	ordersRepo order_repo.IOrderRepo,
	warehouseRepo warehouse_repo.IWarehouseRepo,
	warehouseOrdersRepo warehouse_orders_repo.IWarehouseOrdersRepo,
	txManager *transactor.TransactionManager,
) *Order {
	return &Order{
		ordersRepo:          ordersRepo,
		warehouseRepo:       warehouseRepo,
		warehouseOrdersRepo: warehouseOrdersRepo,
		txManager:           txManager,
	}
}
