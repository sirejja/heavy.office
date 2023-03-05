package products

import (
	"context"
	"fmt"
	"route256/checkout/internal/models"
	product_service "route256/product_service/pkg/v1/api"
)

func (c *Client) GetProduct(ctx context.Context, sku uint32) (*models.ProductAttrs, error) {
	op := "Client.GetProduct"
	request := product_service.GetProductRequest{Sku: sku, Token: c.token}
	response, err := c.client.GetProduct(ctx, &request)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	product := models.ProductAttrs{Name: response.GetName(), Price: response.GetPrice()}
	return &product, nil
}
