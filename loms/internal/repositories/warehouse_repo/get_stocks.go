package warehouse_repo

import "route256/loms/internal/models"

func (w WarehouseRepo) GetStocks(SKU uint32) ([]models.Stock, error) {
	return []models.Stock{models.Stock{WarehouseID: 111, Count: 11}}, nil
}
