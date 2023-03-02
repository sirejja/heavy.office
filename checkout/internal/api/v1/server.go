package v1

import (
	"route256/checkout/internal/services/cart"
	desc "route256/checkout/pkg/grpc/server"
)

var _ desc.CheckoutServer = (*Server)(nil)

type Server struct {
	model cart.Cart

	desc.UnimplementedCheckoutServer
}

func New(service cart.Cart) *Server {
	return &Server{
		model: service,
	}
}
