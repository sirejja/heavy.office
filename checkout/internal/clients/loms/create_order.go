package loms

import (
	"context"
	"fmt"
	"route256/libs/client_wrapper"
)

type CreateOrderRequest struct {
	User  int64 `json:"user"`
	Items []Items
}

type Items struct {
	SKU   uint32 `json:"sku"`
	Count uint16 `json:"count"`
}

type CreateOrderResponse struct {
	OrderID uint64 `json:"orderID"`
}

func (c *Client) CreateOrder(ctx context.Context, user int64, items []Items) (*uint64, error) {
	op := "Client.CreateOrder"

	request := CreateOrderRequest{User: user, Items: items}
	response, err := client_wrapper.Post[CreateOrderRequest, CreateOrderResponse](ctx, c.urlCreateOrder, request)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return &response.OrderID, nil
}
