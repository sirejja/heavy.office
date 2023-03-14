package warehouse_repo

import (
	"context"
	"route256/libs/transactor"
	"route256/loms/internal/repositories/schema"
)

type IWarehouseRepo interface {
	GetStocks(ctx context.Context, filter *GetStocksFilter) ([]*schema.Stock, error)
	UpdateStocks(ctx context.Context, filter *UpdateStocksFilter, data *UpdateStocksData) (uint64, error)
}

var _ IWarehouseRepo = (*warehouseRepo)(nil)

type warehouseRepo struct {
	db   *transactor.TransactionManager
	name string
}

func New(pool *transactor.TransactionManager) *warehouseRepo {
	return &warehouseRepo{
		db:   pool,
		name: "warehouse",
	}
}
