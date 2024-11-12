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

func (s *menuService) AddMenuItem(i models.MenuItem) error {
	return nil
}

func (s *menuService) DeleteMenuItem(id string) error {
	return nil
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

func (s *menuService) UpdateMenuItem(id string, item models.MenuItem) error {
	panic("unimplemented")
}
