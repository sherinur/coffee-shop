package service

import (
	"context"

	"coffee-shop/internal/model"
)

type inventoryService struct {
	InventoryRepo InventoryRepo
}

func NewInventoryService(repo InventoryRepo) *inventoryService {
	return &inventoryService{InventoryRepo: repo}
}

// AddInventoryItem adds a new inventory item to the repository.
// Returns nil if the addition is successful.
// The following errors may be returned:
// - ErrNotUniqueID if the item with the same ID already exists.
// - An error if there is a validation issue or a failure when adding the item to the repository.
func (s *inventoryService) AddInventoryItem(ctx context.Context, item model.Inventory) error {
	// Item validation
	if err := item.Validate(); err != nil {
		return err
	}

	err := s.InventoryRepo.Create(ctx, item)
	if err != nil {
		return err
	}

	return nil
}

// RetrieveInventoryItems retrieves all inventory items from the repository.
// Returns the items data in JSON format as a byte slice.
// The following error may be returned:
// - An error if there is a failure when retrieving items from the repository or when marshalling the data.
func (s *inventoryService) RetrieveInventoryItems(ctx context.Context) ([]model.Inventory, error) {
	return s.InventoryRepo.GetAll(ctx)
}

// RetrieveInventoryItem retrieves a single inventory item by its ID.
// Returns the item data in JSON format as a byte slice if found.
// The following errors may be returned:
// - ErrNoItem if the item with the specified ID is not found.
// - An error if there is a failure when retrieving items from the repository or when marshalling the item data.
func (s *inventoryService) RetrieveInventoryItem(ctx context.Context, id int) (*model.Inventory, error) {
	item, err := s.InventoryRepo.Get(ctx, id)
	return &item, err
}

// UpdateInventoryItem updates the old inventory item with the new one.
// Returns nil if the update is successful.
// The following errors may be returned:
// - ErrNoItem if the old item is not found by id.
// - ErrNotUniqueID if new item id not unique.
// - An error if there is a validation issue or a failure when updating the repository.
func (s *inventoryService) UpdateInventoryItem(ctx context.Context, id int, item model.Inventory) error {
	// New item validation
	if err := item.Validate(); err != nil {
		return err
	}

	// Rewriting old item in repo
	err := s.InventoryRepo.Update(ctx, id, item)
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
func (s *inventoryService) DeleteInventoryItem(ctx context.Context, id int) error {
	return s.InventoryRepo.Delete(ctx, id)
}
