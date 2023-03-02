package loms

import (
	loms_service "route256/checkout/internal/grpc/clients/loms"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	client loms_service.LomsClient
}

func New(url string) (*Client, error) {
	conn, err := grpc.Dial(url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return &Client{
		client: loms_service.NewLomsClient(conn),
	}, nil
}
