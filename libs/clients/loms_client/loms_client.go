package loms_client

import (
	"context"
	"github.com/pkg/errors"
	"route256/libs/client_wrapper"
)

const (
	StocksPath      = "/stocks"
	CreateOrderPath = "/createOrder"
)

type LOMS interface {
	Stocks(ctx context.Context, sku uint32) ([]Stock, error)
	CreateOrder(ctx context.Context, user int64) error
}

type Client struct {
	url   string
	token string
}

func New(url string) *Client {
	return &Client{
		url: url,
	}
}

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
	request := StocksRequest{SKU: sku}
	response, err := client_wrapper.Post[StocksRequest, StocksResponse](ctx, c.url+StocksPath, request)
	if err != nil {
		return nil, errors.WithMessage(err, "marshaling json")
	}
	stocks := make([]Stock, 0, len(response.Stocks))
	for _, stock := range response.Stocks {
		stocks = append(stocks, Stock{
			WarehouseID: stock.WarehouseID,
			Count:       stock.Count,
		})
	}

	return stocks, nil
}

type CreateOrderRequest struct {
	User  int64 `json:"user"`
	Items []Items
}

type Items struct {
	SKU   uint32 `json:"sku"`
	Count uint16 `json:"count"`
}

type CreateOrderResponse struct {
	User int64 `json:"user"`
}

func (c *Client) CreateOrder(ctx context.Context, user int64) error {
	request := CreateOrderRequest{User: user, Items: []Items{{SKU: 1, Count: 1}}}
	_, err := client_wrapper.Post[CreateOrderRequest, CreateOrderResponse](ctx, c.url+CreateOrderPath, request)
	if err != nil {
		return errors.WithMessage(err, "marshaling json")
	}
	return nil
}
