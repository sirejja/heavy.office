package loms

import (
	"context"
	"fmt"
	"route256/libs/client_wrapper"
)

type StocksRequest struct {
	SKU uint32 `json:"sku"`
}

type StocksItem struct {
	WarehouseID int64  `json:"warehouseID"`
	Count       uint64 `json:"count"`
}

type StocksResponse struct {
	Stocks []StocksItem `json:"stocks"`
}

type Stock struct {
	WarehouseID int64
	Count       uint64
}

func (c *Client) Stocks(ctx context.Context, sku uint32) ([]Stock, error) {
	op := "Client.Stocks"

	request := StocksRequest{SKU: sku}
	response, err := client_wrapper.Post[StocksRequest, StocksResponse](ctx, c.urlStocks, request)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	stocks := make([]Stock, 0, len(response.Stocks))
	for _, stock := range response.Stocks {
		stocks = append(stocks, Stock(stock))
	}

	return stocks, nil
}
