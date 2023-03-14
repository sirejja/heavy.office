package warehouse_orders_repo

import (
	"context"
	"fmt"
	"route256/loms/internal/models"

	sq "github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/pgxscan"
)

type FillOrderProductsIns struct {
	OrderID     uint64
	WarehouseID uint64
	Count       uint32
}

func (w *warehouseOrdersRepo) FillOrderProducts(ctx context.Context, ins *FillOrderProductsIns) (uint64, error) {
	op := "WarehouseOrdersRepo.FillOrderProducts"
	db := w.db.GetQueryEngine(ctx)

	if ins == nil {
		return 0, fmt.Errorf("%s: %w", op, models.ErrNoDataProvided)
	}

	query := sq.Insert(w.name).
		Columns("order_id, warehouse_id, cnt").
		Values(ins.OrderID, ins.WarehouseID, ins.Count).
		Suffix("RETURNING \"id\"").
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
