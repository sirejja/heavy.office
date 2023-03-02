package cart

import (
	"context"
	"route256/checkout/internal/clients/loms"
	"route256/checkout/internal/clients/products"
	"route256/checkout/internal/models"
)

type ICartProcessor interface {
	AddToCart(ctx context.Context, user int64, sku uint32, count uint16) error
	ListCart(ctx context.Context, user int64) (*[]models.Product, *uint32, error)
	DeleteFromCart(ctx context.Context, user int64, sku uint32, count uint16) error
	PurchaseCart(ctx context.Context, user int64) (*uint64, error)
}

type Cart struct {
	lomsClient     *loms.Client
	productsClient *products.Client
}

func New(lomsClient *loms.Client, productsClient *products.Client) ICartProcessor {
	return &Cart{
		lomsClient:     lomsClient,
		productsClient: productsClient,
	}
}
