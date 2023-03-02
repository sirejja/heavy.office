package orders

import (
	"context"
	"fmt"
	"route256/loms/internal/models"
)

func (o *Order) CreateOrder(ctx context.Context, user int64, items []models.Item) (*uint64, error) {
	op := "Order.CreateOrder"

	orderId, err := o.warehouseRepo.BookProducts(user, items)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return orderId, nil
}
