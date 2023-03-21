package order_repo

import (
	"context"
	"fmt"
	"route256/loms/internal/models"
	"route256/loms/internal/repositories/schema"

	sq "github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/pgxscan"
)

func (o *OrderRepo) ListOrderStacked(ctx context.Context, orderID uint64) ([]models.StackedOrder, error) {
	op := "OrderRepo.ListOrderStacked"
	db := o.db.GetQueryEngine(ctx)

	query := sq.
		Select("sum(wo.cnt) as cnt ,w.sku").
		From(fmt.Sprintf("%s o", o.name)).
		LeftJoin("warehouse_orders wo on wo.order_id=o.id").
		LeftJoin("warehouse w on w.id=wo.warehouse_id").
		Where(sq.Eq{"o.id": orderID}).
		Where(sq.Eq{"o.cancelled_at": nil}).
		Where(sq.Eq{"wo.deleted_at": nil}).
		Where(sq.Eq{"w.deleted_at": nil}).
		GroupBy("w.sku").
		PlaceholderFormat(sq.Dollar)

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	var listOrder []*schema.ListOrderStackedSchema
	if err = pgxscan.Select(ctx, db, &listOrder, sql, args...); err != nil {
		if pgxscan.NotFound(err) {
			return nil, nil
		}
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	result := make([]models.StackedOrder, 0, len(listOrder))
	for _, product := range listOrder {
		result = append(result, models.StackedOrder{Count: int32(product.Count), SKU: product.SKU})
	}

	return result, nil
}
