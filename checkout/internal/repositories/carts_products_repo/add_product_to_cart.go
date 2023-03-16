package carts_products_repo

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/pgxscan"
)

func (c *cartsProductsRepo) AddProductToCart(ctx context.Context, CartID uint64, sku uint32, Count uint32) (uint64, error) {
	op := "cartsProductsRepo.AddProductToCart"
	db := c.db.GetQueryEngine(ctx)

	query := sq.Insert(c.name).
		Columns("cart_id", "sku", "cnt").
		Values(CartID, sku, Count).
		Suffix("RETURNING \"id\"").
		PlaceholderFormat(sq.Dollar)

	sql, args, err := query.ToSql()
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	var id uint64
	if err = pgxscan.Get(ctx, db, &id, sql, args...); err != nil {
		if pgxscan.NotFound(err) {
			return 0, nil
		}
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return id, nil
}
