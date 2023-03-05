package orders

import (
	"context"
	"fmt"
	"route256/loms/internal/models"
)

func (o *Order) ListOrder(ctx context.Context, orderID int64) (*models.Order, error) {
	op := "Order.ListOrder"

	order, err := o.ordersRepo.ListOrder(orderID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return order, nil
}
