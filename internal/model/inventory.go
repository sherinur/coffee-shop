package model

type Inventory struct {
	IngredientID int
	Name         string
	Quantity     int
	Unit         string
}

func (r *Inventory) Validate() error {
	switch {
	case r.IngredientID <= 0:
		return ErrNotValidIngredientID
	case r.Name == "":
		return ErrNotValidIngredientName
	case r.Quantity <= 0:
		return ErrNotValidQuantity
	case r.Unit == "":
		return ErrNotValidUnit
	default:
		return nil
	}
}
