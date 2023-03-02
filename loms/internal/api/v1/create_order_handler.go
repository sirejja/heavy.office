package v1

import (
	"context"
	"fmt"
	"log"
	"route256/loms/internal/models"
	desc "route256/loms/pkg/grpc/server"
)

func (s *Server) CreateOrder(ctx context.Context, request *desc.CreateOrderRequest) (*desc.CreateOrderResponse, error) {
	op := "Server.CreateOrder"
	log.Printf("create_order_handler: %+v", request)
	items := make([]models.Item, len(request.GetItems()))
	for _, item := range request.GetItems() {
		items = append(items, models.Item{SKU: item.GetCount(), Count: item.GetCount()})
	}
	orderID, err := s.orders.CreateOrder(ctx, request.User, items)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return &desc.CreateOrderResponse{OrderID: orderID}, nil
}
