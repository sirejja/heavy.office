package warehouse_repo

import (
	"context"
	"route256/loms/internal/repositories/schema"

	"github.com/jackc/pgx/v4/pgxpool"
)

type IWarehouseRepo interface {
	GetStocks(ctx context.Context, filter *GetStocksFilter) ([]*schema.Stock, error)
	UpdateStocks(ctx context.Context, filter *UpdateStocksFilter, data *UpdateStocksData) (uint64, error)
}

var _ IWarehouseRepo = (*warehouseRepo)(nil)

type warehouseRepo struct {
	db   *pgxpool.Pool
	name string
}

func New(pool *pgxpool.Pool) *warehouseRepo {
	return &warehouseRepo{
		db:   pool,
		name: "warehouse",
	}
}
