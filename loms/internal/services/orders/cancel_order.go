package orders

import (
	"context"
	"fmt"
)

func (o *Order) CancelOrder(ctx context.Context, orderID int64) error {
	op := "Order.CancelOrder"

	err := o.ordersRepo.CancelOrder(orderID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}
