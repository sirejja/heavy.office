package orders

import (
	"context"
	"fmt"
	"route256/loms/internal/models"
	"route256/loms/internal/repositories/order_repo"
	"route256/loms/internal/repositories/warehouse_orders_repo"
	"route256/loms/internal/repositories/warehouse_repo"
)

func (o *Order) CreateOrder(ctx context.Context, user int64, items []models.Item) (int64, error) {
	op := "Order.CreateOrder"

	var OrderID int64
	var err error

	err = o.txManager.RunRepeatableRead(ctx, func(ctxTX context.Context) error {
		OrderID, err = o.processOrderCreation(ctxTX, user, items)
		if err != nil {
			return fmt.Errorf("%s: %w", op, err)
		}
		return nil
	})
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return OrderID, nil
}

func (o *Order) processOrderCreation(ctx context.Context, user int64, items []models.Item) (int64, error) {
	op := "Order.CreateOrder"

	productsToReserve, err := o.checkProductsForOrderCreation(ctx, items)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	orderID, err := o.ordersRepo.CreateOrder(ctx, &order_repo.CreateOrderIns{UserID: user, Status: "new"})
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	err = o.reserveOrderProducts(ctx, productsToReserve, orderID)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	err = o.updateStockAfterReserve(ctx, productsToReserve)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return int64(orderID), nil
}

func (o *Order) checkProductsForOrderCreation(ctx context.Context, items []models.Item) ([]models.ProductToReserve, error) {
	op := "Order.checkProductsForOrderCreation"

	productsToReserve := make([]models.ProductToReserve, 0)
	for _, item := range items {
		stocks, err := o.warehouseRepo.GetStocks(ctx, &warehouse_repo.GetStocksFilter{SKU: item.SKU, IsDeleted: false})
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		countToReserve := int32(item.Count)
		var countReserved int32
		for _, stock := range stocks {
			if cnt := int32(stock.Count); cnt < countToReserve {
				// К резерву весь свободный товар
				countToReserve -= int32(stock.Count)
				countReserved = int32(stock.Count)
			} else if cnt >= countToReserve {
				// К резерву часть товара, идем к следующей лоту
				countToReserve -= int32(item.Count)
				countReserved = int32(item.Count)
			} else {
				// Пустой сток
				continue
			}
			productsToReserve = append(productsToReserve, models.ProductToReserve{WarehouseID: stock.ID, Count: countReserved})
		}
		if countToReserve != 0 {
			return nil, models.ErrInsufficientStocks
		}
	}
	return productsToReserve, nil
}

func (o *Order) reserveOrderProducts(ctx context.Context, productsToReserve []models.ProductToReserve, orderID uint64) error {
	op := "Order.checkProductsForOrderCreation"

	for _, product := range productsToReserve {
		_, err := o.warehouseOrdersRepo.FillOrderProducts(
			ctx,
			&warehouse_orders_repo.FillOrderProductsIns{
				OrderID:     orderID,
				WarehouseID: product.WarehouseID,
				Count:       uint32(product.Count),
			})
		if err != nil {
			return fmt.Errorf("%s: %w", op, err)
		}
	}

	return nil
}

func (o *Order) updateStockAfterReserve(ctx context.Context, productsToReserve []models.ProductToReserve) error {
	op := "Order.checkProductsForOrderCreation"

	for _, product := range productsToReserve {
		_, err := o.warehouseRepo.UpdateStocks(
			ctx,
			&warehouse_repo.UpdateStocksFilter{ID: product.WarehouseID},
			&warehouse_repo.UpdateStocksData{StockDiff: -product.Count},
		)
		if err != nil {
			return fmt.Errorf("%s: %w", op, err)
		}
	}
	return nil
}
