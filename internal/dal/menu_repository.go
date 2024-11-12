package dal

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"

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
	items, err := r.GetAllMenuItems()
	if err != nil {
		return models.MenuItem{}, err
	}

	items = append(items, i)

	err = r.SaveMenuItems(items)
	if err != nil {
		return models.MenuItem{}, err
	}

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

func (r *menuRepository) SaveMenuItems(menuItems []models.MenuItem) error {
	// Checking the existence of a directory for a file
	dir := filepath.Dir(r.filePath)
	err := utils.CreateDir(dir)
	if err != nil {
		return fmt.Errorf("failed to create directory for file %s: %w", dir, err)
	}

	// Checking file write permissions
	exists, err := utils.FileExists(r.filePath)
	if err != nil {
		return fmt.Errorf("error checking if file exists: %w", err)
	}

	if !exists {
		// If the file does not exist, create it
		err := utils.CreateFile(r.filePath)
		if err != nil {
			return fmt.Errorf("error creating file %s: %w", r.filePath, err)
		}
	}

	jsonData, err := json.MarshalIndent(menuItems, "", " ")
	if err != nil {
		return err
	}

	err = os.WriteFile(r.filePath, jsonData, 0o644)
	if err != nil {
		return err
	}

	return nil
}

func (r *menuRepository) MenuItemExists(i models.MenuItem) (bool, error) {
	menuItems, err := r.GetAllMenuItems()
	if err != nil {
		return false, err
	}

	for _, item := range menuItems {
		if item.ID == i.ID {
			return true, nil
		}
	}

	return false, nil
}

func (r *menuRepository) RewriteMenuItem(id string, newItem models.MenuItem) error {
	items, err := r.GetAllMenuItems()
	if err != nil {
		return err
	}

	for i, item := range items {
		if item.ID == id {
			items[i] = newItem
			break
		}
	}

	err = r.SaveMenuItems(items)
	if err != nil {
		return err
	}

	return nil
}
