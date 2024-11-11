package service

import (
	"encoding/json"
	"errors"
	"hot-coffee/internal/dal"
	"hot-coffee/models"
)

type InventoryService interface {
	AddInventoryItem(i models.InventoryItem) error
	RetrieveInventoryItems() ([]byte, error)
	RetrieveInventoryItem(id string) ([]byte, error)
	DeleteInventoryItem(id string) error
}

type inventoryService struct {
	InventoryRepository dal.InventoryRepository
}

func NewInventoryService(repo dal.InventoryRepository) *inventoryService {
	if repo == nil {
		return nil
	}
	return &inventoryService{InventoryRepository: repo}
}

var (
	ErrNotUniqueID error = errors.New("ingrediend ID must be unique")
	ErrNoItem      error = errors.New("item not found")
)

func (s *inventoryService) AddInventoryItem(i models.InventoryItem) error {
	inventoryItems, err := s.InventoryRepository.GetAllItems()
	if err != nil {
		return err
	}

	// Uniqueness test (id)
	for _, item := range inventoryItems {
		if item.IngredientID == i.IngredientID {
			return ErrNotUniqueID
		}
	}

	// Appending item to the slice
	inventoryItems = append(inventoryItems, i)

	err = s.InventoryRepository.SaveItems(inventoryItems)
	if err != nil {
		return err
	}

	return nil
}

func (s *inventoryService) RetrieveInventoryItems() ([]byte, error) {
	inventoryItems, err := s.InventoryRepository.GetAllItems()
	if err != nil {
		return nil, err
	}

	data, err := json.MarshalIndent(inventoryItems, "", " ")
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s *inventoryService) RetrieveInventoryItem(id string) ([]byte, error) {
	inventoryItems, err := s.InventoryRepository.GetAllItems()
	if err != nil {
		if err.Error() == "EOF" {
			return nil, ErrNoItem
		}
		return nil, err
	}

	var inventoryItem models.InventoryItem

	isFound := false
	for _, item := range inventoryItems {
		if item.IngredientID == id {
			inventoryItem = item
			isFound = true
			break
		}
	}

	if !isFound {
		return nil, ErrNoItem
	}

	data, err := json.MarshalIndent(inventoryItem, "", " ")
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s *inventoryService) DeleteInventoryItem(id string) error {
	inventoryItems, err := s.InventoryRepository.GetAllItems()
	if err != nil {
		if err.Error() == "EOF" {
			return ErrNoItem
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
		return ErrNoItem
	}

	s.InventoryRepository.SaveItems(inventoryItems)

	return nil
}
