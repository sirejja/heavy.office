package warehouse

import (
	"context"
	"fmt"
	"route256/loms/internal/models"
	"route256/loms/internal/repositories/warehouse_repo"
)

func (w *Warehouse) Stocks(ctx context.Context, SKU uint32) ([]models.Stock, error) {
	op := "Warehouse.Stocks"

	stocks, err := w.warehouseRepo.GetStocks(ctx, &warehouse_repo.GetStocksFilter{SKU: SKU, IsDeleted: false})
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	domainStock := make([]models.Stock, 0)
	for _, stock := range stocks {
		domainStock = append(domainStock, models.Stock{WarehouseID: stock.ID, Count: stock.Count})
	}

	return domainStock, nil
}
