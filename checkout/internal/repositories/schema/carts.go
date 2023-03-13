package schema

import (
	"database/sql"
	"time"
)

type CartSchema struct {
	Id          uint64       `db:"id"`
	UserId      int64        `db:"user_id"`
	CreatedAt   time.Time    `db:"created_at"`
	UpdatedAt   time.Time    `db:"updated_at"`
	DeletedAt   sql.NullTime `db:"deleted_at"`
	PurchasedAt sql.NullTime `db:"purchased_at"`
}
