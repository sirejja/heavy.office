package products

import (
	"context"
	"route256/checkout/internal/models"
	product_service "route256/product_service/pkg/v1/api"

	"google.golang.org/grpc"
)

type IProductServiceClient interface {
	GetProduct(ctx context.Context, sku uint32) (*models.ProductAttrs, error)
}

type Client struct {
	client product_service.ProductServiceClient
	token  string
}

func New(conn *grpc.ClientConn, token string) (*Client, error) {
	return &Client{
		token: token,

		client: product_service.NewProductServiceClient(conn),
	}, nil
}
