package carts_repo

import (
	"context"
	"fmt"
	"route256/checkout/internal/models"

	sq "github.com/Masterminds/squirrel"
)

type CreateCartIns struct {
	UserID int64
}

func (c *cartsRepo) CreateCart(ctx context.Context, ins *CreateCartIns) (uint64, error) {
	op := "cartsRepo.CreateCart"

	if ins == nil {
		return 0, fmt.Errorf("%s: %w", op, models.ErrNoDataProvided)
	}

	query := sq.Insert(c.name).Columns("user_id").
		Values(ins.UserID).Suffix("RETURNING \"id\"").
		PlaceholderFormat(sq.Dollar)

	sql, args, err := query.ToSql()

	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	var id uint64

	if err = c.db.QueryRow(ctx, sql, args...).Scan(&id); err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return id, nil
}
