package model

type MenuItemIngredients struct {
	MenuID       int
	IngredientID int
	Quantity     int
}

func (r *MenuItemIngredients) Validate() error {
	switch {
	case r.MenuID <= 0:
		return ErrNotValidMenuID
	case r.IngredientID <= 0:
		return ErrNotValidIngredientID
	case r.Quantity <= 0:
		return ErrNotValidQuantity
	default:
		return nil
	}
}
