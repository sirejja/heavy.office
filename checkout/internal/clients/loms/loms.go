package loms

import (
	"context"
	"route256/checkout/internal/models"
	loms_service "route256/loms/pkg/v1/api"

	"google.golang.org/grpc"
)

type ILOMSClient interface {
	Stocks(ctx context.Context, sku uint32) ([]models.Stock, error)
	CreateOrder(ctx context.Context, user int64, items []models.Item) (int64, error)
}

type Client struct {
	client loms_service.LomsClient
}

func New(conn *grpc.ClientConn) (*Client, error) {
	return &Client{
		client: loms_service.NewLomsClient(conn),
	}, nil
}
