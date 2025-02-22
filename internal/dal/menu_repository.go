package dal

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"coffee-shop/internal/utils"
	"coffee-shop/models"
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

// AddMenuItem adds a new menu item to the repository.
// It first retrieves the current list of all menu items,
// then appends the new item to the list, and saves the updated list back to the repository.
// Returns the added item if successful, or an error if there was a failure during retrieval or saving of the items.
func (r *menuRepository) AddMenuItem(i models.MenuItem) (models.MenuItem, error) {
	items, err := r.GetAllMenuItems()
	if err != nil {
		return models.MenuItem{}, err
	}

	itemsID := []string{}

	for _, item := range items {
		itemsID = append(itemsID, item.ID)
	}

	if i.ID == "" {
		i.ID = utils.GenerateNewID(itemsID, "menu")
	}

	items = append(items, i)

	err = r.SaveMenuItems(items)
	if err != nil {
		return models.MenuItem{}, err
	}

	return i, nil
}

// GetAllMenuItems retrieves all menu items from the repository.
// It checks if the file exists, and if it does, opens it and decodes the list of menu items.
// Returns the list of menu items if successful, or an empty list and an error if there was an issue reading the file.
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

// GetMenuItemById retrieves a menu item by its ID from the repository.
// It fetches all menu items and searches for the item with the matching ID.
// Returns the item if found, or an error if the item is not found.
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

// SaveMenuItems saves the provided menu items to a file in JSON format.
// It ensures that the file's directory exists, creates the file if necessary,
// and checks for write permissions before writing the data.
func (r *menuRepository) SaveMenuItems(menuItems []models.MenuItem) error {
	dir := filepath.Dir(r.filePath)
	err := utils.CreateDir(dir)
	if err != nil {
		return fmt.Errorf("failed to create directory for file %s: %w", dir, err)
	}

	exists, err := utils.FileExists(r.filePath)
	if err != nil {
		return fmt.Errorf("error checking if file exists: %w", err)
	}

	if !exists {
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

// MenuItemExists checks whether a menu item with the specified ID already exists in the repository.
// It retrieves all menu items and compares each item's ID with the provided item's ID.
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

// RewriteMenuItem updates an existing menu item in the repository with a new item.
// It searches for the item by its ID and replaces it with the provided new item.
// If successful, the updated list of menu items is saved back to the repository.
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
