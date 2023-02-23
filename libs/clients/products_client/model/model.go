package model

import (
	"context"
)

type ProductsService interface {
	GetProduct(ctx context.Context, SKU uint32) (*Stock, error)
	ListSKUs(ctx context.Context, SKU uint32) error
}

type Model struct {
	productsService ProductsService
}

func New(productsService ProductsService) *Model {
	return &Model{
		productsService: productsService,
	}
}

type Stock struct {
	WarehouseID int64
	Count       uint64
}
