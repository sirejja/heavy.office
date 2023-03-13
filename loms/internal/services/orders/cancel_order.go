package orders

import (
	"context"
	"fmt"
	"route256/loms/internal/repositories/order_repo"
	"route256/loms/internal/repositories/warehouse_repo"
)

func (o *Order) CancelOrder(ctx context.Context, orderID int64) error {
	op := "Order.CancelOrder"
	// TODO make transactor usage

	productsToRestore, err := o.ordersRepo.ListOrder(ctx, &order_repo.ListOrderFilter{OrderID: uint64(orderID)})
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	// TODO maybe batch update
	for _, product := range productsToRestore {
		_, err = o.warehouseRepo.UpdateStocks(
			ctx,
			&warehouse_repo.UpdateStocksFilter{ID: product.WarehouseID},
			&warehouse_repo.UpdateStocksData{StockDiff: int32(product.Count)},
		)
		if err != nil {
			return fmt.Errorf("%s: %w", op, err)
		}
	}

	_, err = o.ordersRepo.CancelOrder(ctx, orderID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
