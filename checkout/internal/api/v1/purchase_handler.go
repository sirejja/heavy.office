package v1

import (
	"context"
	"fmt"
	"route256/checkout/internal/models"
	desc "route256/checkout/pkg/v1/api"
)

func ValidatePurchase(r *desc.PurchaseRequest) error {
	if r.GetUser() == 0 {
		fmt.Println(1111)
		return models.ErrEmptyUser
	}
	return nil
}

func (s *Implementation) Purchase(ctx context.Context, req *desc.PurchaseRequest) (*desc.PurchaseResponse, error) {
	op := "Implementation.Purchase"

	if err := ValidatePurchase(req); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	orderID, err := s.model.PurchaseCart(ctx, req.GetUser())
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &desc.PurchaseResponse{OrderID: orderID}, nil
}
