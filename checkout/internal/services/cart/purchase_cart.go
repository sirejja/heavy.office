package cart

import (
	"context"
	"fmt"
)

func (c *Cart) PurchaseCart(ctx context.Context, user int64) (int64, error) {
	op := "Cart.PurchaseCart"

	var OrderID int64
	err := c.txManager.RunRepeatableRead(ctx, func(ctxTX context.Context) error {
		cartProducts, err := c.cartsProductsRepo.GetCartsProducts(ctxTX, user)
		if err != nil {
			return fmt.Errorf("%s: %w", op, err)
		}

		OrderID, err = c.lomsClient.CreateOrder(ctxTX, user, cartProducts)
		if err != nil {
			return fmt.Errorf("%s: %w", op, err)
		}

		_, err = c.cartsRepo.PurchaseCart(ctxTX, user)
		if err != nil {
			return fmt.Errorf("%s: %w", op, err)
		}
		return nil
	})
	if err != nil {
		return int64(0), fmt.Errorf("%s: %w", op, err)
	}

	return OrderID, nil
}
