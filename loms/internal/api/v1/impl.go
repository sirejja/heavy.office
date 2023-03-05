package v1

import (
	"route256/loms/internal/services/orders"
	"route256/loms/internal/services/warehouse"
	desc "route256/loms/pkg/v1/api"
)

type Implementation struct {
	warehouse warehouse.IWarehouseService
	orders    orders.IOrdersService

	desc.UnimplementedLomsServer
}

var _ desc.LomsServer = (*Implementation)(nil)

func New(warehouse warehouse.IWarehouseService, orders orders.IOrdersService) *Implementation {
	return &Implementation{
		warehouse: warehouse,
		orders:    orders,
	}
}
