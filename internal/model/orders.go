package model

import "coffee-shop/internal/service"

type Orders struct {
	ID           int
	CustomerName string
	Status       string
	Notes        string
	CreatedAt    string
}

func (r *Orders) Validate() error {
	switch {
	case r.ID <= 0:
		return service.ErrNotValidOrderID
	case r.CustomerName == "":
		return service.ErrNotValidOrderCustomerName
	case r.Status == "":
		return service.ErrNotValidOrderStatus
	case r.CreatedAt == "":
		return service.ErrNotValidCreatedAt
	default:
		return nil
	}
}
