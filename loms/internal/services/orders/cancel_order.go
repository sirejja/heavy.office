package orders

import (
	"context"
	"fmt"
	"route256/loms/internal/kafka/outbox_producer"
	"route256/loms/internal/models"
)

func (o *Order) CancelOrder(ctx context.Context, orderID int64) error {
	op := "Order.CancelOrder"

	err := o.txManager.RunRepeatableRead(ctx, func(ctxTX context.Context) error {
		err := o.processCancellationOrder(ctxTX, orderID)
		if err != nil {
			return fmt.Errorf("%s: %w", op, err)
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	err = o.outboxRepo.ProcessOutboxTaskCreation(
		ctx,
		o.cfg.Kafka.Topics.OrderStatus,
		outbox_producer.OrderStatusMsg{ID: orderID, Status: models.OrderStatusCancelled.ToString()},
	)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (o *Order) processCancellationOrder(ctx context.Context, orderID int64) error {
	op := "Order.processCancellationOrder"

	productsToRestore, err := o.ordersRepo.ListOrder(ctx, uint64(orderID))
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	err = o.restoreProductsViaList(ctx, productsToRestore)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	_, err = o.ordersRepo.UpdateOrderStatus(ctx, orderID, models.OrderStatusCancelled)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (o *Order) restoreProductsViaList(ctx context.Context, productsToRestore []models.ListOrder) error {
	op := "Order.restoreProductsViaList"

	for _, product := range productsToRestore {
		_, err := o.warehouseRepo.ChangeStocks(ctx, product.WarehouseID, product.Count)
		if err != nil {
			return fmt.Errorf("%s: %w", op, err)
		}
	}

	return nil
}
