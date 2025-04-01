package model

import (
	"coffee-shop/internal/service"
)

type OrderItems struct {
	OrderID   int
	ProductID int
	Quantity  int
}

func (r *OrderItems) Validate() error {
	switch {
	case r.OrderID <= 0:
		return service.ErrNotValidOrderID
	case r.ProductID <= 0:
		return service.ErrNotValidMenuID
	case r.Quantity <= 0:
		return service.ErrNotValidQuantity
	default:
		return nil
	}
}
