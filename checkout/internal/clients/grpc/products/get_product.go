package products

import (
	"context"
	"fmt"
	"route256/checkout/internal/models"
	product_service "route256/product_service/pkg/v1/api"
)

func (c *Client) GetProduct(ctx context.Context, Sku uint32) (*models.ProductAttrs, error) {
	op := "Client.GetProduct"

	response, err := c.client.GetProduct(ctx, &product_service.GetProductRequest{Sku: Sku, Token: c.token})
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	product := models.ProductAttrs{Name: response.GetName(), Price: response.GetPrice()}

	return &product, nil
}
