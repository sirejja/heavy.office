package carts_products_repo

import (
	"context"
	"fmt"
	"route256/checkout/internal/models"

	sq "github.com/Masterminds/squirrel"
)

type AddProductToCartInsert struct {
	CartID uint64
	SKU    uint32
	Count  uint32
}

func (c *cartsProductsRepo) AddProductToCart(ctx context.Context, ins *AddProductToCartInsert) (uint64, error) {
	op := "cartsProductsRepo.AddProductToCart"

	if ins == nil {
		return 0, fmt.Errorf("%s: %w", op, models.ErrNoDataProvided)
	}

	query := sq.Insert(c.name).Columns("cart_id", "sku", "cnt").
		Values(ins.CartID, ins.SKU, ins.Count).Suffix("RETURNING \"id\"").
		PlaceholderFormat(sq.Dollar)

	sql, args, err := query.ToSql()
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	row := c.db.QueryRow(ctx, sql, args...)

	var id uint64
	err = row.Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return id, nil
}
