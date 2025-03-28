package service

import (
	"strings"
	"time"

	"coffee-shop/internal/repository"
	"coffee-shop/models"
)

type OrderService interface {
	AddOrder(o models.Order) error
	RetrieveOrders() ([]models.Order, error)
	RetrieveOrder(id string) (models.Order, error)
	UpdateOrder(id string, item models.Order) error
	DeleteOrder(id string) error
	CloseOrder(id string) error
	IsInventorySufficient(orderItems []models.OrderItem) (bool, error)
	ReduceIngredients(orderItems []models.OrderItem) error
	CalculateTotalSales() (float64, error)
}

type orderService struct {
	OrderRepository     repository.OrderRepository
	MenuRepository      repository.MenuRepository
	InventoryRepository repository.InventoryRepository
	ReportRepository    repository.ReportRepository
}

func NewOrderService(or repository.OrderRepository, menu repository.MenuRepository, ir repository.InventoryRepository, re repository.ReportRepository) *orderService {
	if or == nil || ir == nil {
		return nil
	}
	return &orderService{OrderRepository: or, MenuRepository: menu, InventoryRepository: ir, ReportRepository: re}
}

func ValidateOrder(o models.Order) error {
	if strings.Contains(o.ID, " ") {
		return ErrNotValidOrderID
	}

	if o.CustomerName == "" {
		return ErrNotValidOrderCustomerName
	}

	err := ValidateOrderItems(o.Items)
	if err != nil {
		return err
	}

	if o.Status != "" || o.Status == "closed" {
		return ErrNotValidStatusField
	}

	if o.CreatedAt != "" {
		return ErrNotValidCreatedAt
	}

	return nil
}

func ValidateOrderItems(items []models.OrderItem) error {
	if len(items) < 1 {
		return ErrNotValidOrderItems
	}

	for k, item := range items {
		if item.ProductID == "" || strings.Contains(item.ProductID, " ") {
			return ErrNotValidIngredientID
		}

		for l, item2 := range items {
			if item.ProductID == item2.ProductID && k != l {
				return ErrDuplicateOrderItems
			}
		}

		if item.Quantity < 1 {
			return ErrNotValidQuantity
		}
	}
	return nil
}

func (s *orderService) AddOrder(order models.Order) error {
	if exists, err := s.OrderRepository.OrderExists(order); err != nil {
		return err
	} else if exists {
		return ErrNotUniqueOrder
	}

	_, err := s.IsInventorySufficient(order.Items)
	if err != nil {
		return err
	}

	// Order validation
	if err := ValidateOrder(order); err != nil {
		return err
	}

	order.Status = "Open"
	order.CreatedAt = time.Now().Format(time.RFC3339)

	if _, err := s.OrderRepository.AddOrder(order); err != nil {
		return err
	}
	return nil
}

func (s *orderService) RetrieveOrders() ([]models.Order, error) {
	return s.OrderRepository.GetAllOrders()
}

func (s *orderService) RetrieveOrder(id string) (models.Order, error) {
	if len(id) == 0 {
		return models.Order{}, ErrNotValidOrderID
	}

	order, err := s.OrderRepository.GetOrderById(id)
	if err != nil {
		return models.Order{}, err
	}

	return order, nil
}

func (s *orderService) UpdateOrder(id string, order models.Order) error {
	if len(id) == 0 {
		return ErrNotValidOrderID
	}

	if err := ValidateOrder(order); err != nil {
		return err
	}
	if order.ID == "" {
		order.ID = id
	}
	order.Status = "open"

	err := s.OrderRepository.RewriteOrder(id, order)
	if err != nil {
		return err
	}

	return nil
}

func (s *orderService) DeleteOrder(id string) error {
	if len(id) == 0 {
		return ErrNotValidOrderID
	}

	return s.OrderRepository.DeleteOrderById(id)
}

func (s *orderService) CloseOrder(id string) error {
	if len(id) == 0 {
		return ErrNotValidOrderID
	}

	order, err := s.OrderRepository.GetOrderById(id)
	if err != nil {
		return err
	}

	if order.Status != "open" {
		return ErrOrderClosed
	}

	err = s.ReduceIngredients(order.Items)
	if err != nil {
		return err
	}

	totalOrderPrice, err := s.CalculateTotalSales()
	if err != nil {
		return err
	}

	err = s.ReportRepository.UpdateTotalSales(totalOrderPrice)
	if err != nil {
		return err
	}

	order.Status = "closed"

	err = s.OrderRepository.RewriteOrder(id, order)
	if err != nil {
		return err
	}

	return nil
}

func (s *orderService) IsInventorySufficient(orderItems []models.OrderItem) (bool, error) {
	inventoryMap := make(map[string]models.InventoryItem)
	inventoryItems, err := s.InventoryRepository.GetAllItems()
	if err != nil {
		return false, err
	}
	for _, item := range inventoryItems {
		inventoryMap[item.IngredientID] = item
	}

	existingOrders, err := s.OrderRepository.GetAllOrders()
	if err != nil {
		return false, err
	}

	for _, existingOrder := range existingOrders {
		if existingOrder.Status == "closed" {
			continue
		}
		for _, existingOrderItem := range existingOrder.Items {
			menuItem, err := s.MenuRepository.GetMenuItemById(existingOrderItem.ProductID)
			if err != nil {
				return false, err
			}

			for _, ingredient := range menuItem.Ingredients {
				inventoryItem, exists := inventoryMap[ingredient.IngredientID]
				if exists {
					reservedQuantity := ingredient.Quantity * float64(existingOrderItem.Quantity)
					inventoryItem.Quantity -= reservedQuantity
					inventoryMap[ingredient.IngredientID] = inventoryItem
				}
			}
		}
	}

	menuMap := make(map[string]models.MenuItem)
	menuItems, err := s.MenuRepository.GetAllMenuItems()
	if err != nil {
		return false, err
	}
	for _, item := range menuItems {
		menuMap[item.ID] = item
	}

	for _, orderItem := range orderItems {
		menuItem, exists := menuMap[orderItem.ProductID]
		if !exists {
			return false, ErrOrderProductNotFound
		}

		for _, ingredient := range menuItem.Ingredients {
			inventoryItem, exists := inventoryMap[ingredient.IngredientID]
			if !exists {
				return false, ErrInventoryItemNotFound
			}

			requiredQuantity := ingredient.Quantity * float64(orderItem.Quantity)
			if requiredQuantity > inventoryItem.Quantity {
				return false, ErrNotEnoughInventoryQuantity
			}
		}
	}

	return true, nil
}

func (s *orderService) ReduceIngredients(orderItems []models.OrderItem) error {
	inventoryMap := make(map[string]models.InventoryItem)
	inventoryItems, err := s.InventoryRepository.GetAllItems()
	if err != nil {
		return err
	}

	for _, item := range inventoryItems {
		inventoryMap[item.IngredientID] = item
	}

	menuMap := make(map[string]models.MenuItem)
	menuItems, err := s.MenuRepository.GetAllMenuItems()
	if err != nil {
		return err
	}
	for _, item := range menuItems {
		menuMap[item.ID] = item
	}

	for _, orderItem := range orderItems {
		menuItem, exists := menuMap[orderItem.ProductID]
		if !exists {
			return ErrOrderProductNotFound
		}

		for _, ingredient := range menuItem.Ingredients {
			inventoryItem, exists := inventoryMap[ingredient.IngredientID]
			if !exists {
				return ErrInventoryItemNotFound
			}

			requiredQuantity := ingredient.Quantity * float64(orderItem.Quantity)
			if requiredQuantity > inventoryItem.Quantity {
				return ErrNotEnoughInventoryQuantity
			}

			inventoryItem.Quantity -= requiredQuantity
			inventoryMap[ingredient.IngredientID] = inventoryItem
		}
	}

	var updatedItems []models.InventoryItem
	for _, item := range inventoryMap {
		updatedItems = append(updatedItems, item)
	}

	if err := s.InventoryRepository.SaveItems(updatedItems); err != nil {
		return err
	}

	return nil
}

func (s *orderService) CalculateTotalSales() (float64, error) {
	totalSales := 0.0

	orders, err := s.OrderRepository.GetAllOrders()
	if err != nil {
		return 0.0, err
	}

	for _, order := range orders {
		for _, orderItem := range order.Items {
			menuItem, err := s.MenuRepository.GetMenuItemById(orderItem.ProductID)
			if err != nil {
				return 0.0, err
			}

			itemTotal := menuItem.Price * float64(orderItem.Quantity)
			totalSales += itemTotal
		}
	}

	return totalSales, nil
}
