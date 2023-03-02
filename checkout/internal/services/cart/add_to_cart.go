package cart

import (
	"context"
	"fmt"
	"route256/checkout/internal/models"
)

func (c *Cart) AddToCart(ctx context.Context, user int64, sku uint32, count uint16) error {
	op := "Cart.AddToCart"

	stocks, err := c.lomsClient.Stocks(ctx, sku)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	counter := int64(count)
	for _, stock := range stocks {
		counter -= int64(stock.Count)
		if counter <= 0 {
			break
		}
	}
	if counter > 0 {
		return fmt.Errorf("%s: %w", op, models.ErrInsufficientStocks)
	}
	return nil
}
