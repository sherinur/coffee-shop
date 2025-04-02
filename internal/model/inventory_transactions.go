package model

import "time"

type InventoryTransactions struct {
	TransactionID  int
	IngredientId   int
	QuantityChange int
	Reason         string
	CreatedAt      time.Time
}

func (r *InventoryTransactions) Validate() error {
	switch {
	case r.TransactionID <= 0:
		return ErrInventoryItemNotFound
	case r.IngredientId <= 0:
		return ErrDuplicateMenuIngredients
	case r.QuantityChange <= 0:
		return ErrDuplicateMenuIngredients
	case r.Reason == "":
		return ErrDuplicateMenuIngredients
	default:
		return nil
	}
}
