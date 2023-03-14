package orders

import (
	"context"
	"fmt"
	"route256/loms/internal/models"
	"route256/loms/internal/repositories/order_repo"
	"route256/loms/internal/repositories/warehouse_repo"
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

	return nil
}

func (o *Order) processCancellationOrder(ctx context.Context, orderID int64) error {
	op := "Order.restoreProductsViaList"
	productsListRows, err := o.ordersRepo.ListOrder(ctx, &order_repo.ListOrderFilter{OrderID: uint64(orderID)})
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	productsToRestore := order_repo.FromSchemaToRestoringProducts(productsListRows)

	err = o.restoreProductsViaList(ctx, productsToRestore)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	_, err = o.ordersRepo.CancelOrder(ctx, orderID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (o *Order) restoreProductsViaList(ctx context.Context, productsToRestore []models.RestoringProducts) error {
	op := "Order.restoreProductsViaList"

	for _, product := range productsToRestore {
		_, err := o.warehouseRepo.UpdateStocks(
			ctx,
			&warehouse_repo.UpdateStocksFilter{ID: product.WarehouseID},
			&warehouse_repo.UpdateStocksData{StockDiff: product.Count},
		)
		if err != nil {
			return fmt.Errorf("%s: %w", op, err)
		}
	}

	return nil
}
