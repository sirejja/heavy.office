package v1

import (
	"context"
	"fmt"
	"log"
	"route256/checkout/internal/models"
	desc "route256/checkout/pkg/v1/api"

	"google.golang.org/protobuf/types/known/emptypb"
)

func ValidateDeleteFromCart(r *desc.DeleteFromCartRequest) error {
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

func (s *Implementation) DeleteFromCart(ctx context.Context, req *desc.DeleteFromCartRequest) (*emptypb.Empty, error) {
	op := "Implementation.DeleteFromCart"
	log.Printf("delete_from_cart_handler: %+v", req)

	if err := ValidateDeleteFromCart(req); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	if err := s.model.DeleteFromCart(ctx, req.GetUser(), req.GetSku(), req.GetCount()); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &emptypb.Empty{}, nil
}
