package transactor

import (
	"context"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/multierr"
)

type IQueryEngine interface {
	Query(ctx context.Context, query string, args ...interface{}) (pgx.Rows, error)
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
}

type IQueryEngineProvider interface {
	GetQueryEngine(ctx context.Context) IQueryEngine
}
type ITransactor interface {
	RunRepeatableRead(ctx context.Context, fx func(ctxTX context.Context) error) error
}

type transactKey string

const ctxTransactKey = transactKey("transact")

type TransactionManager struct {
	pool   *pgxpool.Pool
	ctxKey transactKey
}

func New(pool *pgxpool.Pool) *TransactionManager {
	return &TransactionManager{
		pool:   pool,
		ctxKey: ctxTransactKey,
	}
}

func (tm *TransactionManager) RunRepeatableRead(ctx context.Context, fx func(ctxTX context.Context) error) error {
	tx, err := tm.pool.BeginTx(ctx,
		pgx.TxOptions{
			IsoLevel: pgx.RepeatableRead,
		})
	if err != nil {
		return err
	}

	if err = fx(context.WithValue(ctx, tm.ctxKey, tx)); err != nil {
		return multierr.Combine(err, tx.Rollback(ctx))
	}

	if err = tx.Commit(ctx); err != nil {
		return multierr.Combine(err, tx.Rollback(ctx))
	}

	return nil
}

func (tm *TransactionManager) GetQueryEngine(ctx context.Context) IQueryEngine {
	tx, ok := ctx.Value(tm.ctxKey).(IQueryEngine)
	if ok && tx != nil {
		return tx
	}

	return tm.pool
}
