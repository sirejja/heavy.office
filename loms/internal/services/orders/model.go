package orders

import (
	"route256/loms/internal/repo/order_repo"
	"route256/loms/internal/repo/warehouse_repo"
)

type Order struct {
	ordersRepo    order_repo.OrderRepo
	warehouseRepo warehouse_repo.WarehouseRepo
}

func New(ordersRepo order_repo.OrderRepo, warehouseRepo warehouse_repo.WarehouseRepo) Order {
	return Order{
		ordersRepo:    ordersRepo,
		warehouseRepo: warehouseRepo,
	}
}
