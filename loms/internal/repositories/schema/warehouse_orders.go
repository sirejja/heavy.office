package schema

import (
	"database/sql"
	"time"
)

type WarehouseOrders struct {
	Id          uint64       `db:"id"`
	WarehouseID uint64       `db:"warehouse_id"`
	OrderID     uint64       `db:"order_id"`
	Count       uint32       `db:"cnt"`
	CreatedAt   time.Time    `db:"created_at"`
	UpdatedAt   time.Time    `db:"updated_at"`
	DeletedAt   sql.NullTime `db:"deleted_at"`
}

type WarehouseOrdersList struct {
	Id          uint64 `db:"id"`
	WarehouseID uint64 `db:"warehouse_id"`
	Count       uint32 `db:"cnt"`
	Status      string `db:"status"`
	UserID      uint64 `db:"user_id"`
	SKU         uint64 `db:"sku"`
}
