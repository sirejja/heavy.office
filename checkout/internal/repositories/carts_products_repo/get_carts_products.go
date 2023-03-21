package carts_products_repo

import (
	"context"
	"fmt"
	"route256/checkout/internal/models"
	"route256/checkout/internal/repositories/schema"

	sq "github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/pgxscan"
)

func (c *cartsProductsRepo) GetCartsProducts(ctx context.Context, userID int64) ([]models.Item, error) {
	op := "cartsProductsRepo.GetCartsProducts"
	db := c.db.GetQueryEngine(ctx)

	query := sq.Select("cp.id, cp.cart_id, cp.sku, cp.cnt").
		From(fmt.Sprintf("%s cp", c.name)).
		LeftJoin("carts c on c.id = cp.cart_id").
		Where(sq.Eq{"cp.deleted_at": nil}).
		Where(sq.Eq{"c.user_id": userID}).
		PlaceholderFormat(sq.Dollar)

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	var cartsProductsRows []*schema.CartProductsSchema
	if err = pgxscan.Select(ctx, db, &cartsProductsRows, sql, args...); err != nil {
		if pgxscan.NotFound(err) {
			return nil, nil
		}
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	products := make([]models.Item, 0, len(cartsProductsRows))
	for _, item := range cartsProductsRows {
		products = append(products, models.Item{SKU: item.SKU, Count: item.Count})
	}

	return products, nil
}
