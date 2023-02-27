package cart

import (
	"context"
	"fmt"
	"route256/checkout/internal/clients/loms"
)

func (c *Cart) PurchaseCart(ctx context.Context, user int64) (*uint64, error) {
	op := "Cart.PurchaseCart"
	// TODO get items
	items := []loms.Items{{SKU: 1, Count: 1}}
	orderID, err := c.lomsClient.CreateOrder(ctx, user, items)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return orderID, nil
}
