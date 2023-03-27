package transactor

//go:generate sh -c "rm -rf mocks && mkdir -p mocks"
//go:generate minimock -i ITransactor -o ./mocks/ -s "_minimock.go"
//go:generate minimock -i IQueryEngine -o ./mocks/ -s "_minimock.go"
//go:generate minimock -i IQueryEngineProvider -o ./mocks/ -s "_minimock.go"
//go:generate minimock -i github.com/jackc/pgx/v4.Tx -o ./mocks/tx_minimock.go

import (
	"context"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"go.uber.org/multierr"
)

type IQueryEngine interface {
	Query(ctx context.Context, query string, args ...interface{}) (pgx.Rows, error)
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	BeginTx(ctx context.Context, txOptions pgx.TxOptions) (pgx.Tx, error)
}

type IQueryEngineProvider interface {
	GetQueryEngine(ctx context.Context) IQueryEngine
}
type ITransactor interface {
	RunRepeatableRead(ctx context.Context, fx func(ctxTX context.Context) error) error
}

type transactKey string

const TXKey = transactKey("transact")

type TransactionManager struct {
	pool   IQueryEngine
	ctxKey transactKey
}

func New(pool IQueryEngine) *TransactionManager {
	return &TransactionManager{
		pool:   pool,
		ctxKey: TXKey,
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
	ctxTX := context.WithValue(ctx, tm.ctxKey, tx)
	if err = fx(ctxTX); err != nil {
		return multierr.Combine(err, tx.Rollback(ctxTX))
	}

	if err = tx.Commit(ctxTX); err != nil {
		return multierr.Combine(err, tx.Rollback(ctxTX))
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
