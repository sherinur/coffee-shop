package model

import "coffee-shop/internal/service"

type MenuItems struct {
	ID          int
	Name        string
	Description string
	Price       float64
}

func (r *MenuItems) Validate() error {
	switch {
	case r.ID <= 0:
		return service.ErrNotValidMenuID
	case r.Name == "":
		return service.ErrNotValidMenuName
	case r.Description == "":
		return service.ErrNotValidMenuDescription
	case r.Price <= 0:
		return service.ErrNotValidPrice
	default:
		return nil
	}
}
