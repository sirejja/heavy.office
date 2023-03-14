package carts_products_repo

import (
	"context"
	"route256/checkout/internal/repositories/schema"
	"route256/libs/transactor"
)

type ICartsProductsRepo interface {
	AddProductToCart(ctx context.Context, ins *AddProductToCartInsert) (uint64, error)
	GetCartsProducts(ctx context.Context, filter *GetCartProductsFilter) ([]*schema.CartProductsSchema, error)
	DeleteProductFromCart(ctx context.Context, filter *DeleteProductFromCartFilter) (uint64, error)
	UpdateCartProduct(ctx context.Context, upd *UpdateProductCartValues, filter *UpdateProductCartFilter) (uint64, error)
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
