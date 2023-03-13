package order_repo

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/pgxscan"
)

func (o *OrderRepo) PayedOrder(ctx context.Context, orderID int64) (uint64, error) {
	op := "OrderRepo.PayedOrder"

	query := sq.Update(o.name).
		Set("updated_at", sq.Expr("current_timestamp")).
		Set("status", "payed").
		Where(sq.Eq{"id": orderID}).
		PlaceholderFormat(sq.Dollar)

	sql, args, err := query.ToSql()
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	row := o.db.QueryRow(ctx, sql, args...)
	var id uint64
	if err = row.Scan(&id); err != nil {
		if pgxscan.NotFound(err) {
			return 0, nil
		}
		return 0, fmt.Errorf("%s: %w", op, err)
	}
	return id, nil
}
