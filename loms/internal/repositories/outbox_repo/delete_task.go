package outbox_repo

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/pgxscan"
)

func (o *outboxRepo) DeleteTask(ctx context.Context, taskID uint64) (uint64, error) {
	op := "outboxRepo.DeleteTask"
	db := o.db.GetQueryEngine(ctx)

	query := sq.Update(o.name).
		Set("deleted_at", sq.Expr("current_timestamp")).
		Where(sq.Eq{"id": taskID}).
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
