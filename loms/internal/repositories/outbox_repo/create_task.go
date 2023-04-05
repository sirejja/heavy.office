package outbox_repo

import (
	"context"
	"encoding/json"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/pgxscan"
)

func (o *outboxRepo) ProcessOutboxTaskCreation(ctx context.Context, topic string, args interface{}) error {
	op := "outboxRepo.ProcessOutboxTaskCreation"

	b, err := json.Marshal(args)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	_, err = o.createTask(ctx, topic, b)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

func (o *outboxRepo) createTask(ctx context.Context, topic string, taskArgs []byte) (uint64, error) {
	op := "outboxRepo.CreateTask"
	db := o.db.GetQueryEngine(ctx)

	query := sq.Insert(o.name).
		Columns("topic, args").
		Values(topic, taskArgs).
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
