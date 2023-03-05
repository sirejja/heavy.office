package v1

import (
	"context"
	"fmt"
	"route256/checkout/internal/models"
	desc "route256/checkout/pkg/v1/api"
)

func ValidateListCart(r *desc.ListCartRequest) error {
	if r.GetUser() == 0 {
		return models.ErrEmptyUser
	}
	return nil
}

func (s *Implementation) ListCart(ctx context.Context, req *desc.ListCartRequest) (*desc.ListCartResponse, error) {
	op := "Implementation.ListCart"

	if err := ValidateListCart(req); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

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
