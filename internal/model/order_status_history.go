package model

import "time"

type OrderStatusHistory struct {
	ID       int
	OrderID  int
	OpenedAt time.Time
	ClosedAt time.Time
}

func (r *OrderStatusHistory) Validate() error {
	switch {
	case r.ID <= 0:
		return ErrNotValidOrderID
	case r.OrderID <= 0:
		return ErrNotValidOrderID
	default:
		return nil
	}
}
