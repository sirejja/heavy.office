package orders

import (
	"context"
	"fmt"
	"route256/loms/internal/models"
)

func (o *Order) ListOrder(ctx context.Context, orderID int64) (*models.Order, error) {
	op := "Order.ListOrder"

	orderDetails, err := o.ordersRepo.GetOrderDetails(ctx, uint64(orderID))
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	listOrder, err := o.ordersRepo.ListOrderStacked(ctx, uint64(orderID))
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	if len(listOrder) == 0 {
		return nil, fmt.Errorf("%s: %w", op, models.ErrNotFound)
	}

	var order models.Order
	for _, product := range listOrder {
		order.Items = append(order.Items, models.Item{Count: uint32(product.Count), SKU: product.SKU})
	}
	order.User = orderDetails.UserID
	order.Status = orderDetails.Status

	return &order, nil
}
