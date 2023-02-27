package warehouse

import (
	"context"
	"route256/loms/internal/models"
	"route256/loms/internal/repo/warehouse_repo"
)

type WarehouseProcessor interface {
	Stocks(ctx context.Context, SKU uint32) (*[]models.Stock, error)
}

type Warehouse struct {
	warehouseRepo warehouse_repo.IWarehouseRepo
}

func New(warehouseRepo warehouse_repo.IWarehouseRepo) WarehouseProcessor {
	return &Warehouse{
		warehouseRepo: warehouseRepo,
	}
}
