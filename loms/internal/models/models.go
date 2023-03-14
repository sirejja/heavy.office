package models

type Item struct {
	SKU   uint32
	Count uint32
}

type Stock struct {
	WarehouseID uint64
	Count       uint32
}

type Order struct {
	User   int64
	Status string
	Items  []Item
}

type Warehouse struct{}

type ProductToReserve struct {
	WarehouseID uint64
	Count       int32
}

type RestoringProducts struct {
	WarehouseID uint64
	Count       int32
}
