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

// AddMenuItem adds a new menu item to the repository.
// It first retrieves the current list of all menu items,
// then appends the new item to the list, and saves the updated list back to the repository.
// Returns the added item if successful, or an error if there was a failure during retrieval or saving of the items.
func (r *menuRepository) AddMenuItem(i models.MenuItem) (models.MenuItem, error) {
	// Retrieve all current menu items from the repository.
	items, err := r.GetAllMenuItems()
	if err != nil {
		// Return an error if there is a failure retrieving the list of menu items.
		return models.MenuItem{}, err
	}

	// Append the new menu item to the existing list.
	items = append(items, i)

	// Save the updated list of menu items back to the repository.
	err = r.SaveMenuItems(items)
	if err != nil {
		// Return an error if there is a failure in saving the updated list of items.
		return models.MenuItem{}, err
	}

	// Return the newly added menu item if the operation was successful.
	return i, nil
}

// GetAllMenuItems retrieves all menu items from the repository.
// It checks if the file exists, and if it does, opens it and decodes the list of menu items.
// Returns the list of menu items if successful, or an empty list and an error if there was an issue reading the file.
func (r *menuRepository) GetAllMenuItems() ([]models.MenuItem, error) {
	// Initialize an empty slice to hold the menu items.
	menuItems := []models.MenuItem{}

	// Check if the file containing the menu items exists.
	exists, err := utils.FileExists(r.filePath)
	if err != nil {
		// Return an empty list and error if checking file existence fails.
		return []models.MenuItem{}, err
	}
	if !exists {
		// Return an empty list if the file doesn't exist.
		return []models.MenuItem{}, nil
	}

	// Open the file containing the menu items.
	file, err := os.Open(r.filePath)
	if err != nil {
		// Return an empty list and error if opening the file fails.
		return []models.MenuItem{}, err
	}
	defer file.Close()

	// Create a new JSON decoder to read the file contents.
	decoder := json.NewDecoder(file)

	// Check if the file is empty, returning an empty list if it is.
	if utils.FileEmpty(file) {
		return []models.MenuItem{}, nil
	}

	// Decode the file contents into the menuItems slice.
	err = decoder.Decode(&menuItems)
	if err != nil {
		// Return an empty list and error if decoding the JSON fails.
		return []models.MenuItem{}, err
	}

	// Return the list of menu items.
	return menuItems, nil
}

// GetMenuItemById retrieves a menu item by its ID from the repository.
// It fetches all menu items and searches for the item with the matching ID.
// Returns the item if found, or an error if the item is not found.
func (r *menuRepository) GetMenuItemById(id string) (models.MenuItem, error) {
	// Retrieve all menu items from the repository.
	items, err := r.GetAllMenuItems()
	if err != nil {
		// Return an empty item and the error if retrieving the menu items fails.
		return models.MenuItem{}, err
	}

	// Iterate through the list of items to find the item with the matching ID.
	for _, item := range items {
		if item.ID == id {
			// Return the item if it matches the provided ID.
			return item, nil
		}
	}

	// Return an error if no item with the given ID is found.
	return models.MenuItem{}, errors.New("item not found")
}

// SaveMenuItems saves the provided menu items to a file in JSON format.
// It ensures that the file's directory exists, creates the file if necessary,
// and checks for write permissions before writing the data.
func (r *menuRepository) SaveMenuItems(menuItems []models.MenuItem) error {
	// Checking the existence of the directory for the file path.
	dir := filepath.Dir(r.filePath)
	err := utils.CreateDir(dir)
	if err != nil {
		// Return an error if the directory creation fails.
		return fmt.Errorf("failed to create directory for file %s: %w", dir, err)
	}

	// Checking if the file already exists.
	exists, err := utils.FileExists(r.filePath)
	if err != nil {
		// Return an error if checking file existence fails.
		return fmt.Errorf("error checking if file exists: %w", err)
	}

	// If the file does not exist, create it.
	if !exists {
		err := utils.CreateFile(r.filePath)
		if err != nil {
			// Return an error if creating the file fails.
			return fmt.Errorf("error creating file %s: %w", r.filePath, err)
		}
	}

	// Marshal the menu items into a JSON format with indentation.
	jsonData, err := json.MarshalIndent(menuItems, "", " ")
	if err != nil {
		// Return an error if marshaling the data to JSON fails.
		return err
	}

	// Write the JSON data to the file.
	err = os.WriteFile(r.filePath, jsonData, 0o644)
	if err != nil {
		// Return an error if writing the file fails.
		return err
	}

	// Return nil if everything was successful.
	return nil
}

// MenuItemExists checks whether a menu item with the specified ID already exists in the repository.
// It retrieves all menu items and compares each item's ID with the provided item's ID.
func (r *menuRepository) MenuItemExists(i models.MenuItem) (bool, error) {
	// Retrieve all menu items from the repository.
	menuItems, err := r.GetAllMenuItems()
	if err != nil {
		// Return an error if retrieving the menu items fails.
		return false, err
	}

	// Iterate over all menu items to check if any item's ID matches the provided ID.
	for _, item := range menuItems {
		if item.ID == i.ID {
			// Return true if a match is found, indicating that the menu item already exists.
			return true, nil
		}
	}

	// Return false if no match is found, indicating that the menu item does not exist.
	return false, nil
}

// RewriteMenuItem updates an existing menu item in the repository with a new item.
// It searches for the item by its ID and replaces it with the provided new item.
// If successful, the updated list of menu items is saved back to the repository.
func (r *menuRepository) RewriteMenuItem(id string, newItem models.MenuItem) error {
	// Retrieve all current menu items from the repository.
	items, err := r.GetAllMenuItems()
	if err != nil {
		// Return an error if there is a failure in retrieving the items.
		return err
	}

	// Iterate through the items and find the one with the matching ID.
	for i, item := range items {
		if item.ID == id {
			// Replace the found item with the new item.
			items[i] = newItem
			break
		}
	}

	// Save the updated list of menu items back to the repository.
	err = r.SaveMenuItems(items)
	if err != nil {
		// Return an error if there is a failure in saving the items.
		return err
	}

	// Return nil if the item was successfully rewritten and saved.
	return nil
}
