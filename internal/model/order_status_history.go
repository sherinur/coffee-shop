package model

type OrderStatusHistory struct {
	ID      int
	OrderID int
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
