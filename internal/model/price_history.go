package model

type PriceHistory struct {
	HistoryID  int
	MenuItemID int
	Price      float64
	ChangedAt  string
}
