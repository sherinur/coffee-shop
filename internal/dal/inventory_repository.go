package dal

import (
	"encoding/json"
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
	// TODO: open and read file.  return error if file is not found
	// TODO: decode  json and return slice of items

	var inventoryItems []models.InventoryItem

	exists, err := utils.FileExists(r.filePath)
	if err != nil {
		return inventoryItems, err
	}
	if !exists {
		return inventoryItems, nil // Return empty slice if file doesn't exist
	}

	file, err := os.Open(r.filePath)
	if err != nil {
		return inventoryItems, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)

	err = decoder.Decode(&inventoryItems)
	if err != nil {
		return inventoryItems, err
	}

	return inventoryItems, nil
}

func (r *inventoryRepository) GetItemById(id string) (models.InventoryItem, error) {
	var inventoryItem models.InventoryItem
	// TODO: Implement GetItemById() logic
	items, err := r.GetAllItems()
	if err != nil {
		return models.InventoryItem{}, err
	}

	for _, item := range items {
		if item.IngredientID == id {
			return item, nil
		}
	}
	return inventoryItem, nil
}

func (r *inventoryRepository) SaveItems(inventoryItems []models.InventoryItem) error {
	// TODO: Marshal items to JSON and write new file(ioutil.WriteFile, 0644)

	dir := filepath.Dir(r.filePath)
	err := utils.CreateDir(dir)
	if err != nil {
		return err
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
