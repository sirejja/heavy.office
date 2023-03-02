package cart

import (
	"route256/checkout/internal/clients/loms"
	"route256/checkout/internal/clients/products"
)

type Cart struct {
	lomsClient     *loms.Client
	productsClient *products.Client
}

func New(lomsClient *loms.Client, productsClient *products.Client) *Cart {
	return &Cart{
		lomsClient:     lomsClient,
		productsClient: productsClient,
	}
}
