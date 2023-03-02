package orders

import (
	"context"
	"fmt"
	"route256/loms/internal/repo/order_repo"
)

func (o *Order) ListOrder(ctx context.Context, orderID int64) (*order_repo.Order, error) {
	op := "Order.ListOrder"

	order, err := o.ordersRepo.ListOrder(orderID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return order, nil
}
