package v1

import (
	"context"
	"fmt"
	"log"
	desc "route256/checkout/pkg/grpc/server"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) DeleteFromCart(ctx context.Context, req *desc.DeleteFromCartRequest) (*emptypb.Empty, error) {
	op := "Server.DeleteFromCart"
	log.Printf("delete_from_cart_handler: %+v", req)

	if err := s.model.DeleteFromCart(ctx, req.GetUser(), req.GetSku(), req.GetCount()); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &emptypb.Empty{}, nil
}
