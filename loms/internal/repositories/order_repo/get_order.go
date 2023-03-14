package order_repo

import (
	"context"
	"fmt"
	"route256/loms/internal/models"
	"route256/loms/internal/repositories/schema"

	sq "github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/pgxscan"
)

type GetOrderFilter struct {
	Id     uint64
	UserId int64
	Status string
}

func (o *OrderRepo) GetOrder(ctx context.Context, filter *GetOrderFilter) (*schema.OrdersSchema, error) {
	op := "OrderRepo.GetOrder"
	db := o.db.GetQueryEngine(ctx)

	if filter == nil {
		return nil, fmt.Errorf("%s: %w", op, models.ErrNoFiltersProvided)
	}

	query := sq.Select("id, user_id, status, created_at, updated_at, cancelled_at").
		From(o.name).PlaceholderFormat(sq.Dollar)

	if filter.Id != 0 {
		query = query.Where(sq.Eq{"id": filter.Id})
	}
	if filter.UserId != 0 {
		query = query.Where(sq.Eq{"user_id": filter.UserId})
	}

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	var orderRow schema.OrdersSchema
	if err = pgxscan.Get(ctx, db, &orderRow, sql, args...); err != nil {
		if pgxscan.NotFound(err) {
			return nil, nil
		}
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &orderRow, nil
}
