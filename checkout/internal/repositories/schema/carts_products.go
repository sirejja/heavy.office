package schema

import (
	"database/sql"
)

type CartProductsSchema struct {
	ID        uint64       `db:"id"`
	CartID    uint64       `db:"cart_id"`
	SKU       uint32       `db:"sku"`
	Count     uint32       `db:"cnt"`
	CreatedAt sql.NullTime `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
	DeletedAt sql.NullTime `db:"deleted_at"`
}
