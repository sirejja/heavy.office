package v1

import (
	"context"
	"fmt"
	"log"
	desc "route256/checkout/pkg/grpc/server"
)

func (s *Server) ListCart(ctx context.Context, req *desc.ListCartRequest) (*desc.ListCartResponse, error) {
	op := "Server.ListCart"
	log.Printf("list_cart_handler: %+v", req)

	products, totalPrice, err := s.model.ListCart(ctx, req.GetUser())
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	var response desc.ListCartResponse
	for _, product := range products {
		response.Items = append(response.Items, &desc.CartItem{
			Sku:   product.SKU,
			Count: product.Count,
			Name:  product.Name,
			Price: product.Price,
		})
	}
	response.TotalPrice = totalPrice
	return &response, nil
}
