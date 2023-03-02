package v1

import (
	"context"
	"fmt"
	"log"
	desc "route256/loms/pkg/grpc/server"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) OrderPayed(ctx context.Context, request *desc.OrderPayedRequest) (*emptypb.Empty, error) {
	op := "Handler.Handle"
	log.Printf("payed_order_handler: %+v", request)

	err := s.orders.PayedOrder(ctx, request.OrderID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return &emptypb.Empty{}, nil
}
