package carts_products_repo

import (
	"context"
	"fmt"
	"route256/checkout/internal/models"

	sq "github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/pgxscan"
)

func (c *cartsProductsRepo) GetCartProductCount(ctx context.Context, SKU uint32) (uint32, error) {
	op := "cartsProductsRepo.GetCartsProductsCount"
	db := c.db.GetQueryEngine(ctx)

	if SKU == 0 {
		return 0, fmt.Errorf("%s: %w", op, models.ErrDBEmptySKU)
	}

	query := sq.Select("cnt").
		From(c.name).
		Where(sq.Eq{"deleted_at": nil}).
		Where(sq.Eq{"sku": SKU}).
		PlaceholderFormat(sq.Dollar)

	sql, args, err := query.ToSql()
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	var cnt uint32
	if err = pgxscan.Get(ctx, db, &cnt, sql, args...); err != nil {
		if pgxscan.NotFound(err) {
			return 0, nil
		}
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return cnt, nil
}
