package cart

import (
	"context"
	"route256/checkout/internal/clients/loms"
	"route256/checkout/internal/clients/products"
	"route256/checkout/internal/models"
)

type ICartProcessor interface {
	AddToCart(ctx context.Context, user int64, sku uint32, count uint32) error
	ListCart(ctx context.Context, user int64) ([]models.CartProduct, uint32, error)
	DeleteFromCart(ctx context.Context, user int64, sku uint32, count uint32) error
	PurchaseCart(ctx context.Context, user int64) (int64, error)
}

type Cart struct {
	lomsClient     loms.ILOMSClient
	productsClient products.IProductServiceClient
}

func New(lomsClient loms.ILOMSClient, productsClient products.IProductServiceClient) *Cart {
	return &Cart{
		lomsClient:     lomsClient,
		productsClient: productsClient,
	}
}
