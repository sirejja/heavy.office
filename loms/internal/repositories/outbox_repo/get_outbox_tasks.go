package outbox_repo

import (
	"context"
	"fmt"
	"route256/loms/internal/repositories/schema"

	sq "github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/pgxscan"
)

func (o *outboxRepo) GetTasks(ctx context.Context) ([]*schema.OutboxTask, error) {
	op := "outboxRepo.GetTasks"
	db := o.db.GetQueryEngine(ctx)

	query := sq.
		Select("id, topic, args").
		From(o.name).
		Where(sq.Eq{"deleted_at": nil}).
		OrderBy("created_at asc").
		PlaceholderFormat(sq.Dollar)

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	var tasks []*schema.OutboxTask
	if err = pgxscan.Select(ctx, db, &tasks, sql, args...); err != nil {
		if pgxscan.NotFound(err) {
			return nil, nil
		}
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return tasks, nil
}
