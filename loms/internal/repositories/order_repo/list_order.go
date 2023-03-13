package order_repo

import (
	"context"
	"fmt"
	"route256/loms/internal/models"
	"route256/loms/internal/repositories/schema"

	sq "github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/pgxscan"
)

type ListOrderFilter struct {
	OrderID uint64
}

func (o *OrderRepo) ListOrder(ctx context.Context, filter *ListOrderFilter) ([]*schema.WarehouseOrdersList, error) {
	op := "OrderRepo.ListOrder"
	if filter == nil {
		return nil, fmt.Errorf("%s: %w", op, models.ErrNoFiltersProvided)
	}

	query := sq.
		Select("wo.id, w.id warehouse_id, wo.cnt, o.user_id, o.status, w.sku").
		From(fmt.Sprintf("%s o", o.name)).
		LeftJoin("warehouse_orders wo on wo.order_id=o.id").
		LeftJoin("warehouse w on w.id=wo.warehouse_id").
		Where(sq.Eq{"o.id": filter.OrderID}).
		Where(sq.Eq{"o.cancelled_at": nil}).
		Where(sq.Eq{"wo.deleted_at": nil}).
		PlaceholderFormat(sq.Dollar)

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	var listOrder []*schema.WarehouseOrdersList
	if err = pgxscan.Select(ctx, o.db, &listOrder, sql, args...); err != nil {
		if pgxscan.NotFound(err) {
			return nil, nil
		}
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return listOrder, nil
}
