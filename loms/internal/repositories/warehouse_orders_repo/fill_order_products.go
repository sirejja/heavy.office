package warehouse_orders_repo

import (
	"context"
	"fmt"
	"route256/loms/internal/models"

	sq "github.com/Masterminds/squirrel"
)

type FillOrderProductsIns struct {
	OrderID     uint64
	WarehouseID uint64
	Count       uint32
}

func (w *warehouseOrdersRepo) FillOrderProducts(ctx context.Context, ins *FillOrderProductsIns) (uint64, error) {
	op := "WarehouseOrdersRepo.FillOrderProducts"

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

	if err = w.db.QueryRow(ctx, sql, args...).Scan(&id); err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return id, nil
}
