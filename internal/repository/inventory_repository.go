package repository

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"coffee-shop/internal/utils"
	"coffee-shop/models"
)

type InventoryRepository interface {
	AddItem(i models.InventoryItem) (models.InventoryItem, error)
	GetAllItems() ([]models.InventoryItem, error)
	GetItemById(id string) (models.InventoryItem, error)
	SaveItems(inventoryItems []models.InventoryItem) error
	ItemExists(i models.InventoryItem) (bool, error)
	// ItemExistsById(id string) (bool, error)
	RewriteItem(id string, newItem models.InventoryItem) error
	DeleteItemByID(id string) error
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
	items, err := r.GetAllItems()
	if err != nil {
		return models.InventoryItem{}, err
	}

	items = append(items, i)

	err = r.SaveItems(items)
	if err != nil {
		return models.InventoryItem{}, err
	}

	return i, nil
}

// GetAllItems retrieves all inventory items from the repository.
// Returns an empty slice if the file is empty or does not exist.
// The following errors may be returned:
// - An error if there is a failure in checking file existence or reading the file.
func (r *inventoryRepository) GetAllItems() ([]models.InventoryItem, error) {
	inventoryItems := []models.InventoryItem{}

	exists, err := utils.FileExists(r.filePath)
	if err != nil {
		return []models.InventoryItem{}, err
	}
	if !exists {
		return []models.InventoryItem{}, nil
	}

	file, err := os.Open(r.filePath)
	if err != nil {
		return []models.InventoryItem{}, err
	}
	defer file.Close()

	if utils.FileEmpty(file) {
		return []models.InventoryItem{}, nil
	}

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
	items, err := r.GetAllItems()
	if err != nil {
		return models.InventoryItem{}, err
	}

	for _, item := range items {
		if item.IngredientID == id {
			return item, nil
		}
	}

	return models.InventoryItem{}, errors.New("item not found")
}

// RewriteItem updates an existing inventory item identified by its ID.
// Returns an error if updating the repository fails.
func (r *inventoryRepository) RewriteItem(id string, newItem models.InventoryItem) error {
	items, err := r.GetAllItems()
	if err != nil {
		return err
	}

	for i, item := range items {
		if item.IngredientID == id {
			items[i] = newItem
			break
		}
	}

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

	jsonData, err := json.MarshalIndent(inventoryItems, "", " ")
	if err != nil {
		return err
	}

	err = os.WriteFile(r.filePath, jsonData, 0o644)
	if err != nil {
		return err
	}

	return nil
}

// ItemExists checks if an inventory item with the same ID already exists in the repository.
// Returns true if the item exists, false otherwise.
func (r *inventoryRepository) ItemExists(i models.InventoryItem) (bool, error) {
	inventoryItems, err := r.GetAllItems()
	if err != nil {
		return false, err
	}

	for _, item := range inventoryItems {
		if item.IngredientID == i.IngredientID {
			return true, nil
		}
	}

	return false, nil
}

// func (r *inventoryRepository) ItemExistsById(id string) (bool, error) {
// 	inventoryItems, err := r.GetAllItems()
// 	if err != nil {
// 		return false, err
// 	}

// 	for _, item := range inventoryItems {
// 		if item.IngredientID == id {
// 			return true, nil
// 		}
// 	}

// 	return false, nil
// }

// DeleteInventoryItem deletes an inventory item by its ID.
// Returns nil if the deletion is successful.
// The following errors may be returned:
// - ErrNoItem if the item with the specified ID is not found.
// - An error if there is a failure when retrieving or saving items in the repository.
func (r *inventoryRepository) DeleteItemByID(id string) error {
	inventoryItems, err := r.GetAllItems()
	if err != nil {
		if err.Error() == "EOF" {
			return errors.New("item not found")
		}
		return err
	}

	isFound := false
	// deleting from the slice
	for i, item := range inventoryItems {
		if item.IngredientID == id {
			inventoryItems = append(inventoryItems[:i], inventoryItems[i+1:]...)
			isFound = true
			break
		}
	}

	if !isFound {
		return errors.New("item not found")
	}

	err = r.SaveItems(inventoryItems)
	if err != nil {
		return err
	}

	return nil
}
