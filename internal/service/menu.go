package service

import (
	"coffee-shop/internal/model"
	"coffee-shop/internal/repository/postgres"
	"context"
)

type menuService struct {
	MenuRepo            postgres.Menu
	MenuIngredientsRepo postgres.MenuItemIngredients
}

func NewMenuService(menuRepo postgres.Menu, menuIngRepo postgres.MenuItemIngredients) *menuService {
	return &menuService{
		MenuRepo:            menuRepo,
		MenuIngredientsRepo: menuIngRepo}
}

// AddMenuItem adds a new menu item to the repository.
// Returns nil if the addition is successful.
// The following errors may be returned:
// - ErrNotUniqueID if the item with the same ID already exists.
// - An error if there is a validation issue or a failure when adding the item to the repository.
func (s *menuService) AddMenuItem(ctx context.Context, menu model.MenuItem, ingredients []model.MenuItemIngredients) error {
	// Item validation
	if err := menu.Validate(); err != nil {
		return err
	}

	for _, i := range ingredients {
		if err := i.Validate(); err != nil {
			return err
		}
	}

	err := s.MenuRepo.Create(ctx, menu)
	if err != nil {
		return err
	}

	for _, i := range ingredients {
		err := s.MenuIngredientsRepo.Create(ctx, i)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *menuService) RetrieveMenuItems(ctx context.Context) ([]model.MenuItem, error) {
	return s.MenuRepo.GetAll(ctx)
}

func (s *menuService) RetrieveMenuItemWithId(ctx context.Context, id int) (*model.MenuItem, []model.MenuItemIngredients, error) {
	menuItem, err := s.MenuRepo.Get(ctx, id)
	if err != nil {
		return nil, nil, err
	}

	menuItemIngredients, err := s.MenuIngredientsRepo.GetAllWithID(ctx, id)

	return &menuItem, menuItemIngredients, nil
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
