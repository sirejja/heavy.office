package orders

import (
	"context"
	"fmt"
	"route256/loms/internal/models"
	"route256/loms/internal/repositories/order_repo"
)

func (o *Order) ListOrder(ctx context.Context, orderID int64) (*models.Order, error) {
	op := "Order.ListOrder"

	orderDetails, err := o.ordersRepo.GetOrder(ctx, &order_repo.GetOrderFilter{Id: uint64(orderID)})
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	listOrder, err := o.ordersRepo.ListOrder(ctx, &order_repo.ListOrderFilter{OrderID: uint64(orderID)})
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	if len(listOrder) == 0 {
		return nil, fmt.Errorf("%s: %w", op, models.ErrNotFound)
	}

	SKUCount := make(map[uint64]uint32)
	for _, product := range listOrder {
		SKUCount[product.SKU] += product.Count
	}

	var order models.Order
	for SKU, Count := range SKUCount {
		order.Items = append(order.Items, models.Item{Count: Count, SKU: uint32(SKU)})
	}
	order.User = orderDetails.UserID
	order.Status = orderDetails.Status

	return &order, nil
}
