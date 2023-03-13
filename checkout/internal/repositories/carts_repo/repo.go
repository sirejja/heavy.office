package carts_repo

import (
	"context"
	"route256/checkout/internal/repositories/schema"

	"github.com/jackc/pgx/v4/pgxpool"
)

type ICartsRepo interface {
	GetCarts(ctx context.Context, filter *GetCartFilter) ([]*schema.CartSchema, error)
	CreateCart(ctx context.Context, ins *CreateCartIns) (uint64, error)
	UpdateCart(ctx context.Context, upd *UpdatCartValues, filter *UpdatCartFilter) (uint64, error)
}

var _ ICartsRepo = (*cartsRepo)(nil)

type cartsRepo struct {
	db   *pgxpool.Pool
	name string
}

func New(pool *pgxpool.Pool) *cartsRepo {
	return &cartsRepo{
		db:   pool,
		name: "carts",
	}
}
