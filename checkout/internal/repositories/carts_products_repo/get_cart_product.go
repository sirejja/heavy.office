package carts_products_repo

import (
	"context"
	"fmt"
	"route256/checkout/internal/models"
	"route256/checkout/internal/repositories/schema"

	sq "github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/pgxscan"
)

func (c *cartsProductsRepo) GetCartProduct(ctx context.Context, sku uint32, userID int64) (*models.ItemCart, error) {
	op := "cartsProductsRepo.GetCartProduct"
	db := c.db.GetQueryEngine(ctx)

	query := sq.Select("cp.id, cp.cart_id, cp.sku, cp.cnt").
		From(fmt.Sprintf("%s cp", c.name)).
		LeftJoin("carts c on c.id=cp.cart_id").
		Where(sq.Eq{"c.user_id": userID}).
		Where(sq.Eq{"cp.sku": sku}).
		Where(sq.Eq{"cp.deleted_at": nil}).
		Where(sq.Eq{"c.deleted_at": nil}).
		PlaceholderFormat(sq.Dollar)

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	var cartProduct schema.CartProductsSchema
	if err = pgxscan.Get(ctx, db, &cartProduct, sql, args...); err != nil {
		if pgxscan.NotFound(err) {
			return nil, nil
		}
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &models.ItemCart{CartID: cartProduct.CartID, Count: cartProduct.Count, SKU: cartProduct.SKU}, nil
}
