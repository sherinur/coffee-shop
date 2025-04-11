package json

// type InventoryRepository interface {
// 	AddItem(i models.InventoryItem) (*models.InventoryItem, error)
// 	GetAllItems() ([]models.InventoryItem, error)
// 	GetItemById(id string) (*models.InventoryItem, error)
// 	SaveItems(inventoryItems []models.InventoryItem) error
// 	ItemExists(i models.InventoryItem) (bool, error)
// 	// ItemExistsById(id string) (bool, error)
// 	RewriteItem(id string, newItem models.InventoryItem) error
// 	DeleteItemByID(id string) error
// }

// type inventoryRepository struct {
// 	filePath string
// }

// func NewInventoryRepository(filePath string) *inventoryRepository {
// 	return &inventoryRepository{filePath: filePath}
// }

// // AddItem adds a new inventory item to the repository.
// // Returns the added item if successful.
// // The following errors may be returned:
// // - An error if there is a failure in retrieving or saving the items.
// func (r *inventoryRepository) AddItem(i models.InventoryItem) (*models.InventoryItem, error) {
// 	return nil, nil
// }

// // GetAllItems retrieves all inventory items from the repository.
// // Returns an empty slice if the file is empty or does not exist.
// // The following errors may be returned:
// // - An error if there is a failure in checking file existence or reading the file.
// func (r *inventoryRepository) GetAllItems() ([]models.InventoryItem, error) {
// 	return nil, nil
// }

// // GetItemById retrieves a specific inventory item by its ID.
// // Returns an error if the item is not found.
// func (r *inventoryRepository) GetItemById(id string) (*models.InventoryItem, error) {
// 	return nil, nil
// }

// // RewriteItem updates an existing inventory item identified by its ID.
// // Returns an error if updating the repository fails.
// func (r *inventoryRepository) RewriteItem(id string, newItem models.InventoryItem) error {
// 	return nil
// }

// // SaveItems writes the provided inventory items to the repository file.
// // Creates the directory and file if they do not exist.
// // The following errors may be returned:
// // - An error if creating the directory or file fails.
// // - An error if writing to the file fails.
// func (r *inventoryRepository) SaveItems(inventoryItems []models.InventoryItem) error {
// 	return nil
// }

// // ItemExists checks if an inventory item with the same ID already exists in the repository.
// // Returns true if the item exists, false otherwise.
// func (r *inventoryRepository) ItemExists(i models.InventoryItem) (bool, error) {
// 	return false, nil
// }

// // func (r *inventoryRepository) ItemExistsById(id string) (bool, error) {
// // 	inventoryItems, err := r.GetAllItems()
// // 	if err != nil {
// // 		return false, err
// // 	}

// // 	for _, item := range inventoryItems {
// // 		if item.IngredientID == id {
// // 			return true, nil
// // 		}
// // 	}

// // 	return false, nil
// // }

// // DeleteInventoryItem deletes an inventory item by its ID.
// // Returns nil if the deletion is successful.
// // The following errors may be returned:
// // - ErrNoItem if the item with the specified ID is not found.
// // - An error if there is a failure when retrieving or saving items in the repository.
// func (r *inventoryRepository) DeleteItemByID(id string) error {
// 	return nil
// }
