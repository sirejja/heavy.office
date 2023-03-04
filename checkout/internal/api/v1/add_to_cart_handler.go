package v1

import (
	"context"
	"fmt"
	"log"
	"route256/checkout/internal/models"
	desc "route256/checkout/pkg/v1/api"

	"google.golang.org/protobuf/types/known/emptypb"
)

func ValidateAddToCart(r *desc.AddToCartRequest) error {
	if r.GetUser() == 0 {
		return models.ErrEmptyUser
	}
	if r.GetSku() == 0 {
		return models.ErrEmptySKU
	}
	if r.GetCount() == 0 {
		return models.ErrEmptyCount
	}
	return nil
}

func (s *Implementation) AddToCart(ctx context.Context, req *desc.AddToCartRequest) (*emptypb.Empty, error) {
	op := "Implementation.AddToCart"
	log.Printf("add_to_cart_handler: %+v", req)

	if err := ValidateAddToCart(req); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	err := s.model.AddToCart(ctx, req.GetUser(), req.GetSku(), req.GetCount())
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &emptypb.Empty{}, nil
}
