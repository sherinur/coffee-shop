package service

import (
	"context"

	"coffee-shop/internal/model"
	"coffee-shop/internal/repository/postgres"
)

type orderService struct {
	OrderRepo     postgres.Order
	MenuRepo      postgres.Menu
	InventoryRepo postgres.Inventory
	ReportRepo    any
}

func NewOrderService(or postgres.Order, menu postgres.Menu, ir postgres.Inventory, re any) *orderService {
	return &orderService{OrderRepo: or, MenuRepo: menu, InventoryRepo: ir, ReportRepo: re}
}

func (s *orderService) AddOrder(ctx context.Context, order model.Order) error {
	if err := order.Validate(); err != nil {
		return err
	}

	if err := s.OrderRepo.Create(ctx, order); err != nil {
		return err
	}

	return nil
}

func (s *orderService) RetrieveOrders(ctx context.Context) ([]model.Order, error) {
	return s.OrderRepo.GetAll(ctx)
}

func (s *orderService) RetrieveOrder(ctx context.Context, id int) (*model.Order, error) {
	order, err := s.OrderRepo.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return &order, nil
}

func (s *orderService) UpdateOrder(ctx context.Context, id int, order model.Order) error {
	if err := order.Validate(); err != nil {
		return err
	}

	err := s.OrderRepo.Update(ctx, id, order)
	if err != nil {
		return err
	}

	return nil
}

func (s *orderService) DeleteOrder(ctx context.Context, id int) error {
	return s.OrderRepo.Delete(ctx, id)
}

func (s *orderService) CloseOrder(ctx context.Context, id int) error {

	// if order.Status != "open" {
	// 	return ErrOrderClosed
	// }

	// err = s.ReduceIngredients(order.Items)
	// if err != nil {
	// 	return err
	// }

	// totalOrderPrice, err := s.CalculateTotalSales()
	// if err != nil {
	// 	return err
	// }

	// err = s.ReportRepo.UpdateTotalSales(totalOrderPrice)
	// if err != nil {
	// 	return err
	// }

	// order.Status = "closed"

	// err = s.OrderRepo.RewriteOrder(id, order)
	// if err != nil {
	// 	return err
	// }

	return nil
}

// func (s *orderService) IsInventorySufficient(ctx context.Context, orderItems []model.OrderItem) (bool, error) {
// 	inventoryMap := make(map[string]model.InventoryItem)
// 	inventoryItems, err := s.InventoryRepo.GetAllItems()
// 	if err != nil {
// 		return false, err
// 	}
// 	for _, item := range inventoryItems {
// 		inventoryMap[item.IngredientID] = item
// 	}

// 	existingOrders, err := s.OrderRepo.GetAllOrders()
// 	if err != nil {
// 		return false, err
// 	}

// 	for _, existingOrder := range existingOrders {
// 		if existingOrder.Status == "closed" {
// 			continue
// 		}
// 		for _, existingOrderItem := range existingOrder.Items {
// 			menuItem, err := s.MenuRepo.GetMenuItemById(existingOrderItem.ProductID)
// 			if err != nil {
// 				return false, err
// 			}

// 			for _, ingredient := range menuItem.Ingredients {
// 				inventoryItem, exists := inventoryMap[ingredient.IngredientID]
// 				if exists {
// 					reservedQuantity := ingredient.Quantity * float64(existingOrderItem.Quantity)
// 					inventoryItem.Quantity -= reservedQuantity
// 					inventoryMap[ingredient.IngredientID] = inventoryItem
// 				}
// 			}
// 		}
// 	}

// 	menuMap := make(map[string]models.MenuItem)
// 	menuItems, err := s.MenuRepo.GetAllMenuItems()
// 	if err != nil {
// 		return false, err
// 	}
// 	for _, item := range menuItems {
// 		menuMap[item.ID] = item
// 	}

// 	for _, orderItem := range orderItems {
// 		menuItem, exists := menuMap[orderItem.ProductID]
// 		if !exists {
// 			return false, ErrOrderProductNotFound
// 		}

// 		for _, ingredient := range menuItem.Ingredients {
// 			inventoryItem, exists := inventoryMap[ingredient.IngredientID]
// 			if !exists {
// 				return false, ErrInventoryItemNotFound
// 			}

// 			requiredQuantity := ingredient.Quantity * float64(orderItem.Quantity)
// 			if requiredQuantity > inventoryItem.Quantity {
// 				return false, ErrNotEnoughInventoryQuantity
// 			}
// 		}
// 	}

// 	return true, nil
// }

// func (s *orderService) ReduceIngredients(ctx context.Context, orderItems []model.OrderItem) error {
// 	inventoryMap := make(map[string]models.InventoryItem)
// 	inventoryItems, err := s.InventoryRepo.GetAllItems()
// 	if err != nil {
// 		return err
// 	}

// 	for _, item := range inventoryItems {
// 		inventoryMap[item.IngredientID] = item
// 	}

// 	menuMap := make(map[string]models.MenuItem)
// 	menuItems, err := s.MenuRepo.GetAllMenuItems()
// 	if err != nil {
// 		return err
// 	}
// 	for _, item := range menuItems {
// 		menuMap[item.ID] = item
// 	}

// 	for _, orderItem := range orderItems {
// 		menuItem, exists := menuMap[orderItem.ProductID]
// 		if !exists {
// 			return ErrOrderProductNotFound
// 		}

// 		for _, ingredient := range menuItem.Ingredients {
// 			inventoryItem, exists := inventoryMap[ingredient.IngredientID]
// 			if !exists {
// 				return ErrInventoryItemNotFound
// 			}

// 			requiredQuantity := ingredient.Quantity * float64(orderItem.Quantity)
// 			if requiredQuantity > inventoryItem.Quantity {
// 				return ErrNotEnoughInventoryQuantity
// 			}

// 			inventoryItem.Quantity -= requiredQuantity
// 			inventoryMap[ingredient.IngredientID] = inventoryItem
// 		}
// 	}

// 	var updatedItems []models.InventoryItem
// 	for _, item := range inventoryMap {
// 		updatedItems = append(updatedItems, item)
// 	}

// 	if err := s.InventoryRepo.SaveItems(updatedItems); err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (s *orderService) CalculateTotalSales() (float64, error) {
// 	totalSales := 0.0

// 	orders, err := s.OrderRepo.GetAllOrders()
// 	if err != nil {
// 		return 0.0, err
// 	}

// 	for _, order := range orders {
// 		for _, orderItem := range order.Items {
// 			menuItem, err := s.MenuRepo.GetMenuItemById(orderItem.ProductID)
// 			if err != nil {
// 				return 0.0, err
// 			}

// 			itemTotal := menuItem.Price * float64(orderItem.Quantity)
// 			totalSales += itemTotal
// 		}
// 	}

// 	return totalSales, nil
// }
