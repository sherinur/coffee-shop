package dao

type Order struct {
	OrderID      int    `json:"order_id" db:"order_id"`
	CustomerName string `json:"customer_name" db:"customer_name"`
	Status       string `json:"status" db:"status"`
	Notes        string `json:"notes" db:"notes"`
	CreatedAt    string `json:"created_at" db:"created_at"`
}

type OrderItems struct {
	OrderID   int `json:"order_id" db:"order_id"`
	ProductID int `json:"product_id" db:"product_id"`
	Quantity  int `json:"quantity" db:"quantity"`
}

type OrderStatusHistory struct {
	ID       int    `json:"id" db:"id"`
	OrderID  int    `json:"order_id" db:"order_id"`
	OpenedAt string `json:"opened_at" db:"opened_at"`
	ClosedAt string `json:"closed_at" db:"closed_at"`
}
