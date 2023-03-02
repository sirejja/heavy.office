package warehouse

import (
	"route256/loms/internal/repo/warehouse_repo"
)

type Warehouse struct {
	warehouseRepo warehouse_repo.WarehouseRepo
}

func New(warehouseRepo warehouse_repo.WarehouseRepo) *Warehouse {
	return &Warehouse{
		warehouseRepo: warehouseRepo,
	}
}
