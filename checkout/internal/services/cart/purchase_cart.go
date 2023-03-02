package cart

import (
	"context"
	"fmt"
	"route256/checkout/internal/models"
)

func (c *Cart) PurchaseCart(ctx context.Context, user int64) (int64, error) {
	op := "Cart.PurchaseCart"
	// TODO get items
	items := []models.Items{{SKU: 1, Count: 1}}
	orderID, err := c.lomsClient.CreateOrder(ctx, user, items)
	if err != nil {
		return int64(0), fmt.Errorf("%s: %w", op, err)
	}
	return orderID, nil
}
