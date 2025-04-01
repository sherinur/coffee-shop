package model

import "coffee-shop/internal/service"

type Inventory struct {
	IngredientID int
	Name         string
	Quantity     int
	Unit         string
}

func (r *Inventory) Validate() error {
	switch {
	case r.IngredientID <= 0:
		return service.ErrNotValidIngredientID
	case r.Name == "":
		return service.ErrNotValidIngredientName
	case r.Quantity <= 0:
		return service.ErrNotValidQuantity
	case r.Unit == "":
		return service.ErrNotValidUnit
	default:
		return nil
	}
}
