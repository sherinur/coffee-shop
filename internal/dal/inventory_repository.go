package dal

import "hot-coffee/models"

type InventoryRepository interface {
	AddItem(i models.InventoryItem) (models.InventoryItem, error)
	GetAllItems(i models.InventoryItem) ([]models.InventoryItem, error)
}

type inventoryRepository struct {
	filePath string
}

func NewInventoryRepository(filePath string) InventoryRepository {
	return &inventoryRepository{filePath: filePath}
}

func (r *inventoryRepository) AddItem(i models.InventoryItem) (models.InventoryItem, error) {
	// TODO: getAllItems and append new order
	// TODO: Marshal items to JSON and write new file(ioutil.WriteFile, 0644)

	return models.InventoryItem{}, nil
}

func (r *inventoryRepository) GetAllItems(i models.InventoryItem) ([]models.InventoryItem, error) {
	// TODO: open and read file.  return error if file is not found
	// TODO: decode  json and return slice of items

	return []models.InventoryItem{}, nil
}
