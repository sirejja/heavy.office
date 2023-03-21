package v1

import (
	"context"
	"fmt"
	"route256/loms/internal/models"
	desc "route256/loms/pkg/v1/api"
)

func ValidateStocksRequest(r *desc.StocksRequest) error {
	if r.GetSku() == 0 {
		return models.ErrEmptySKU
	}
	return nil
}

func (s *Implementation) Stocks(ctx context.Context, req *desc.StocksRequest) (*desc.StocksResponse, error) {
	op := "Implementation.Stocks"

	if err := ValidateStocksRequest(req); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	stocks, err := s.warehouse.Stocks(ctx, req.GetSku())
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	stocksResp := make([]*desc.Stock, 0, len(stocks))
	for _, stock := range stocks {
		stocksResp = append(stocksResp, &desc.Stock{Count: uint64(stock.Count), WarehouseID: int64(stock.WarehouseID)})
	}

	return &desc.StocksResponse{Stocks: stocksResp}, nil
}
