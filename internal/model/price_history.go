package model

import "coffee-shop/internal/service"

type PriceHistory struct {
	HistoryID  int
	MenuItemID int
	Price      float64
	ChangedAt  string
}

func (r *PriceHistory) Validate() error {
	switch {
	case r.HistoryID <= 0:
		return service.ErrNotValidPriceHistoryID
	case r.MenuItemID <= 0:
		return service.ErrNotValidMenuID
	case r.Price <= 0:
		return service.ErrNotValidPrice
	case r.ChangedAt == "":
		return service.ErrNotValidChangedAtTime
	default:
		return nil
	}
}
