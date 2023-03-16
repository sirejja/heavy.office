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

type OrderDetails struct {
	UserID int64  `db:"user_id"`
	Status string `db:"status"`
}

type WarehouseOrdersListSchema struct {
	Id          uint64 `db:"id"`
	WarehouseID uint64 `db:"warehouse_id"`
	Count       uint32 `db:"cnt"`
	Status      string `db:"status"`
	UserID      uint64 `db:"user_id"`
	SKU         uint32 `db:"sku"`
}

type ListOrderStackedSchema struct {
	Count uint32 `db:"cnt"`
	SKU   uint32 `db:"sku"`
}
