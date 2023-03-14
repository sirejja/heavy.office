package carts_products_repo

import (
	"context"
	"fmt"
	"route256/checkout/internal/models"

	sq "github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/pgxscan"
)

type DeleteProductFromCartFilter struct {
	ID uint64
}

func (c *cartsProductsRepo) DeleteProductFromCart(ctx context.Context, filter *DeleteProductFromCartFilter) (uint64, error) {
	op := "cartsProductsRepo.DeleteProductFromCart"
	db := c.db.GetQueryEngine(ctx)

	if filter == nil {
		return 0, fmt.Errorf("%s: %w", op, models.ErrNoFiltersProvided)
	}

	query := sq.Update(c.name).Set("deleted_at", sq.Expr("current_timestamp")).
		Where(sq.Eq{"id": filter.ID}).Suffix("RETURNING \"id\"").
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
