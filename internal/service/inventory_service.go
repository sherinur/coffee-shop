package service

import (
	"hot-coffee/internal/dal"
	"hot-coffee/models"
)

type InventoryService interface {
	AddInventoryItem(i models.InventoryItem) error
}

type inventoryService struct {
	InventoryRepository dal.InventoryRepository
}

func NewInventoryService(r dal.InventoryRepository) InventoryService {
	return &inventoryService{InventoryRepository: r}
}

func (s *inventoryService) AddInventoryItem(i models.InventoryItem) error {
	// TODO: implement logic to Add a new inventory item.
	return nil
}
