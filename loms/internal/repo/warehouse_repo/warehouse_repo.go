package warehouse_repo

import (
	"route256/loms/internal/models"
)

type Warehouse struct{}

type IWarehouseRepo interface {
	BookProducts(user int64, items []models.Item) (*uint64, error)
	GetStocks(SKU uint32) (*[]models.Stock, error)
}

func New() IWarehouseRepo {
	return &Warehouse{}
}

func (w Warehouse) BookProducts(user int64, items []models.Item) (*uint64, error) {
	var orderId uint64 = 555
	return &orderId, nil
}

func (w Warehouse) GetStocks(SKU uint32) (*[]models.Stock, error) {
	return &[]models.Stock{{WarehouseID: 111, Count: 11}}, nil
}
