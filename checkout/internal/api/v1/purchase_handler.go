package v1

import (
	"context"
	"fmt"
	"log"
	"route256/checkout/internal/models"
	desc "route256/checkout/pkg/v1/api"
)

func ValidatePurchase(r *desc.PurchaseRequest) error {
	if r.GetUser() == 0 {
		return models.ErrEmptyUser
	}
	return nil
}

func (s *Implementation) Purchase(ctx context.Context, req *desc.PurchaseRequest) (*desc.PurchaseResponse, error) {
	op := "Implementation.Purchase"
	log.Printf("purchase_handler: %+v", req)

	if err := ValidatePurchase(req); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	orderID, err := s.model.PurchaseCart(ctx, req.User)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &desc.PurchaseResponse{OrderID: orderID}, nil
}
