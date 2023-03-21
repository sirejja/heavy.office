package cart

import (
	"context"
	"fmt"
	"route256/checkout/internal/models"
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

	cartProduct, err := c.cartsProductsRepo.GetCartProduct(ctx, sku, user)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	if cartProduct == nil {
		return fmt.Errorf("%s: %w", op, models.ErrNothingToDelete)
	}

	product := *cartProduct

	if product.Count <= count || (product.Count-count) == 0 {
		_, err = c.cartsProductsRepo.DeleteProductFromCart(ctx, product.CartID)
		if err != nil {
			return fmt.Errorf("%s: %w", op, err)
		}
		return nil
	}

	decreasedCount := product.Count - count
	_, err = c.cartsProductsRepo.UpdateCartProduct(ctx, uint64(sku), decreasedCount, uint32(product.CartID))
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
