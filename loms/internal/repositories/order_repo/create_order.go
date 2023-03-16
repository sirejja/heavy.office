package order_repo

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/pgxscan"
)

func (o *OrderRepo) CreateOrder(ctx context.Context, userID int64, status string) (uint64, error) {
	op := "OrderRepo.CreateOrder"
	db := o.db.GetQueryEngine(ctx)

	query := sq.Insert(o.name).
		Columns("user_id, status").
		Values(userID, status).
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
