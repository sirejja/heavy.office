package warehouse

import "route256/loms/pkg/models"

type Warehouse struct{}

type Repo interface {
	BookProducts(user int64, items []models.Item) (*uint64, error)
	ListOrder(orderID int64) (*Order, error)
	PayedOrder(orderID int64) error
	CancellOrder(orderID int64) error
	GetStocks(SKU uint32) (*[]models.Stock, error)
}

func New() *Warehouse {
	return &Warehouse{}
}

func (w Warehouse) BookProducts(user int64, items []models.Item) (*uint64, error) {
	var orderId uint64 = 555
	return &orderId, nil
}

func (w Warehouse) ListOrder(orderID int64) (*Order, error) {
	order := Order{User: 111, Status: models.NewOrderStatus, Items: []models.Item{{111, 10}}}
	return &order, nil
}

func (w Warehouse) PayedOrder(orderID int64) error {
	return nil
}

func (w Warehouse) CancellOrder(orderID int64) error {
	return nil
}

func (w Warehouse) GetStocks(SKU uint32) (*[]models.Stock, error) {
	return &[]models.Stock{{WarehouseID: 111, Count: 11}}, nil
}

type Order struct {
	User   uint64             `json:"user"`
	Status models.OrderStatus `json:"status"`
	Items  []models.Item      `json:"items"`
}
