package warehouse_repo

import (
	"context"
	"fmt"
	"route256/loms/internal/models"

	sq "github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/pgxscan"
)

type UpdateStocksFilter struct {
	SKU uint64
	ID  uint64
}

type UpdateStocksData struct {
	StockDiff int32
}

func (w *warehouseRepo) UpdateStocks(ctx context.Context, filter *UpdateStocksFilter, data *UpdateStocksData) (uint64, error) {
	op := "WarehouseRepo.UpdateStocks"

	if filter == nil {
		return 0, fmt.Errorf("%s: %w", op, models.ErrNoFiltersProvided)
	}
	if data == nil {
		return 0, fmt.Errorf("%s: %w", op, models.ErrNoDataProvided)
	}
	if data.StockDiff == 0 {
		return 0, fmt.Errorf("%s: %w", op, models.ErrNoDataProvided)
	}

	query := sq.Update(w.name).
		Suffix("RETURNING id").
		PlaceholderFormat(sq.Dollar).
		Set("updated_at", sq.Expr("current_timestamp")).
		Set("stock", sq.Expr("stock + ?", data.StockDiff))

	if filter.SKU != 0 {
		query = query.Where(sq.Eq{"id": filter.ID})
	}
	if filter.SKU != 0 {
		query = query.Where(sq.Eq{"sku": filter.SKU})
	}

	sql, args, err := query.ToSql()
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	row := w.db.QueryRow(ctx, sql, args...)

	var id uint64
	if err = row.Scan(&id); err != nil {
		if pgxscan.NotFound(err) {
			return 0, nil
		}
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return id, nil
}
