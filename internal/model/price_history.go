package model

type PriceHistory struct {
	HistoryID  int
	MenuItemID int
	Price      float64
	ChangedAt  string
}

func (r *PriceHistory) Validate() error {
	switch {
	case r.HistoryID <= 0:
		return ErrNotValidPriceHistoryID
	case r.MenuItemID <= 0:
		return ErrNotValidMenuID
	case r.Price <= 0:
		return ErrNotValidPrice
	case r.ChangedAt == "":
		return ErrNotValidChangedAtTime
	default:
		return nil
	}
}
