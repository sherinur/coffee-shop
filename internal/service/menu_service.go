package service

import (
	"encoding/json"

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

func ValidateMenuItem(i models.MenuItem) error {
	if i.ID == "" {
		return ErrNotValidID
	}

	if i.Name == "" {
		return ErrNotValidName
	}

	if i.Description == "" {
		return ErrNotValidDescription
	}

	if i.Price <= 0 {
		return ErrNotValidPrice
	}

	if i.Ingredients == nil {
		return ErrNotValidIngredients
	}

	return nil
}

func (s *menuService) AddMenuItem(i models.MenuItem) error {
	if exists, err := s.MenuRepository.MenuItemExists(i); err != nil {
		return err
	} else if exists {
		return ErrNotUniqueID
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
			return ErrNotUniqueID
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
