package dal

type InventoryRepository interface{}

type inventoryRepository struct {
	filePath string
}

func NewInventoryRepository(filePath string) InventoryRepository {
	return &inventoryRepository{filePath: filePath}
}
