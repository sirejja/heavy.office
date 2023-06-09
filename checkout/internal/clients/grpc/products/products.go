package products

//go:generate sh -c "rm -rf mocks && mkdir -p mocks"
//go:generate minimock -i IProductServiceClient -o ./mocks/ -s "_minimock.go"

import (
	"context"
	"route256/checkout/internal/models"
	"route256/libs/cache/inmemory"
	product_service "route256/product_service/pkg/v1/api"

	"google.golang.org/grpc"
)

type IProductServiceClient interface {
	GetProduct(ctx context.Context, Sku uint32) (*models.ProductAttrs, error)
	GetProductCached(Sku uint32) (*models.ProductAttrs, error)
}

type Client struct {
	client product_service.ProductServiceClient
	token  string
	cache  inmemory.ICache
}

func New(conn *grpc.ClientConn, token string, cache inmemory.ICache) (*Client, error) {
	return &Client{
		token: token,

		client: product_service.NewProductServiceClient(conn),
		cache:  cache,
	}, nil
}
