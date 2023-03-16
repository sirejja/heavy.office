package order_repo

import (
	"context"
	"fmt"
	"route256/loms/internal/models"
	"route256/loms/internal/repositories/schema"

	sq "github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/pgxscan"
)

func (o *OrderRepo) GetOrderDetails(ctx context.Context, orderID uint64) (*models.OrderDetails, error) {
	op := "OrderRepo.GetOrderDetails"
	db := o.db.GetQueryEngine(ctx)

	query := sq.Select("user_id, status").
		From(o.name).
		Where(sq.Eq{"id": orderID}).
		Where(sq.Eq{"cancelled_at": nil}).
		PlaceholderFormat(sq.Dollar)

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	var orderRow schema.OrderDetails
	if err = pgxscan.Get(ctx, db, &orderRow, sql, args...); err != nil {
		if pgxscan.NotFound(err) {
			return nil, nil
		}
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &models.OrderDetails{UserID: orderRow.UserID, Status: orderRow.Status}, nil
}
