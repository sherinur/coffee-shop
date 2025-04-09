package hp

import "coffee-shop/internal/transport/http/handler"

type Inventory interface {
	handler.InventoryService
	handler.MenuService
}
