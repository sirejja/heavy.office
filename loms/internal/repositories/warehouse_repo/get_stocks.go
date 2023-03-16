package warehouse_repo

import (
	"context"
	"fmt"
	"route256/loms/internal/models"
	"route256/loms/internal/repositories/schema"

	sq "github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/pgxscan"
)

func (w *warehouseRepo) GetStocks(ctx context.Context, sku uint32) ([]models.Stock, error) {
	op := "WarehouseRepo.GetStocks"
	db := w.db.GetQueryEngine(ctx)

	query := sq.Select("id, stock, sku").
		From(w.name).
		Where(sq.Eq{"deleted_at": nil}).
		Where(sq.Eq{"sku": sku}).
		Where(sq.Gt{"stock": 0}).
		PlaceholderFormat(sq.Dollar)

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	var stocks []*schema.StocksSchema
	if err = pgxscan.Select(ctx, db, &stocks, sql, args...); err != nil {
		if pgxscan.NotFound(err) {
			return nil, nil
		}
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	warehouseStocks := make([]models.Stock, 0)
	for _, stock := range stocks {
		warehouseStocks = append(warehouseStocks, models.Stock{WarehouseID: stock.ID, Count: stock.Count})
	}

	return warehouseStocks, nil
}
