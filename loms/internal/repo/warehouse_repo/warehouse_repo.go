package warehouse_repo

import (
	"route256/loms/internal/models"
)

type IWarehouseRepo interface {
	BookProducts(user int64, items []models.Item) (int64, error)
	GetStocks(SKU uint32) (*[]models.Stock, error)
}

var _ IWarehouseRepo = (*WarehouseRepo)(nil)

type WarehouseRepo struct {
}

func New() WarehouseRepo {
	return WarehouseRepo{}
}
