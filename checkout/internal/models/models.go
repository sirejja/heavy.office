package models

type CartProduct struct {
	SKU   uint32
	Count uint32
	Name  string
	Price uint32
}

type Stock struct {
	WarehouseID int64
	Count       uint64
}

type Item struct {
	SKU   uint32
	Count uint32
}

type ProductAttrs struct {
	Name  string
	Price uint32
}
