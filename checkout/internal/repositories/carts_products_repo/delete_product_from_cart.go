package carts_products_repo

import (
	"context"
	"fmt"
	"route256/checkout/internal/models"

	sq "github.com/Masterminds/squirrel"
)

type DeleteProductFromCartFilter struct {
	ID uint64
}

func (c *cartsProductsRepo) DeleteProductFromCart(ctx context.Context, filter *DeleteProductFromCartFilter) (uint64, error) {
	op := "cartsProductsRepo.DeleteProductFromCart"

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

	row := c.db.QueryRow(ctx, sql, args...)

	var id uint64

	if err = row.Scan(&id); err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return id, nil
}
