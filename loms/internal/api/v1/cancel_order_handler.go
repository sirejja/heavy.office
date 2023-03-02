package v1

import (
	"context"
	"fmt"
	"log"
	desc "route256/loms/pkg/grpc/server"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) CancelOrder(ctx context.Context, request *desc.CancelOrderRequest) (*emptypb.Empty, error) {
	op := "Server.CancelOrder"
	log.Printf("payed_order_handler: %+v", request)

	err := s.orders.CancelOrder(ctx, request.GetOrderID())
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return &emptypb.Empty{}, nil
}
