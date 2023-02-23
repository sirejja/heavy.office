package models

import "github.com/pkg/errors"

var (
	ErrEmptyOrder         = errors.New("empty order")
	ErrEmptySKU           = errors.New("empty sku")
	ErrEmptyUser          = errors.New("empty user")
	ErrEmptyCount         = errors.New("empty count")
	ErrInsufficientStocks = errors.New("insufficient stocks_handler")
	ErrSomeErr            = errors.New("mock error")
)

type Item struct {
	SKU   uint32 `json:"sku"`
	Count uint16 `json:"count"`
}

type Stock struct {
	WarehouseID int64  `json:"warehouseID"`
	Count       uint64 `json:"count"`
}

type OrderStatus string

const (
	NewOrderStatus             OrderStatus = "new"
	AwaitingPaymentOrderStatus OrderStatus = "awaiting payment"
	FailedOrderStatus          OrderStatus = "failed"
	PayedOrderStatus           OrderStatus = "payed"
	CancelledOrderStatus       OrderStatus = "cancelled"
)

func (s OrderStatus) String() string {
	switch s {
	case NewOrderStatus:
		return "new"
	case AwaitingPaymentOrderStatus:
		return "awaiting payment"
	case FailedOrderStatus:
		return "failed"
	case PayedOrderStatus:
		return "payed"
	case CancelledOrderStatus:
		return "cancelled"
	}
	return "unknown"
}

func (s OrderStatus) FromString(status string) OrderStatus {
	switch status {
	case string(NewOrderStatus):
		return NewOrderStatus
	case string(AwaitingPaymentOrderStatus):
		return AwaitingPaymentOrderStatus
	case string(FailedOrderStatus):
		return FailedOrderStatus
	case string(PayedOrderStatus):
		return PayedOrderStatus
	case string(CancelledOrderStatus):
		return CancelledOrderStatus
	}
	return "unknown"
}
