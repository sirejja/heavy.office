package order_repo

import (
	"context"
	"fmt"
	"route256/loms/internal/models"

	sq "github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/pgxscan"
)

func (o *OrderRepo) UpdateOrderStatus(ctx context.Context, orderID int64, status models.OrderStatus) (uint64, error) {
	op := "OrderRepo.CancelOrder"
	db := o.db.GetQueryEngine(ctx)

	query := sq.Update(o.name).
		Set("status", status.ToString()).
		Set("updated_at", sq.Expr("current_timestamp")).
		Where(sq.Eq{"id": orderID}).
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
