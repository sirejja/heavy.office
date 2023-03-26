package carts_repo

//go:generate sh -c "rm -rf mocks && mkdir -p mocks"
//go:generate minimock -i ICartsRepo -o ./mocks/ -s "_minimock.go"

import (
	"context"
	"route256/libs/transactor"
)

type ICartsRepo interface {
	GetCartID(ctx context.Context, userID int64) (uint64, error)
	CreateCart(ctx context.Context, UserID int64) (uint64, error)
	PurchaseCart(ctx context.Context, userID int64) (uint64, error)
}

var _ ICartsRepo = (*cartsRepo)(nil)

type cartsRepo struct {
	db   *transactor.TransactionManager
	name string
}

func New(pool *transactor.TransactionManager) *cartsRepo {
	return &cartsRepo{
		db:   pool,
		name: "carts",
	}
}
