package model

import "coffee-shop/internal/service"

type OrderStatusHistory struct {
	ID       int
	OrderID  int
	OpenedAt string
	ClosedAt string
}

func (r *OrderStatusHistory) Validate() error {
	switch {
	case r.ID <= 0:
		return service.ErrNotValidOrderID
	case r.OrderID <= 0:
		return service.ErrNotValidOrderID
	case r.OpenedAt == "":
		return service.ErrNotValidStatusHistoryTime
	case r.ClosedAt == "":
		return service.ErrNotValidStatusHistoryTime
	default:
		return nil
	}
}
