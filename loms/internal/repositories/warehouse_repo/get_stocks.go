package warehouse_repo

import (
	"context"
	"fmt"
	"route256/loms/internal/models"
	"route256/loms/internal/repositories/schema"

	sq "github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/pgxscan"
)

type GetStocksFilter struct {
	SKU       uint32
	IsDeleted bool
}

func (w *warehouseRepo) GetStocks(ctx context.Context, filter *GetStocksFilter) ([]*schema.Stock, error) {
	op := "WarehouseRepo.GetStocks"

	if filter == nil {
		return nil, fmt.Errorf("%s: %w", op, models.ErrNoFiltersProvided)
	}

	query := sq.Select("id, stock, sku").
		From(w.name).PlaceholderFormat(sq.Dollar)

	if filter.SKU != 0 {
		query = query.Where(sq.Eq{"sku": filter.SKU})
	}
	if !filter.IsDeleted {
		query = query.Where(sq.Eq{"deleted_at": nil})
	} else {
		query = query.Where(sq.Expr("deleted_at is not null"))
	}

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	var stocks []*schema.Stock
	if err = pgxscan.Select(ctx, w.db, &stocks, sql, args...); err != nil {
		if pgxscan.NotFound(err) {
			return nil, nil
		}
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return stocks, nil
}
