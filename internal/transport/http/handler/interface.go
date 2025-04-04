package handler

import "coffee-shop/internal/model"

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
