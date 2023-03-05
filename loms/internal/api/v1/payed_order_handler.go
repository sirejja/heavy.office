package v1

import (
	"context"
	"fmt"
	"route256/loms/internal/models"
	desc "route256/loms/pkg/v1/api"

	"google.golang.org/protobuf/types/known/emptypb"
)

func ValidateOrderPayed(r *desc.OrderPayedRequest) error {
	if r.GetOrderID() == 0 {
		return models.ErrEmptyOrderID
	}
	return nil
}

func (s *Implementation) OrderPayed(ctx context.Context, req *desc.OrderPayedRequest) (*emptypb.Empty, error) {
	op := "Implementation.OrderPayed"

	if err := ValidateOrderPayed(req); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	err := s.orders.PayedOrder(ctx, req.OrderID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &emptypb.Empty{}, nil
}
