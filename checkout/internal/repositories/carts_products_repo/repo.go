package carts_products_repo

import (
	"context"
	"route256/checkout/internal/models"
	"route256/libs/transactor"
)

type ICartsProductsRepo interface {
	GetCartsProducts(ctx context.Context, userID int64) ([]models.Item, error)
	GetCartProduct(ctx context.Context, sku uint32, userID int64) (*models.ItemCart, error)
	GetCartProductCount(ctx context.Context, SKU uint32) (uint32, error)
	AddProductToCart(ctx context.Context, CartID uint64, sku uint32, Count uint32) (uint64, error)
	UpdateCartProduct(ctx context.Context, sku uint64, count uint32, cartID uint32) (uint64, error)
	DeleteProductFromCart(ctx context.Context, id uint64) (uint64, error)
}

var _ ICartsProductsRepo = (*cartsProductsRepo)(nil)

type cartsProductsRepo struct {
	db   *transactor.TransactionManager
	name string
}

func New(pool *transactor.TransactionManager) *cartsProductsRepo {
	return &cartsProductsRepo{
		db:   pool,
		name: "carts_products",
	}
}
