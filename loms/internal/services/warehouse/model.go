package warehouse

import (
	"context"
	"route256/libs/transactor"
	"route256/loms/internal/models"
	"route256/loms/internal/repositories/warehouse_repo"
)

type IWarehouseService interface {
	Stocks(ctx context.Context, SKU uint32) ([]models.Stock, error)
}

type Warehouse struct {
	warehouseRepo warehouse_repo.IWarehouseRepo
	txManager     transactor.ITransactor
}

var _ IWarehouseService = (*Warehouse)(nil)

func New(warehouseRepo warehouse_repo.IWarehouseRepo, txManager *transactor.TransactionManager) *Warehouse {
	return &Warehouse{
		warehouseRepo: warehouseRepo,
		txManager:     txManager,
	}
}
