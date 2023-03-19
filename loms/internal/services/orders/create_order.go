package orders

import (
	"context"
	"fmt"
	"route256/loms/internal/models"
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
	op := "Order.processOrderCreation"

	productsToReserve, err := o.checkProductsForOrderCreation(ctx, items)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	orderID, err := o.ordersRepo.CreateOrder(ctx, user, models.OrderStatusNew.ToString())
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	err = o.reserveOrderProducts(ctx, productsToReserve, orderID)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}
	_, err = o.ordersRepo.UpdateOrderStatus(ctx, int64(orderID), models.OrderStatusWaitPayment)
	if err != nil {
		_, errStatus := o.ordersRepo.UpdateOrderStatus(ctx, int64(orderID), models.OrderStatusFailed)
		if errStatus != nil {
			return 0, fmt.Errorf("%s: %w", op, errStatus)
		}
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	err = o.changeStockAfterReserve(ctx, productsToReserve)
	if err != nil {
		_, err = o.ordersRepo.UpdateOrderStatus(ctx, int64(orderID), models.OrderStatusFailed)
		if err != nil {
			return 0, fmt.Errorf("%s: %w", op, err)
		}
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return int64(orderID), nil
}

func (o *Order) checkProductsForOrderCreation(ctx context.Context, items []models.Item) ([]models.ProductToReserve, error) {
	op := "Order.checkProductsForOrderCreation"

	productsToReserve := make([]models.ProductToReserve, 0)
	for _, item := range items {
		stocks, err := o.warehouseRepo.GetStocks(ctx, item.SKU)
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
				// К резерву часть стока
				countToReserve -= int32(item.Count)
				countReserved = int32(item.Count)
			}
			productsToReserve = append(productsToReserve, models.ProductToReserve{WarehouseID: stock.WarehouseID, Count: countReserved})
		}
		if countToReserve != 0 {
			return nil, models.ErrInsufficientStocks
		}
	}

	return productsToReserve, nil
}

func (o *Order) reserveOrderProducts(ctx context.Context, productsToReserve []models.ProductToReserve, orderID uint64) error {
	op := "Order.reserveOrderProducts"

	for _, product := range productsToReserve {
		_, err := o.warehouseOrdersRepo.FillOrderProducts(ctx, orderID, product.WarehouseID, uint32(product.Count))
		if err != nil {
			return fmt.Errorf("%s: %w", op, err)
		}
	}

	return nil
}

func (o *Order) changeStockAfterReserve(ctx context.Context, productsToReserve []models.ProductToReserve) error {
	op := "Order.changeStockAfterReserve"

	for _, product := range productsToReserve {
		_, err := o.warehouseRepo.ChangeStocks(ctx, product.WarehouseID, -product.Count)
		if err != nil {
			return fmt.Errorf("%s: %w", op, err)
		}
	}

	return nil
}
