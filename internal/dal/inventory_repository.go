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

type InventoryRepository interface {
	AddItem(i models.InventoryItem) (models.InventoryItem, error)
	GetAllItems() ([]models.InventoryItem, error)
	GetItemById(id string) (models.InventoryItem, error)
	SaveItems(inventoryItems []models.InventoryItem) error
	ItemExists(i models.InventoryItem) (bool, error)
	RewriteItem(id string, newItem models.InventoryItem) error
}

type inventoryRepository struct {
	filePath string
}

func NewInventoryRepository(filePath string) *inventoryRepository {
	return &inventoryRepository{filePath: filePath}
}

// AddItem adds a new inventory item to the repository.
// Returns the added item if successful.
// The following errors may be returned:
// - An error if there is a failure in retrieving or saving the items.
func (r *inventoryRepository) AddItem(i models.InventoryItem) (models.InventoryItem, error) {
	// Retrieve all current inventory items from the repository.
	items, err := r.GetAllItems()
	if err != nil {
		// Return an error if there is a failure in retrieving items.
		return models.InventoryItem{}, err
	}

	// Append the new inventory item to the list.
	items = append(items, i)

	// Save the updated list of items back to the repository.
	err = r.SaveItems(items)
	if err != nil {
		// Return an error if there is a failure in saving items.
		return models.InventoryItem{}, err
	}

	// Return the added inventory item if successful.
	return i, nil
}

// GetAllItems retrieves all inventory items from the repository.
// Returns an empty slice if the file is empty or does not exist.
// The following errors may be returned:
// - An error if there is a failure in checking file existence or reading the file.
func (r *inventoryRepository) GetAllItems() ([]models.InventoryItem, error) {
	inventoryItems := []models.InventoryItem{}

	// Check if the file exists.
	exists, err := utils.FileExists(r.filePath)
	if err != nil {
		return []models.InventoryItem{}, err
	}
	if !exists {
		// Return an empty slice if the file does not exist.
		return []models.InventoryItem{}, nil
	}

	// Open the file for reading.
	file, err := os.Open(r.filePath)
	if err != nil {
		return []models.InventoryItem{}, err
	}
	defer file.Close()

	// Check if the file is empty.
	if utils.FileEmpty(file) {
		return []models.InventoryItem{}, nil
	}

	// Decode JSON data into the inventory items slice.
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&inventoryItems)
	if err != nil {
		return []models.InventoryItem{}, err
	}

	return inventoryItems, nil
}

// GetItemById retrieves a specific inventory item by its ID.
// Returns an error if the item is not found.
func (r *inventoryRepository) GetItemById(id string) (models.InventoryItem, error) {
	// Retrieve all items from the repository.
	items, err := r.GetAllItems()
	if err != nil {
		return models.InventoryItem{}, err
	}

	// Search for the item with the specified ID.
	for _, item := range items {
		if item.IngredientID == id {
			return item, nil
		}
	}

	// Return an error if the item is not found.
	return models.InventoryItem{}, errors.New("item not found")
}

// RewriteItem updates an existing inventory item identified by its ID.
// Returns an error if updating the repository fails.
func (r *inventoryRepository) RewriteItem(id string, newItem models.InventoryItem) error {
	// Retrieve all items from the repository.
	items, err := r.GetAllItems()
	if err != nil {
		return err
	}

	// Search for the item and update it.
	for i, item := range items {
		if item.IngredientID == id {
			items[i] = newItem
			break
		}
	}

	// Save the updated list of items.
	err = r.SaveItems(items)
	if err != nil {
		return err
	}

	return nil
}

// SaveItems writes the provided inventory items to the repository file.
// Creates the directory and file if they do not exist.
// The following errors may be returned:
// - An error if creating the directory or file fails.
// - An error if writing to the file fails.
func (r *inventoryRepository) SaveItems(inventoryItems []models.InventoryItem) error {
	// Ensure the directory for the file exists.
	dir := filepath.Dir(r.filePath)
	err := utils.CreateDir(dir)
	if err != nil {
		return fmt.Errorf("failed to create directory for file %s: %w", dir, err)
	}

	// Check if the file exists, and create it if necessary.
	exists, err := utils.FileExists(r.filePath)
	if err != nil {
		return fmt.Errorf("error checking if file exists: %w", err)
	}

	if !exists {
		// Create the file if it does not exist.
		err := utils.CreateFile(r.filePath)
		if err != nil {
			return fmt.Errorf("error creating file %s: %w", r.filePath, err)
		}
	}

	// Marshal the inventory items into JSON format.
	jsonData, err := json.MarshalIndent(inventoryItems, "", " ")
	if err != nil {
		return err
	}

	// Write the JSON data to the file.
	err = os.WriteFile(r.filePath, jsonData, 0o644)
	if err != nil {
		return err
	}

	return nil
}

// ItemExists checks if an inventory item with the same ID already exists in the repository.
// Returns true if the item exists, false otherwise.
func (r *inventoryRepository) ItemExists(i models.InventoryItem) (bool, error) {
	// Retrieve all items from the repository.
	inventoryItems, err := r.GetAllItems()
	if err != nil {
		return false, err
	}

	// Check if an item with the same ID exists.
	for _, item := range inventoryItems {
		if item.IngredientID == i.IngredientID {
			return true, nil
		}
	}

	return false, nil
}
