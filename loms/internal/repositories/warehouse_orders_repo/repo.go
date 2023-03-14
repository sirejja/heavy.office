package warehouse_orders_repo

import (
	"context"
	"route256/libs/transactor"
)

type IWarehouseOrdersRepo interface {
	FillOrderProducts(ctx context.Context, ins *FillOrderProductsIns) (uint64, error)
}

var _ IWarehouseOrdersRepo = (*warehouseOrdersRepo)(nil)

type warehouseOrdersRepo struct {
	db   *transactor.TransactionManager
	name string
}

func New(pool *transactor.TransactionManager) *warehouseOrdersRepo {
	return &warehouseOrdersRepo{
		db:   pool,
		name: "warehouse_orders",
	}
}
