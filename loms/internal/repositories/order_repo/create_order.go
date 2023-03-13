package order_repo

import (
	"context"
	"fmt"
	"route256/loms/internal/models"

	sq "github.com/Masterminds/squirrel"
)

type CreateOrderIns struct {
	UserID int64
	Status string
}

func (o *OrderRepo) CreateOrder(ctx context.Context, ins *CreateOrderIns) (uint64, error) {
	op := "OrderRepo.CreateOrder"

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

	if err = o.db.QueryRow(ctx, sql, args...).Scan(&id); err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return id, nil
}
