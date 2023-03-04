package v1

import (
	"route256/loms/internal/services/orders"
	"route256/loms/internal/services/warehouse"
	desc "route256/loms/pkg/v1/api"
)

type Server struct {
	warehouse warehouse.IWarehouseService
	orders    orders.IOrdersService

	desc.UnimplementedLomsServer
}

var _ desc.LomsServer = (*Server)(nil)

func New(warehouse warehouse.IWarehouseService, orders orders.IOrdersService) *Server {
	return &Server{
		warehouse: warehouse,
		orders:    orders,
	}
}
