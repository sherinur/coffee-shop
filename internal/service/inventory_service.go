package service

import (
	"errors"
	"strings"

	"coffee-shop/internal/repository"
	"coffee-shop/models"
)

type InventoryService interface {
	AddInventoryItem(i models.InventoryItem) error
	RetrieveInventoryItems() ([]models.InventoryItem, error)
	RetrieveInventoryItem(id string) (*models.InventoryItem, error)
	UpdateInventoryItem(id string, item models.InventoryItem) error
	DeleteInventoryItem(id string) error
}

type inventoryService struct {
	InventoryRepository repository.InventoryRepository
}

func NewInventoryService(repo repository.InventoryRepository) *inventoryService {
	if repo == nil {
		return nil
	}
	return &inventoryService{InventoryRepository: repo}
}

// ValidateItem validates the fields of an InventoryItem.
// Returns nil if the item is valid.
// The following errors may be returned:
// - ErrNotValidID if the IngredientID is empty.
// - ErrIDContainsSpace if the IngredientID contains spaces.
// - ErrNotValidName if the Name is empty.
// - ErrNotValidQuantity if the Quantity is zero or negative.
// - ErrNotValidUnit if the Unit is empty.
func ValidateItem(i models.InventoryItem) error {
	if i.IngredientID == "" || strings.Contains(i.IngredientID, " ") {
		return ErrNotValidIngredientID
	}
	if strings.Contains(i.IngredientID, " ") {
		return ErrNotValidIngredientID
	}

	if i.Name == "" {
		return ErrNotValidIngredientName
	}

	if i.Quantity < 0 || i.Quantity == -0 {
		return ErrNotValidQuantity
	}

	if i.Unit == "" {
		return ErrNotValidUnit
	}

	return nil
}

// AddInventoryItem adds a new inventory item to the repository.
// Returns nil if the addition is successful.
// The following errors may be returned:
// - ErrNotUniqueID if the item with the same ID already exists.
// - An error if there is a validation issue or a failure when adding the item to the repository.
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

// RetrieveInventoryItems retrieves all inventory items from the repository.
// Returns the items data in JSON format as a byte slice.
// The following error may be returned:
// - An error if there is a failure when retrieving items from the repository or when marshalling the data.
func (s *inventoryService) RetrieveInventoryItems() ([]models.InventoryItem, error) {
	return s.InventoryRepository.GetAllItems()
}

// RetrieveInventoryItem retrieves a single inventory item by its ID.
// Returns the item data in JSON format as a byte slice if found.
// The following errors may be returned:
// - ErrNoItem if the item with the specified ID is not found.
// - An error if there is a failure when retrieving items from the repository or when marshalling the item data.
func (s *inventoryService) RetrieveInventoryItem(id string) (*models.InventoryItem, error) {
	if len(id) == 0 {
		return nil, errors.New("identificator is not valid")
	}

	inventoryItem, err := s.InventoryRepository.GetItemById(id)
	if err != nil {
		if err.Error() == "EOF" {
			return nil, ErrNoItem
		}
		return nil, err
	}

	return inventoryItem, nil
}

// UpdateInventoryItem updates the old inventory item with the new one.
// Returns nil if the update is successful.
// The following errors may be returned:
// - ErrNoItem if the old item is not found by id.
// - ErrNotUniqueID if new item id not unique.
// - An error if there is a validation issue or a failure when updating the repository.
func (s *inventoryService) UpdateInventoryItem(id string, i models.InventoryItem) error {
	if len(id) == 0 {
		return errors.New("identificator is not valid")
	}

	if exists, err := s.InventoryRepository.ItemExists(models.InventoryItem{IngredientID: id}); err != nil {
		return err
	} else if !exists {
		return ErrNoItem
	}

	// Uniqueness test of new item
	if i.IngredientID != id {
		if exists, err := s.InventoryRepository.ItemExists(i); err != nil {
			return err
		} else if exists {
			return ErrNotUniqueID
		}
	}

	// New item validation
	if err := ValidateItem(i); err != nil {
		return err
	}

	// Rewriting old item in repo
	err := s.InventoryRepository.RewriteItem(id, i)
	if err != nil {
		return nil
	}

	return nil
}

// DeleteInventoryItem deletes an inventory item by its ID.
// Returns nil if the deletion is successful.
// The following errors may be returned:
// - ErrNoItem if the item with the specified ID is not found.
// - An error if there is a failure when retrieving or saving items in the repository.
func (s *inventoryService) DeleteInventoryItem(id string) error {
	if len(id) == 0 {
		return errors.New("identificator is not valid")
	}
	return s.InventoryRepository.DeleteItemByID(id)
}
