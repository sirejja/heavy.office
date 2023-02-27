package orders

import (
	"context"
	"fmt"
)

func (o *Order) PayedOrder(ctx context.Context, orderID int64) error {
	op := "Order.PayedOrder"

	err := o.ordersRepo.PayedOrder(orderID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}
