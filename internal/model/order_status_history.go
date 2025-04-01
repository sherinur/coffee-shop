package model

type OrderStatusHistory struct {
	ID       int
	OrderID  int
	OpenedAt string
	ClosedAt string
}
