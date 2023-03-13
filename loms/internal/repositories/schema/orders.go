package schema

import (
	"database/sql"
	"time"
)

type OrdersSchema struct {
	Id          uint64       `db:"id"`
	UserID      int64        `db:"user_id"`
	Status      string       `db:"status"`
	CreatedAt   time.Time    `db:"created_at"`
	UpdatedAt   time.Time    `db:"updated_at"`
	CancelledAt sql.NullTime `db:"cancelled_at"`
}
