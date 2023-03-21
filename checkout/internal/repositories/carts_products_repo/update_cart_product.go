package carts_products_repo

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/pgxscan"
)

func (c *cartsProductsRepo) UpdateCartProduct(ctx context.Context, sku uint64, count uint32, cartID uint32) (uint64, error) {
	op := "cartsProductsRepo.UpdateProductCart"
	db := c.db.GetQueryEngine(ctx)

	query := sq.Update(c.name).
		Set("cnt", count).
		Set("updated_at", sq.Expr("current_timestamp")).
		Where(sq.Eq{"cart_id": cartID}).
		Where(sq.Eq{"sku": sku}).
		Where(sq.Eq{"deleted_at": nil}).
		Suffix("RETURNING id").
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
