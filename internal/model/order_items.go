package model

type OrderItems struct {
	OrderID   int
	ProductID int
	Quantity  int
}

func (r *OrderItems) Validate() error {
	switch {
	case r.OrderID <= 0:
		return ErrNotValidOrderID
	case r.ProductID <= 0:
		return ErrNotValidMenuID
	case r.Quantity <= 0:
		return ErrNotValidQuantity
	default:
		return nil
	}
}
