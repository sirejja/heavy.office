package carts_products_repo

import (
	"context"
	"fmt"
	"route256/checkout/internal/models"

	sq "github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/pgxscan"
)

type UpdateProductCartFilter struct {
	CartID    uint64
	SKU       uint32
	IsDeleted bool
}
type UpdateProductCartValues struct {
	SKU   uint32
	Count uint32
}

func (c *cartsProductsRepo) UpdateCartProduct(ctx context.Context, upd *UpdateProductCartValues, filter *UpdateProductCartFilter) (uint64, error) {
	op := "cartsProductsRepo.UpdateProductCart"
	db := c.db.GetQueryEngine(ctx)

	if upd == nil {
		return 0, fmt.Errorf("%s: %w", op, models.ErrNoDataProvided)
	}
	if filter == nil {
		return 0, fmt.Errorf("%s: %w", op, models.ErrNoFiltersProvided)
	}

	query := sq.Update(c.name).PlaceholderFormat(sq.Dollar).Suffix("RETURNING id")

	if filter.CartID != 0 && filter.SKU != 0 {
		query = query.Where(sq.Eq{"cart_id": filter.CartID}).Where(sq.Eq{"sku": filter.SKU})
	} else {
		return 0, fmt.Errorf("%s: %w", op, models.ErrNoFiltersProvided)
	}

	if filter.IsDeleted {
		query = query.Where(sq.Expr("deleted_at is not null"))
	} else {
		query = query.Where(sq.Eq{"deleted_at": nil})
	}

	if upd.Count != 0 && upd.SKU != 0 {
		query = query.Set("cnt", upd.Count)
	} else {
		return 0, fmt.Errorf("%s: %w", op, models.ErrNoDataProvided)
	}

	query = query.Set("updated_at", sq.Expr("current_timestamp"))

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
