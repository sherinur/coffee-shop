package model

import "coffee-shop/internal/service"

type InventoryTransactions struct {
	TransactionID  int
	IngredientId   int
	QuantityChange float64
	Reason         string
	CreatedAt      string
}

func (r *InventoryTransactions) Validate() error {
	switch {
	case r.TransactionID <= 0:
		return service.ErrInventoryItemNotFound
	case r.IngredientId <= 0:
		return service.ErrDuplicateMenuIngredients
	case r.QuantityChange <= 0:
		return service.ErrDuplicateMenuIngredients
	case r.Reason == "":
		return service.ErrDuplicateMenuIngredients
	case r.CreatedAt == "":
		return service.ErrDuplicateMenuIngredients
	default:
		return nil
	}
}
