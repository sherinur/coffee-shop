package service

import (
	"hot-coffee/internal/dal"
)

type InventoryService interface {
	AddInventoryItem()
}

type inventoryService struct {
	InventoryRepository dal.InventoryRepository
}

func NewInventoryService(r dal.InventoryRepository) InventoryService {
	return &inventoryService{InventoryRepository: r}
}

func (s *inventoryService) AddInventoryItem() {
	// TODO: implement logic to Add a new inventory item.
}
