package loms

import (
	"context"
	"fmt"
	"route256/checkout/internal/models"
	loms_service "route256/loms/pkg/v1/api"
)

func (c *Client) CreateOrder(ctx context.Context, user int64, items []models.Item) (int64, error) {
	op := "Client.CreateOrder"
	requestItems := make([]*loms_service.Item, len(items))
	for _, item := range items {
		requestItems = append(requestItems, &loms_service.Item{
			Sku: item.SKU, Count: item.Count,
		})
	}
	response, err := c.client.CreateOrder(ctx, &loms_service.CreateOrderRequest{
		User:  user,
		Items: requestItems,
	})
	if err != nil {
		return int64(0), fmt.Errorf("%s: %w", op, err)
	}
	return response.OrderID, nil
}
