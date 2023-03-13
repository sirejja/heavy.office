package orders

import (
	"context"
	"fmt"
)

func (o *Order) PayedOrder(ctx context.Context, orderID int64) error {
	op := "Order.PayedOrder"

	_, err := o.ordersRepo.PayedOrder(ctx, orderID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}
