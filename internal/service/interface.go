package service

import (
	"context"

	"coffee-shop/internal/model"
)

type InventoryRepo interface {
	Create(ctx context.Context, item model.Inventory) error
	Get(ctx context.Context, id int) (model.Inventory, error)
	GetAll(ctx context.Context) ([]model.Inventory, error)
	Update(ctx context.Context, id int, item model.Inventory) error
	Delete(ctx context.Context, id int) error
}
