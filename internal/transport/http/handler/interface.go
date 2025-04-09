package handler

import (
	"coffee-shop/internal/model"
)

type InventoryService interface {
	AddInventoryItem(i model.Inventory) error
	RetrieveInventoryItems() ([]model.Inventory, error)
	RetrieveInventoryItem(id string) (*model.Inventory, error)
	UpdateInventoryItem(id string, item model.Inventory) error
	DeleteInventoryItem(id string) error
}

type MenuService interface {
	AddMenuItem(i model.MenuItem) error
	RetrieveMenuItems() ([]model.MenuItem, error)
	RetrieveMenuItem(id string) (*model.MenuItem, error)
	UpdateMenuItem(id string, item model.MenuItem) error
	DeleteMenuItem(id string) error
}

type OrderService interface {
	AddOrder(o model.Order) error
	RetrieveOrders() ([]model.Order, error)
	RetrieveOrder(id string) (model.Order, error)
	UpdateOrder(id string, item model.Order) error
	DeleteOrder(id string) error
	CloseOrder(id string) error
	// IsInventorySufficient(orderItems []models.OrderItem) (bool, error)
	// ReduceIngredients(orderItems []models.OrderItem) error
	// CalculateTotalSales() (float64, error)
}
