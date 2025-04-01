package model

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
		return ErrInventoryItemNotFound
	case r.IngredientId <= 0:
		return ErrDuplicateMenuIngredients
	case r.QuantityChange <= 0:
		return ErrDuplicateMenuIngredients
	case r.Reason == "":
		return ErrDuplicateMenuIngredients
	case r.CreatedAt == "":
		return ErrDuplicateMenuIngredients
	default:
		return nil
	}
}
