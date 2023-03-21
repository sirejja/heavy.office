package warehouse_orders_repo

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/pgxscan"
)

type FillOrderProductsIns struct {
}

func (w *warehouseOrdersRepo) FillOrderProducts(ctx context.Context, orderID uint64, warehouseID uint64, count uint32) (uint64, error) {
	op := "WarehouseOrdersRepo.FillOrderProducts"
	db := w.db.GetQueryEngine(ctx)

	query := sq.Insert(w.name).
		Columns("order_id, warehouse_id, cnt").
		Values(orderID, warehouseID, count).
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
