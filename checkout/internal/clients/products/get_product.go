package products

import (
	"context"
	"fmt"
	"route256/libs/client_wrapper"
)

type GetProductRequest struct {
	SKU   uint32 `json:"sku"`
	Token string `json:"token"`
}

type GetProductResponse struct {
	Name  string `json:"name"`
	Price uint32 `json:"price"`
}

type Product struct {
	Name  string
	Price uint32
}

func (c *Client) GetProduct(ctx context.Context, sku uint32) (*Product, error) {
	op := "Client.GetProduct"
	request := GetProductRequest{SKU: sku, Token: c.token}
	response, err := client_wrapper.Post[GetProductRequest, GetProductResponse](ctx, c.urlGetProducts, request)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	product := Product(response)
	return &product, nil
}
