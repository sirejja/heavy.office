package products

import (
	"route256/checkout/internal/grpc/clients/product_service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	client product_service.ProductServiceClient
	token  string
}

func New(url string, token string) (*Client, error) {
	conn, err := grpc.Dial(url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return &Client{
		token: token,

		client: product_service.NewProductServiceClient(conn),
	}, nil
}
