package carts_repo

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/pgxscan"
)

func (c *cartsRepo) GetCartID(ctx context.Context, userID int64) (uint64, error) {
	op := "cartsRepo.GetCartID"
	db := c.db.GetQueryEngine(ctx)

	query := sq.Select("id").
		From(c.name).
		Where(sq.Eq{"user_id": userID}).
		Where(sq.Eq{"deleted_at": nil}).
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
