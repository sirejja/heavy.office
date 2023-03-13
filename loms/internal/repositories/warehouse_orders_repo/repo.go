package warehouse_orders_repo

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

type IWarehouseOrdersRepo interface {
	FillOrderProducts(ctx context.Context, ins *FillOrderProductsIns) (uint64, error)
}

var _ IWarehouseOrdersRepo = (*warehouseOrdersRepo)(nil)

type warehouseOrdersRepo struct {
	db   *pgxpool.Pool
	name string
}

func New(pool *pgxpool.Pool) *warehouseOrdersRepo {
	return &warehouseOrdersRepo{
		db:   pool,
		name: "warehouse_orders",
	}
}
