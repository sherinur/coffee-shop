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

type MenuRepo interface {
	Create(ctx context.Context, menu model.MenuItem) error
	Get(ctx context.Context, id int) (model.MenuItem, error)
	GetAll(ctx context.Context) ([]model.MenuItem, error)
	Update(ctx context.Context, id int, menu model.MenuItem) error
	Delete(ctx context.Context, id int) error
}

type OrderRepo interface {
	Create(ctx context.Context, order model.Order) error
	Get(ctx context.Context, id int) (model.Order, error)
	GetAll(ctx context.Context) ([]model.Order, error)
	Update(ctx context.Context, id int, order model.Order) error
	Delete(ctx context.Context, id int) error
}
