package products_client

import (
	"context"
	"github.com/pkg/errors"
	"route256/libs/client_wrapper"
)

const (
	GetProductPath = "/get_product"
)

type Client struct {
	url       string
	token     string
	urlStocks string
}

func New(url, token string) *Client {
	return &Client{
		url:   url,
		token: token,
	}
}

type GetProductRequest struct {
	SKU   uint32 `json:"sku"`
	Token string `json:"token"`
}

type Product struct {
	Name  string `json:"name"`
	Price uint32 `json:"price"`
}

type GetProductResponse struct {
	Product
}

func (c *Client) GetProduct(ctx context.Context, sku uint32) (*GetProductResponse, error) {
	request := GetProductRequest{SKU: sku, Token: c.token}
	response, err := client_wrapper.Post[GetProductRequest, GetProductResponse](ctx, c.url+GetProductPath, request)
	if err != nil {
		return nil, errors.WithMessage(err, "marshaling json")
	}

	return &response, nil
}

func (c *Client) ListSKUs(ctx context.Context, SKU uint32) error {

	return nil
}
