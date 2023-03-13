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

	// TODO make transactor usage

	// check products
	productsToReserve := make([]models.ProductToReserve, 0)
	for _, item := range items {
		stocks, err := o.warehouseRepo.GetStocks(ctx, &warehouse_repo.GetStocksFilter{SKU: item.SKU, IsDeleted: false})
		if err != nil {
			return 0, fmt.Errorf("%s: %w", op, err)
		}

		countToReserve := int32(item.Count)
		var countReserved int32
		for _, stock := range stocks {
			if cnt := int32(stock.Count); cnt < countToReserve {
				countToReserve -= int32(stock.Count)
				countReserved = int32(stock.Count)
			} else if cnt >= countToReserve {
				countToReserve -= int32(item.Count)
				countReserved = int32(item.Count)
			} else {
				continue
			}
			productsToReserve = append(productsToReserve, models.ProductToReserve{WarehouseID: stock.ID, Count: countReserved})
		}
		if countToReserve != 0 {
			return 0, models.ErrInsufficientStocks
		}
	}

	// create order
	// TODO switch to enum
	orderID, err := o.ordersRepo.CreateOrder(ctx, &order_repo.CreateOrderIns{UserID: user, Status: "new"})
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	// reserve products
	// TODO maybe batch ins
	for _, product := range productsToReserve {
		_, err = o.warehouseOrdersRepo.FillOrderProducts(
			ctx,
			&warehouse_orders_repo.FillOrderProductsIns{
				OrderID:     orderID,
				WarehouseID: product.WarehouseID,
				Count:       uint32(product.Count),
			})
		if err != nil {
			return 0, fmt.Errorf("%s: %w", op, err)
		}
	}

	// decrease stocks
	// TODO maybe batch update
	for _, product := range productsToReserve {
		_, err = o.warehouseRepo.UpdateStocks(
			ctx,
			&warehouse_repo.UpdateStocksFilter{ID: product.WarehouseID},
			&warehouse_repo.UpdateStocksData{StockDiff: -product.Count},
		)
		if err != nil {
			return 0, fmt.Errorf("%s: %w", op, err)
		}
	}

	return int64(orderID), nil
}
