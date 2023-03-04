package v1

import (
	"context"
	"fmt"
	"log"
	"route256/loms/internal/models"
	desc "route256/loms/pkg/v1/api"

	"google.golang.org/protobuf/types/known/emptypb"
)

func ValidatCancelOrder(r *desc.CancelOrderRequest) error {
	if r.GetOrderID() == 0 {
		return models.ErrEmptyOrderID
	}
	return nil
}

func (s *Server) CancelOrder(ctx context.Context, req *desc.CancelOrderRequest) (*emptypb.Empty, error) {
	op := "Server.CancelOrder"
	log.Printf("cancel_order_handler: %+v", req)

	if err := ValidatCancelOrder(req); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	err := s.orders.CancelOrder(ctx, req.GetOrderID())
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return &emptypb.Empty{}, nil
}
