package carts_repo

import (
	"context"
	"fmt"
	"route256/checkout/internal/models"

	sq "github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/pgxscan"
)

type UpdatCartFilter struct {
	ID          uint64
	UserID      int64
	IsDeleted   bool
	IsPurchased bool
}
type UpdatCartValues struct {
	IsDeleted   bool
	IsPurchased bool
}

func (c *cartsRepo) UpdateCart(ctx context.Context, upd *UpdatCartValues, filter *UpdatCartFilter) (uint64, error) {
	op := "cartsRepo.UpdateCart"
	db := c.db.GetQueryEngine(ctx)

	if upd == nil {
		return 0, fmt.Errorf("%s: %w", op, models.ErrNoDataProvided)
	}
	if filter == nil {
		return 0, fmt.Errorf("%s: %w", op, models.ErrNoFiltersProvided)
	}

	query := sq.Update(c.name).PlaceholderFormat(sq.Dollar).Suffix("RETURNING id")

	if filter.UserID != 0 {
		query = query.Where(sq.Eq{"user_id": filter.UserID})
	}
	if filter.ID != 0 {
		query = query.Where(sq.Eq{"id": filter.ID})
	}
	if filter.IsDeleted {
		query = query.Where(sq.Expr("deleted_at is not null"))
	} else {
		query = query.Where(sq.Eq{"deleted_at": nil})
	}
	if filter.IsPurchased {
		query = query.Where(sq.Expr("purchased_at is not null"))
	} else {
		query = query.Where(sq.Eq{"purchased_at": nil})
	}

	if upd.IsPurchased {
		query = query.Set("purchased_at", sq.Expr("current_timestamp"))
	}
	if upd.IsDeleted {
		query = query.Set("deleted_at", sq.Expr("current_timestamp"))
	}

	query = query.Set("updated_at", sq.Expr("current_timestamp"))

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
