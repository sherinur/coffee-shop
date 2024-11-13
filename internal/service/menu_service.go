package service

import (
	"encoding/json"
	"strings"

	"hot-coffee/internal/dal"
	"hot-coffee/models"
)

type MenuService interface {
	AddMenuItem(i models.MenuItem) error
	RetrieveMenuItems() ([]byte, error)
	RetrieveMenuItem(id string) ([]byte, error)
	UpdateMenuItem(id string, item models.MenuItem) error
	DeleteMenuItem(id string) error
}

type menuService struct {
	MenuRepository dal.MenuRepository
}

func NewMenuService(repo dal.MenuRepository) *menuService {
	if repo == nil {
		return nil
	}
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

	if i.Ingredients == nil || len(i.Ingredients) < 1 {
		return ErrNotValidIngredientID
	}

	for _, ingredient := range i.Ingredients {
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

func (s *menuService) RetrieveMenuItems() ([]byte, error) {
	menuItems, err := s.MenuRepository.GetAllMenuItems()
	if err != nil {
		return nil, err
	}

	data, err := json.MarshalIndent(menuItems, "", " ")
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s *menuService) RetrieveMenuItem(id string) ([]byte, error) {
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

	data, err := json.MarshalIndent(menuItem, "", " ")
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s *menuService) UpdateMenuItem(id string, i models.MenuItem) error {
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
