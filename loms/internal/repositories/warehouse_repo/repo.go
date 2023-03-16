package warehouse_repo

import (
	"context"
	"route256/libs/transactor"
	"route256/loms/internal/models"
)

type IWarehouseRepo interface {
	GetStocks(ctx context.Context, sku uint32) ([]models.Stock, error)
	ChangeStocks(ctx context.Context, warehouseID uint64, StockDiff int32) (uint64, error)
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
