package dao

import "coffee-shop/internal/model"

type Order struct {
	OrderID      int    `json:"order_id" db:"order_id"`
	CustomerName string `json:"customer_name" db:"customer_name"`
	Status       string `json:"status" db:"status"`
	Notes        string `json:"notes" db:"notes"`
	CreatedAt    string `json:"created_at" db:"created_at"`
}

func FromOrder(o model.Order) Order {
	return Order{
		OrderID:      o.ID,
		CustomerName: o.CustomerName,
		Status:       o.Status,
		Notes:        o.Notes,
	}
}

func ToOrder(o Order) model.Order {
	return model.Order{
		ID:           o.OrderID,
		CustomerName: o.CustomerName,
		Status:       o.Status,
		Notes:        o.Notes,
	}
}

type OrderItems struct {
	OrderID   int `json:"order_id" db:"order_id"`
	ProductID int `json:"product_id" db:"product_id"`
	Quantity  int `json:"quantity" db:"quantity"`
}

func FromOrderItems(o model.OrderItems) OrderItems {
	return OrderItems{
		OrderID:   o.OrderID,
		ProductID: o.ProductID,
		Quantity:  o.Quantity,
	}
}

func ToOrderItems(o OrderItems) model.OrderItems {
	return model.OrderItems{
		OrderID:   o.OrderID,
		ProductID: o.ProductID,
		Quantity:  o.Quantity,
	}
}

type OrderStatusHistory struct {
	ID       int    `json:"id" db:"id"`
	OrderID  int    `json:"order_id" db:"order_id"`
	OpenedAt string `json:"opened_at" db:"opened_at"`
	ClosedAt string `json:"closed_at" db:"closed_at"`
}

func FromOrderStatusHistory(o model.OrderStatusHistory) OrderStatusHistory {
	return OrderStatusHistory{
		ID:      o.ID,
		OrderID: o.OrderID,
	}
}

func ToOrderStatusHistory(o OrderStatusHistory) model.OrderStatusHistory {
	return model.OrderStatusHistory{
		ID:      o.ID,
		OrderID: o.OrderID,
	}
}
