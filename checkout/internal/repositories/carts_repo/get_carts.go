package carts_repo

import (
	"context"
	"fmt"
	"route256/checkout/internal/models"
	"route256/checkout/internal/repositories/schema"

	sq "github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/pgxscan"
)

type GetCartFilter struct {
	Id        uint64
	UserId    int64
	IsDeleted bool
}

func (c *cartsRepo) GetCarts(ctx context.Context, filter *GetCartFilter) ([]*schema.CartSchema, error) {
	op := "cartsRepo.GetCarts"

	if filter == nil {
		return nil, fmt.Errorf("%s: %w", op, models.ErrNoFiltersProvided)
	}

	query := sq.Select("id, user_id, created_at, updated_at, deleted_at, purchased_at").
		From(c.name).PlaceholderFormat(sq.Dollar)

	if filter.Id != 0 {
		query = query.Where(sq.Eq{"id": filter.Id})
	}
	if filter.UserId != 0 {
		query = query.Where(sq.Eq{"user_id": filter.UserId})
	}
	if !filter.IsDeleted {
		query = query.Where(sq.Eq{"deleted_at": nil})
	} else {
		query = query.Where(sq.Expr("deleted_at is not null"))
	}

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	var cartsRows []*schema.CartSchema
	if err = pgxscan.Select(ctx, c.db, &cartsRows, sql, args...); err != nil {
		if pgxscan.NotFound(err) {
			return []*schema.CartSchema{}, nil
		}
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return cartsRows, nil
}
