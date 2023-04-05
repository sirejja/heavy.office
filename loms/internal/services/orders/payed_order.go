package orders

import (
	"context"
	"fmt"
	"route256/loms/internal/kafka/outbox_producer"
	"route256/loms/internal/models"
)

func (o *Order) PayedOrder(ctx context.Context, orderID int64) error {
	op := "Order.PayedOrder"

	_, err := o.ordersRepo.UpdateOrderStatus(ctx, orderID, models.OrderStatusPayed)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	err = o.outboxRepo.ProcessOutboxTaskCreation(ctx, "", outbox_producer.OrderStatusMsg{ID: int64(orderID), Status: models.OrderStatusPayed.ToString()})
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
