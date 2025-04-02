package service

import (
	"strings"

	"coffee-shop/internal/repository/postgres"
	"coffee-shop/models"
)

type MenuService interface {
	AddMenuItem(i models.MenuItem) error
	RetrieveMenuItems() ([]models.MenuItem, error)
	RetrieveMenuItem(id string) (*models.MenuItem, error)
	UpdateMenuItem(id string, item models.MenuItem) error
	DeleteMenuItem(id string) error
}

type menuService struct {
	MenuRepository postgres.Menu
}

func NewMenuService(repo postgres.Menu) *menuService {
	return &menuService{MenuRepository: repo}
}

// TODO: Добавить правило чтобы не повторялись ингредиенты в массиве (один ингредиент и количество сразу пишутся)
// ValidateMenuItem validates the fields of a MenuItem.
// Returns nil if the item is valid.
// The following errors may be returned:
// - ErrNotValidID if the ID is empty.
// - ErrIDContainsSpace if the ID contains spaces.
// - ErrNotValidName if the Name is empty.
// - ErrNotValidDescription if the Description is empty.
// - ErrNotValidPrice if the Price is zero or negative.
// - ErrNotValidIngredients if the Ingredients list is nil or empty.
// - ErrInvalidIngredientID if any ingredient has an invalid ID (empty or contains spaces).
// - ErrInvalidIngredientQty if any ingredient has a quantity less than 1.
func ValidateMenuItem(i models.MenuItem) error {
	if i.ID == "" || strings.Contains(i.ID, " ") {
		return ErrNotValidMenuID
	}

	if i.Name == "" {
		return ErrNotValidMenuName
	}

	if i.Description == "" {
		return ErrNotValidMenuDescription
	}

	if i.Price <= 0 {
		return ErrNotValidPrice
	}

	err := ValidateMenuIngredient(i.Ingredients)
	if err != nil {
		return err
	}

	return nil
}

func ValidateMenuIngredient(i []models.MenuItemIngredient) error {
	if len(i) < 1 {
		return ErrNotEnoughIngredients
	}
	for k, ingredient := range i {
		for l, ingredient2 := range i {
			if ingredient.IngredientID == ingredient2.IngredientID && k != l {
				return ErrDuplicateMenuIngredients
			}
		}
		if ingredient.IngredientID == "" || strings.Contains(ingredient.IngredientID, " ") {
			return ErrNotValidIngredientID
		}

		if ingredient.Quantity < 1 {
			return ErrNotValidQuantity
		}
	}
	return nil
}

// AddMenuItem adds a new menu item to the repository.
// Returns nil if the addition is successful.
// The following errors may be returned:
// - ErrNotUniqueID if the item with the same ID already exists.
// - An error if there is a validation issue or a failure when adding the item to the repository.
func (s *menuService) AddMenuItem(i models.MenuItem) error {
	if exists, err := s.MenuRepository.MenuItemExists(i); err != nil {
		return err
	} else if exists {
		return ErrNotUniqueMenuID
	}

	// Item validation
	if err := ValidateMenuItem(i); err != nil {
		return err
	}

	if _, err := s.MenuRepository.AddMenuItem(i); err != nil {
		return err
	}

	return nil
}

func (s *menuService) RetrieveMenuItems() ([]models.MenuItem, error) {
	return s.MenuRepository.GetAllMenuItems()
}

func (s *menuService) RetrieveMenuItem(id string) (*models.MenuItem, error) {
	if len(id) == 0 {
		return nil, ErrNotValidMenuID
	}

	menuItems, err := s.MenuRepository.GetAllMenuItems()
	if err != nil {
		if err.Error() == "EOF" {
			return nil, ErrNoItem
		}
		return nil, err
	}

	var menuItem models.MenuItem

	isFound := false
	for _, item := range menuItems {
		if item.ID == id {
			menuItem = item
			isFound = true
			break
		}
	}

	if !isFound {
		return nil, ErrNoItem
	}

	return &menuItem, nil
}

func (s *menuService) UpdateMenuItem(id string, i models.MenuItem) error {
	if len(id) == 0 {
		return ErrNotValidMenuID
	}

	// Existence test of old item
	if exists, err := s.MenuRepository.MenuItemExists(models.MenuItem{ID: id}); err != nil {
		return err
	} else if !exists {
		return ErrNoItem
	}

	// Uniqueness test of new item
	if i.ID != id {
		if exists, err := s.MenuRepository.MenuItemExists(i); err != nil {
			return err
		} else if exists {
			return ErrNotUniqueMenuID
		}
	}

	// New item validation
	if err := ValidateMenuItem(i); err != nil {
		return err
	}

	// Rewriting old item in repo
	err := s.MenuRepository.RewriteMenuItem(id, i)
	if err != nil {
		return nil
	}

	return nil
}

func (s *menuService) DeleteMenuItem(id string) error {
	if len(id) == 0 {
		return ErrNotValidMenuID
	}

	menuItems, err := s.MenuRepository.GetAllMenuItems()
	if err != nil {
		if err.Error() == "EOF" {
			return ErrNoItem
		}
		return err
	}

	isFound := false
	// deleting from the slice
	for i, item := range menuItems {
		if item.ID == id {
			menuItems = append(menuItems[:i], menuItems[i+1:]...)
			isFound = true
			break
		}
	}

	if !isFound {
		return ErrNoItem
	}

	err = s.MenuRepository.SaveMenuItems(menuItems)
	if err != nil {
		return err
	}

	return nil
}
