package v1

import (
	"context"
	"fmt"
	"log"
	"route256/loms/internal/models"
	desc "route256/loms/pkg/v1/api"
)

func ValidateCreateOrder(r *desc.CreateOrderRequest) error {
	if r.GetUser() == 0 {
		return models.ErrEmptyUser
	}
	for _, item := range r.GetItems() {
		if item.GetCount() == 0 {
			return models.ErrEmptyCount
		}
		if item.GetSku() == 0 {
			return models.ErrEmptySKU
		}
	}
	return nil
}

func (s *Server) CreateOrder(ctx context.Context, req *desc.CreateOrderRequest) (*desc.CreateOrderResponse, error) {
	op := "Server.CreateOrder"
	log.Printf("create_order_handler: %+v", req)

	if err := ValidateCreateOrder(req); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	items := make([]models.Item, len(req.GetItems()))
	for _, item := range req.GetItems() {
		items = append(items, models.Item{SKU: item.GetCount(), Count: item.GetCount()})
	}
	orderID, err := s.orders.CreateOrder(ctx, req.User, items)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return &desc.CreateOrderResponse{OrderID: orderID}, nil
}
