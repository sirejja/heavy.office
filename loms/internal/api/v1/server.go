package v1

import (
	"route256/loms/internal/services/orders"
	"route256/loms/internal/services/warehouse"
	desc "route256/loms/pkg/grpc/server"
)

type Server struct {
	warehouse warehouse.Warehouse
	orders    orders.Order

	desc.UnimplementedLomsServer
}

var _ desc.LomsServer = (*Server)(nil)

func New(warehouse warehouse.Warehouse, orders orders.Order) *Server {
	return &Server{
		warehouse: warehouse,
		orders:    orders,
	}
}
