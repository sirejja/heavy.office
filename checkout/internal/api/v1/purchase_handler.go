package v1

import (
	"context"
	"fmt"
	"log"
	desc "route256/checkout/pkg/grpc/server"
)

func (s *Server) Purchase(ctx context.Context, req *desc.PurchaseRequest) (*desc.PurchaseResponse, error) {
	op := "Server.Purchase"
	log.Printf("purchase_handler: %+v", req)

	orderID, err := s.model.PurchaseCart(ctx, req.User)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &desc.PurchaseResponse{OrderID: orderID}, nil
}
