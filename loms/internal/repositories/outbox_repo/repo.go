package outbox_repo

import (
	"context"
	"route256/libs/transactor"
	"route256/loms/internal/repositories/schema"
)

type IOutboxRepo interface {
	ProcessOutboxTaskCreation(ctx context.Context, topic string, args interface{}) error
	GetTasks(ctx context.Context) ([]*schema.OutboxTask, error)
	DeleteTask(ctx context.Context, taskID uint64) (uint64, error)
}

var _ IOutboxRepo = (*outboxRepo)(nil)

type outboxRepo struct {
	db   *transactor.TransactionManager
	name string
}

func New(pool *transactor.TransactionManager) *outboxRepo {
	return &outboxRepo{
		db:   pool,
		name: "outbox",
	}
}
