package handler

import (
	"coffee-shop/internal/model"
	"context"
)

type InventoryService interface {
	AddInventoryItem(ctx context.Context, item model.Inventory) error
	RetrieveInventoryItems(ctx context.Context) ([]model.Inventory, error)
	RetrieveInventoryItem(ctx context.Context, id int) (*model.Inventory, error)
	UpdateInventoryItem(ctx context.Context, id int, item model.Inventory) error
	DeleteInventoryItem(ctx context.Context, id int) error
}

type MenuService interface {
	AddMenuItem(ctx context.Context, menu model.MenuItem, ingredients []model.MenuItemIngredients) error
	RetrieveMenuItems(ctx context.Context) ([]model.MenuItem, error)
	RetrieveMenuItemWithId(ctx context.Context, id int) (*model.MenuItem, []model.MenuItemIngredients, error)
	UpdateMenuItem(ctx context.Context, id int, item model.MenuItem, ingredients []model.MenuItemIngredients) error
	DeleteMenuItem(ctx context.Context, id int) error
}

type OrderService interface {
	AddOrder(ctx context.Context, order model.Order) error
	RetrieveOrders(ctx context.Context) ([]model.Order, error)
	RetrieveOrder(ctx context.Context, id int) (*model.Order, error)
	UpdateOrder(ctx context.Context, id int, order model.Order) error
	DeleteOrder(ctx context.Context, id int) error
}
