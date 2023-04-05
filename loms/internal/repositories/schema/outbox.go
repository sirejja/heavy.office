package schema

type OutboxTask struct {
	ID    uint64 `db:"id"`
	Topic string `db:"topic"`
	Args  []byte `db:"args"`
}
