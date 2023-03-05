package warehouse_repo

import "route256/loms/internal/models"

func (w WarehouseRepo) BookProducts(user int64, items []models.Item) (int64, error) {
	var orderID int64 = 555
	return orderID, nil
}
