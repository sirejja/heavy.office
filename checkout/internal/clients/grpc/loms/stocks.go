package loms

import (
	"context"
	"fmt"
	"route256/checkout/internal/models"
	loms_service "route256/loms/pkg/v1/api"
)

func (c *Client) Stocks(ctx context.Context, sku uint32) ([]models.Stock, error) {
	op := "Client.Stocks"

	response, err := c.client.Stocks(ctx, &loms_service.StocksRequest{Sku: sku})
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	stocks := make([]models.Stock, 0, len(response.Stocks))
	for _, stock := range response.GetStocks() {
		stocks = append(stocks, models.Stock{WarehouseID: stock.GetWarehouseID(), Count: stock.GetCount()})
	}

	return stocks, nil
}
