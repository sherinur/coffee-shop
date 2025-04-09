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
}

type OrderService interface {
}
