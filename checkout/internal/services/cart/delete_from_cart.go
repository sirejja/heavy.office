package cart

import (
	"context"
	"fmt"
	"route256/checkout/internal/models"
	"route256/checkout/internal/repositories/carts_products_repo"
)

func (c *Cart) DeleteFromCart(ctx context.Context, user int64, sku uint32, count uint32) error {
	op := "Cart.DeleteFromCart"

	err := c.txManager.RunRepeatableRead(ctx, func(ctxTX context.Context) error {
		err := c.processProductDeletionFromCart(ctxTX, user, sku, count)
		if err != nil {
			return fmt.Errorf("%s: %w", op, err)
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (c *Cart) processProductDeletionFromCart(ctx context.Context, user int64, sku uint32, count uint32) error {
	op := "Cart.processProductDeletionFromCart"

	cartProducts, err := c.cartsProductsRepo.GetCartsProducts(ctx,
		&carts_products_repo.GetCartProductsFilter{SKU: sku, UserID: user, IsDeleted: false})
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	if len(cartProducts) == 0 {
		return fmt.Errorf("%s: %w", op, models.ErrNothingToDelete)
	}

	product := *cartProducts[0]
	if product.Count <= count || (product.Count-count) == 0 {
		_, err = c.cartsProductsRepo.DeleteProductFromCart(ctx,
			&carts_products_repo.DeleteProductFromCartFilter{ID: product.ID})
		if err != nil {
			return fmt.Errorf("%s: %w", op, err)
		}
		return nil
	}

	decreasedCount := product.Count - count
	_, err = c.cartsProductsRepo.UpdateCartProduct(ctx,
		&carts_products_repo.UpdateProductCartValues{SKU: sku, Count: decreasedCount},
		&carts_products_repo.UpdateProductCartFilter{CartID: product.CartID, SKU: sku, IsDeleted: false})
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
