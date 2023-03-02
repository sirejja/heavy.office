package v1

import (
	"context"
	"fmt"
	"log"
	desc "route256/loms/pkg/grpc/server"
)

func (s *Server) ListOrder(ctx context.Context, request *desc.ListOrderRequest) (*desc.ListOrderResponse, error) {
	op := "Server.ListOrder"
	log.Printf("create_order_handler: %+v", request)

	order, err := s.orders.ListOrder(ctx, request.GetOrderID())
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	items := make([]*desc.Item, 0, len(order.Items))
	for _, item := range order.Items {
		items = append(items, &desc.Item{Sku: item.SKU, Count: item.Count})
	}
	return &desc.ListOrderResponse{User: order.User, Status: order.Status, Items: items}, nil
}
