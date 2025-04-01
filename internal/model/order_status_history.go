package model

type OrderStatusHistory struct {
	ID       int
	OrderID  int
	OpenedAt string
	ClosedAt string
}

func (r *OrderStatusHistory) Validate() error {
	switch {
	case r.ID <= 0:
		return ErrNotValidOrderID
	case r.OrderID <= 0:
		return ErrNotValidOrderID
	case r.OpenedAt == "":
		return ErrNotValidStatusHistoryTime
	case r.ClosedAt == "":
		return ErrNotValidStatusHistoryTime
	default:
		return nil
	}
}
