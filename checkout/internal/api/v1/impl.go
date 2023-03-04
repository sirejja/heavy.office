package v1

import (
	"route256/checkout/internal/services/cart"
	desc "route256/checkout/pkg/v1/api"
)

var _ desc.CheckoutServer = (*Implementation)(nil)

type Implementation struct {
	model cart.ICartProcessor

	desc.UnimplementedCheckoutServer
}

func New(service cart.ICartProcessor) *Implementation {
	return &Implementation{
		model: service,
	}
}
