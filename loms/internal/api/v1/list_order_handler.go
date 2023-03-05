package v1

import (
	"context"
	"fmt"
	"route256/loms/internal/models"
	desc "route256/loms/pkg/v1/api"
)

func ValidateListOrder(r *desc.ListOrderRequest) error {
	if r.GetOrderID() == 0 {
		return models.ErrEmptyOrderID
	}
	return nil
}

func (s *Implementation) ListOrder(ctx context.Context, req *desc.ListOrderRequest) (*desc.ListOrderResponse, error) {
	op := "Implementation.ListOrder"

	if err := ValidateListOrder(req); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	order, err := s.orders.ListOrder(ctx, req.GetOrderID())
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	items := make([]*desc.Item, 0, len(order.Items))
	for _, item := range order.Items {
		items = append(items, &desc.Item{Sku: item.SKU, Count: item.Count})
	}
	return &desc.ListOrderResponse{User: order.User, Status: order.Status, Items: items}, nil
}
