package model

import "coffee-shop/internal/service"

type MenuItemIngredients struct {
	MenuID       int
	IngredientID int
	Quantity     int
}

func (r *MenuItemIngredients) Validate() error {
	switch {
	case r.MenuID <= 0:
		return service.ErrNotValidMenuID
	case r.IngredientID <= 0:
		return service.ErrNotValidIngredientID
	case r.Quantity <= 0:
		return service.ErrNotValidQuantity
	default:
		return nil
	}
}
