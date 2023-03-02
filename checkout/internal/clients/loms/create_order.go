package loms

import (
	"context"
	"fmt"
	lomsservice "route256/checkout/internal/grpc/clients/loms"
	"route256/checkout/internal/models"
)

func (c *Client) CreateOrder(ctx context.Context, user int64, items []models.Items) (int64, error) {
	op := "Client.CreateOrder"
	requestItems := make([]*lomsservice.Item, len(items))
	for _, item := range items {
		requestItems = append(requestItems, &lomsservice.Item{
			Sku: item.SKU, Count: item.Count,
		})
	}
	response, err := c.client.CreateOrder(ctx, &lomsservice.CreateOrderRequest{
		User:  user,
		Items: requestItems,
	})
	if err != nil {
		return int64(0), fmt.Errorf("%s: %w", op, err)
	}
	return response.OrderID, nil
}
