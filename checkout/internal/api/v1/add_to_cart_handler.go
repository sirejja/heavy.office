package v1

import (
	"context"
	"fmt"
	"log"
	desc "route256/checkout/pkg/grpc/server"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) AddToCart(ctx context.Context, req *desc.AddToCartRequest) (*emptypb.Empty, error) {
	op := "Server.AddToCart"
	log.Printf("add_to_cart_handler: %+v", req)

	err := s.model.AddToCart(ctx, req.GetUser(), req.GetSku(), req.GetCount())
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &emptypb.Empty{}, nil
}
