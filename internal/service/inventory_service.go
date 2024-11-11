package service

import (
	"encoding/json"

	"hot-coffee/internal/dal"
	"hot-coffee/models"
)

type InventoryService interface {
	AddInventoryItem(i models.InventoryItem) error
	RetrieveInventoryItems() ([]byte, error)
	RetrieveInventoryItem(id string) ([]byte, error)
	UpdateInventoryItem(id string, item models.InventoryItem) error
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

func ValidateItem(i models.InventoryItem) error {
	if i.IngredientID == "" {
		return ErrNotValidID
	}

	if i.Name == "" {
		return ErrNotValidName
	}

	if i.Quantity <= 0 {
		return ErrNotValidQuantity
	}

	if i.Unit == "" {
		return ErrNotValidUnit
	}

	return nil
}

func (s *inventoryService) AddInventoryItem(i models.InventoryItem) error {
	if exists, err := s.InventoryRepository.ItemExists(i); err != nil {
		return err
	} else if exists {
		return ErrNotUniqueID
	}

	// Item validation
	if err := ValidateItem(i); err != nil {
		return err
	}

	if _, err := s.InventoryRepository.AddItem(i); err != nil {
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

func (s *inventoryService) UpdateInventoryItem(id string, item models.InventoryItem) error {
	inventoryItems, err := s.InventoryRepository.GetAllItems()
	if err != nil {
		if err.Error() == "EOF" {
			return ErrNoItem
		}
		return err
	}

	// updating the slice
	isFound := false
	for i, item := range inventoryItems {
		if item.IngredientID == id {
			inventoryItems[i] = item
			isFound = true
			break
		}
	}

	if !isFound {
		return ErrNoItem
	}

	err = s.InventoryRepository.SaveItems(inventoryItems)
	if err != nil {
		return err
	}

	return nil
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

	err = s.InventoryRepository.SaveItems(inventoryItems)
	if err != nil {
		return err
	}

	return nil
}
