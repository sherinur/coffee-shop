package model

type MenuItems struct {
	ID          int
	Name        string
	Description string
	Price       float64
}

func (r *MenuItems) Validate() error {
	switch {
	case r.ID <= 0:
		return ErrNotValidMenuID
	case r.Name == "":
		return ErrNotValidMenuName
	case r.Description == "":
		return ErrNotValidMenuDescription
	case r.Price <= 0:
		return ErrNotValidPrice
	default:
		return nil
	}
}
