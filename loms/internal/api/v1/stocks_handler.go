package v1

import (
	"context"
	"fmt"
	"log"
	desc "route256/loms/pkg/grpc/server"
)

func (s *Server) Stocks(ctx context.Context, request *desc.StocksRequest) (*desc.StocksResponse, error) {
	op := "Server.Stocks"
	log.Printf("stocks_handler: %+v", request)

	stocks, err := s.warehouse.Stocks(ctx, request.GetSku())
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	stocksResp := make([]*desc.Stock, len(*stocks))
	for _, stock := range *stocks {
		stocksResp = append(stocksResp, &desc.Stock{Count: stock.Count, WarehouseID: stock.WarehouseID})
	}
	return &desc.StocksResponse{Stocks: stocksResp}, nil
}
