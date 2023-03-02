package warehouse

import (
	"context"
	"fmt"
	"route256/loms/internal/models"
)

func (w *Warehouse) Stocks(ctx context.Context, SKU uint32) (*[]models.Stock, error) {
	op := "Warehouse.Stocks"

	stocks, err := w.warehouseRepo.GetStocks(SKU)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return stocks, nil
}
