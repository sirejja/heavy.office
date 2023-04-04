package orders

import (
	"context"
	"fmt"
	"route256/loms/internal/kafka/order_sender"
	"route256/loms/internal/models"
)

func (o *Order) PayedOrder(ctx context.Context, orderID int64) error {
	op := "Order.PayedOrder"

	_, err := o.ordersRepo.UpdateOrderStatus(ctx, orderID, models.OrderStatusPayed)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if err = o.brokerSender.SendOrderOrderStatusEvent(order_sender.OrderStatusMsg{ID: orderID, Status: models.OrderStatusPayed}); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
