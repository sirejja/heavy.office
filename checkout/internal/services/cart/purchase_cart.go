package cart

import (
	"context"
	"fmt"
	"route256/checkout/internal/models"
)

func (c *Cart) PurchaseCart(ctx context.Context, user int64) (int64, error) {
	op := "Cart.PurchaseCart"

	itemsCart, _, err := c.ListCart(ctx, user)
	if err != nil {
		return int64(0), fmt.Errorf("%s: %w", op, err)
	}

	purchaseCart := make([]models.Item, len(itemsCart))
	for _, item := range purchaseCart {
		purchaseCart = append(purchaseCart, models.Item{SKU: item.SKU, Count: item.Count})
	}

	orderID, err := c.lomsClient.CreateOrder(ctx, user, purchaseCart)
	if err != nil {
		return int64(0), fmt.Errorf("%s: %w", op, err)
	}

	return orderID, nil
}
