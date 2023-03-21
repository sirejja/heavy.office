package carts_products_repo

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/pgxscan"
)

func (c *cartsProductsRepo) DeleteProductFromCart(ctx context.Context, cartProductroductID uint64) (uint64, error) {
	op := "cartsProductsRepo.DeleteProductFromCart"
	db := c.db.GetQueryEngine(ctx)

	query := sq.Update(c.name).
		Set("deleted_at", sq.Expr("current_timestamp")).
		Where(sq.Eq{"id": cartProductroductID}).
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
