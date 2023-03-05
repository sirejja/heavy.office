package models

type Item struct {
	SKU   uint32
	Count uint32
}

type Stock struct {
	WarehouseID int64
	Count       uint64
}

type Order struct {
	User   int64
	Status string
	Items  []Item
}

type Warehouse struct{}
