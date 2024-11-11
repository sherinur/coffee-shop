package dal

import (
	"encoding/json"
	"errors"
	"fmt"
	"hot-coffee/internal/utils"
	"hot-coffee/models"
	"os"
	"path/filepath"
)

type InventoryRepository interface {
	AddItem(i models.InventoryItem) (models.InventoryItem, error)
	GetAllItems() ([]models.InventoryItem, error)
	GetItemById(id string) (models.InventoryItem, error)
	SaveItems(inventoryItems []models.InventoryItem) error
}

type inventoryRepository struct {
	filePath string
}

func NewInventoryRepository(filePath string) *inventoryRepository {
	return &inventoryRepository{filePath: filePath}
}

func (r *inventoryRepository) AddItem(i models.InventoryItem) (models.InventoryItem, error) {
	// TODO: Marshal item to JSON and append to the file inventory.json (ioutil.WriteFile, 0644)
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
	var inventoryItems []models.InventoryItem

	exists, err := utils.FileExists(r.filePath)
	if err != nil {
		return inventoryItems, err
	}
	if !exists {
		return inventoryItems, nil
	}

	file, err := os.Open(r.filePath)
	if err != nil {
		return inventoryItems, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)

	if stat, _ := file.Stat(); stat.Size() == 0 {
		return inventoryItems, nil
	}

	err = decoder.Decode(&inventoryItems)
	if err != nil {
		return inventoryItems, err
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

	// Opening a file for recording
	file, err := os.OpenFile(r.filePath, os.O_RDWR|os.O_CREATE, 0o644)
	if err != nil {
		return fmt.Errorf("failed to open file %s for writing: %w", r.filePath, err)
	}
	defer file.Close()

	// Checking for file emptiness before writing
	stat, err := file.Stat()
	if err != nil {
		return fmt.Errorf("failed to get file stat: %w", err)
	}
	if stat.Size() == 0 {
		inventoryItems = []models.InventoryItem{}
	}

	// // Converting data to JSON format
	jsonData, err := json.MarshalIndent(inventoryItems, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshalling inventory items: %w", err)
	}

	// Writing data to a file
	_, err = file.Write(jsonData)
	if err != nil {
		return fmt.Errorf("failed to write data to file %s: %w", r.filePath, err)
	}

	return nil
}
