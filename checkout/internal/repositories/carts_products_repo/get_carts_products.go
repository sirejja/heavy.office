package carts_products_repo

import (
	"context"
	"fmt"
	"route256/checkout/internal/models"
	"route256/checkout/internal/repositories/schema"

	sq "github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/pgxscan"
)

type GetCartProductsFilter struct {
	ID        uint64
	SKU       uint32
	UserID    int64
	IsDeleted bool
}

func (c *cartsProductsRepo) GetCartsProducts(ctx context.Context, filter *GetCartProductsFilter) ([]*schema.CartProductsSchema, error) {
	op := "cartsProductsRepo.GetCartsProducts"
	if filter == nil {
		return nil, fmt.Errorf("%s: %w", op, models.ErrNoFiltersProvided)
	}

	query := sq.Select("cp.id, cp.cart_id, cp.sku, cp.cnt, cp.created_at, cp.updated_at, cp.deleted_at").
		From(fmt.Sprintf("%s cp", c.name)).PlaceholderFormat(sq.Dollar)

	if filter.ID != 0 {
		query = query.Where(sq.Eq{"cp.id": filter.ID})
	}
	if filter.SKU != 0 {
		query = query.Where(sq.Eq{"cp.sku": filter.SKU})
	}
	if !filter.IsDeleted {
		query = query.Where(sq.Eq{"cp.deleted_at": nil})
	} else {
		query = query.Where(sq.Expr("cp.deleted_at is not null"))
	}

	if filter.UserID != 0 && filter.ID == 0 {
		query = query.LeftJoin("carts c on c.id=cp.cart_id")
		query = query.Where(sq.Eq{"c.user_id": filter.UserID})
	}

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	var cartsProductsRows []*schema.CartProductsSchema
	if err = pgxscan.Select(ctx, c.db, &cartsProductsRows, sql, args...); err != nil {
		if pgxscan.NotFound(err) {
			return []*schema.CartProductsSchema{}, nil
		}
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return cartsProductsRows, nil
}
