package order_repo

import (
	"context"
	"fmt"
	"route256/loms/internal/models"

	sq "github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/pgxscan"
)

func (o *OrderRepo) GetOrdersForCancel(ctx context.Context) ([]int64, error) {
	op := "OrderRepo.GetOrdersForCancel"
	db := o.db.GetQueryEngine(ctx)

	query := sq.
		Select("id").
		From(o.name).
		Where("created_at < current_timestamp - interval '10 minutes'").
		Where(sq.Eq{"status": models.OrderStatusWaitPayment}).
		Where(sq.Eq{"cancelled_at": nil}).
		PlaceholderFormat(sq.Dollar)

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	var orderIDs []*int64
	if err = pgxscan.Select(ctx, db, &orderIDs, sql, args...); err != nil {
		if pgxscan.NotFound(err) {
			return nil, nil
		}
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	result := make([]int64, 0, len(orderIDs))
	for _, orderID := range orderIDs {
		result = append(result, *orderID)
	}

	return result, nil
}
