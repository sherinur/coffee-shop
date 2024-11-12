package dal

import (
	"encoding/json"
	"errors"
	"os"

	"hot-coffee/internal/utils"
	"hot-coffee/models"
)

type MenuRepository interface {
	AddMenuItem(i models.MenuItem) (models.MenuItem, error)
	GetAllMenuItems() ([]models.MenuItem, error)
	GetMenuItemById(id string) (models.MenuItem, error)
	SaveMenuItems(menuItems []models.MenuItem) error
	MenuItemExists(i models.MenuItem) (bool, error)
	RewriteMenuItem(id string, newItem models.MenuItem) error
}

type menuRepository struct {
	filePath string
}

func NewMenuRepository(filePath string) *menuRepository {
	return &menuRepository{filePath: filePath}
}

func (r *menuRepository) AddMenuItem(i models.MenuItem) (models.MenuItem, error) {
	// TODO:AddMenuItem logic
	return i, nil
}

func (r *menuRepository) GetAllMenuItems() ([]models.MenuItem, error) {
	menuItems := []models.MenuItem{}

	exists, err := utils.FileExists(r.filePath)
	if err != nil {
		return []models.MenuItem{}, err
	}
	if !exists {
		return []models.MenuItem{}, nil
	}

	file, err := os.Open(r.filePath)
	if err != nil {
		return []models.MenuItem{}, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)

	if utils.FileEmpty(file) {
		return []models.MenuItem{}, nil
	}

	err = decoder.Decode(&menuItems)
	if err != nil {
		return []models.MenuItem{}, err
	}

	return menuItems, nil
}

func (r *menuRepository) GetMenuItemById(id string) (models.MenuItem, error) {
	items, err := r.GetAllMenuItems()
	if err != nil {
		return models.MenuItem{}, err
	}

	for _, item := range items {
		if item.ID == id {
			return item, nil
		}
	}
	return models.MenuItem{}, errors.New("item not found")
}

func (r *menuRepository) MenuItemExists(i models.MenuItem) (bool, error) {
	// TODO:MenuItemExists logic
	return false, nil
}

func (r *menuRepository) RewriteMenuItem(id string, newItem models.MenuItem) error {
	// TODO:RewriteMenuItem logic
	return nil
}

func (r *menuRepository) SaveMenuItems(menuItems []models.MenuItem) error {
	// TODO:RewriteMenuItem logic
	return nil
}
