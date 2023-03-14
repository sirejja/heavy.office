package order_repo

import (
	"context"
	"fmt"
	"route256/loms/internal/models"

	sq "github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/pgxscan"
)

type CreateOrderIns struct {
	UserID int64
	Status string
}

func (o *OrderRepo) CreateOrder(ctx context.Context, ins *CreateOrderIns) (uint64, error) {
	op := "OrderRepo.CreateOrder"
	db := o.db.GetQueryEngine(ctx)

	if ins == nil {
		return 0, fmt.Errorf("%s: %w", op, models.ErrNoDataProvided)
	}

	query := sq.Insert(o.name).Columns("user_id, status").
		Values(ins.UserID, ins.Status).Suffix("RETURNING \"id\"").
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
