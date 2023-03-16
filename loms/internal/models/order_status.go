package models

type OrderStatus string

type IOrderStatus interface {
	ToString(status OrderStatus) string
}

const (
	OrderStatusNew         OrderStatus = "new"
	OrderStatusWaitPayment OrderStatus = "awaiting payment"
	OrderStatusFailed      OrderStatus = "failed"
	OrderStatusPayed       OrderStatus = "payed"
	OrderStatusCancelled   OrderStatus = "cancelled"
)

func (o OrderStatus) ToString() string {
	return string(o)
}
