package schema

import (
	"database/sql"
	"time"
)

type WarehouseSchema struct {
	Id        uint64       `db:"id"`
	SKU       uint32       `db:"sku"`
	Stock     uint32       `db:"stock"`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt time.Time    `db:"updated_at"`
	DeletedAt sql.NullTime `db:"deleted_at"`
}

type Stock struct {
	ID    uint64 `db:"id"`
	Count uint32 `db:"stock"`
	SKU   uint64 `db:"sku"`
}
