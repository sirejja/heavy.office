package carts_repo

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/pgxscan"
)

func (c *cartsRepo) PurchaseCart(ctx context.Context, userID int64) (uint64, error) {
	op := "cartsRepo.PurchaseCart"
	db := c.db.GetQueryEngine(ctx)

	query := sq.Update(c.name).
		Set("purchased_at", sq.Expr("current_timestamp")).
		Set("updated_at", sq.Expr("current_timestamp")).
		Where(sq.Eq{"purchased_at": nil}).
		Where(sq.Eq{"deleted_at": nil}).
		Where(sq.Eq{"user_id": userID}).
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
