package warehouse_repo

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/pgxscan"
)

type UpdateStocksFilter struct {
	SKU uint64
	ID  uint64
}

func (w *warehouseRepo) ChangeStocks(ctx context.Context, warehouseID uint64, stockDiff int32) (uint64, error) {
	op := "WarehouseRepo.ChangeStocks"
	db := w.db.GetQueryEngine(ctx)

	query := sq.Update(w.name).
		Set("updated_at", sq.Expr("current_timestamp")).
		Set("stock", sq.Expr("stock + ?", stockDiff)).
		Where(sq.Eq{"id": warehouseID}).
		Suffix("RETURNING id").
		PlaceholderFormat(sq.Dollar)

	sql, args, err := query.ToSql()
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	var id uint64
	if err = pgxscan.Get(ctx, db, &id, sql, args...); err != nil {
		if pgxscan.NotFound(err) {
			return 0, nil
		}
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return id, nil
}
