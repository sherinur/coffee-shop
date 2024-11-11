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
}

type inventoryRepository struct {
	filePath string
}

func NewInventoryRepository(filePath string) *inventoryRepository {
	return &inventoryRepository{filePath: filePath}
}

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

	decoder := json.NewDecoder(file)

	if utils.FileEmpty(file) {
		return []models.InventoryItem{}, nil
	}

	err = decoder.Decode(&inventoryItems)
	if err != nil {
		return []models.InventoryItem{}, err
	}

	return inventoryItems, nil
}

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

func (r *inventoryRepository) SaveItems(inventoryItems []models.InventoryItem) error {
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

func (r *inventoryRepository) ItemExists(i models.InventoryItem) (bool, error) {
	inventoryItems, err := r.GetAllItems()
	if err != nil {
		return false, err
	}

	// Uniqueness test (id)
	for _, item := range inventoryItems {
		if item.IngredientID == i.IngredientID {
			return true, nil
		}
	}

	return false, nil
}
