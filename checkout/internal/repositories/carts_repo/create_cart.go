package carts_repo

import (
	"context"
	"fmt"
	"route256/checkout/internal/models"

	sq "github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/pgxscan"
)

func (c *cartsRepo) CreateCart(ctx context.Context, UserID int64) (uint64, error) {
	op := "cartsRepo.CreateCart"
	db := c.db.GetQueryEngine(ctx)

	if UserID == 0 {
		return 0, fmt.Errorf("%s: %w", op, models.ErrNoDataProvided)
	}

	query := sq.Insert(c.name).
		Columns("user_id").
		Values(UserID).
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
