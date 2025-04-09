package service

import (
	"coffee-shop/internal/model"
	"coffee-shop/internal/repository/postgres"
	"context"
)

type menuService struct {
	MenuRepo postgres.Menu
}

func NewMenuService(repo postgres.Menu) *menuService {
	return &menuService{MenuRepo: repo}
}

// AddMenuItem adds a new menu item to the repository.
// Returns nil if the addition is successful.
// The following errors may be returned:
// - ErrNotUniqueID if the item with the same ID already exists.
// - An error if there is a validation issue or a failure when adding the item to the repository.
func (s *menuService) AddMenuItem(ctx context.Context, item model.MenuItem) error {
	// Item validation
	if err := item.Validate(); err != nil {
		return err
	}

	err := s.MenuRepo.Create(ctx, item)
	if err != nil {
		return err
	}

	return nil
}

func (s *menuService) RetrieveMenuItems(ctx context.Context) ([]model.MenuItem, error) {
	return s.MenuRepo.GetAll(ctx)
}

func (s *menuService) RetrieveMenuItem(ctx context.Context, id int) (*model.MenuItem, error) {
	menuItem, err := s.MenuRepo.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return &menuItem, nil
}

func (s *menuService) UpdateMenuItem(ctx context.Context, id int, item model.MenuItem) error {
	// New item validation
	err := item.Validate()
	if err != nil {
		return err
	}

	// Rewriting old item in repo
	err = s.MenuRepo.Update(ctx, id, item)
	if err != nil {
		return nil
	}

	return nil
}

func (s *menuService) DeleteMenuItem(ctx context.Context, id int) error {
	return s.MenuRepo.Delete(ctx, id)
}
