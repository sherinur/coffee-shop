package service

import (
	"encoding/json"
	"strings"
	"time"

	"hot-coffee/internal/dal"
	"hot-coffee/models"
)

type OrderService interface {
	AddOrder(o models.Order) error
	RetrieveOrders() ([]byte, error)
	RetrieveOrder(id string) ([]byte, error)
	UpdateOrder(id string, item models.Order) error
	DeleteOrder(id string) error
	CloseOrder(id string) error
	IsInventorySufficient(orderItems []models.OrderItem) (bool, error)
}

type orderService struct {
	OrderRepository     dal.OrderRepository
	InventoryRepository dal.InventoryRepository
}

func NewOrderService(or dal.OrderRepository, ir dal.InventoryRepository) *orderService {
	if or == nil || ir == nil {
		return nil
	}
	return &orderService{OrderRepository: or, InventoryRepository: ir}
}

func ValidateOrder(o models.Order) error {
	if o.ID == "" || strings.Contains(o.ID, " ") {
		return ErrNotValidOrderID
	}

	if o.CustomerName == "" {
		return ErrNotValidOrderCustomerName
	}

	err := ValidateOrderItems(o.Items)
	if err != nil {
		return err
	}

	if o.Status != "" {
		return ErrNotValidStatusField
	}

	if o.CreatedAt != "" {
		return ErrNotValidCreatedAt
	}

	return nil
}

func ValidateOrderItems(items []models.OrderItem) error {
	if items == nil || len(items) < 1 {
		return ErrNotValidOrderItems
	}

	for k, item := range items {
		for l, item2 := range items {
			if item.ProductID == item2.ProductID && k != l {
				return ErrDuplicateOrderItems
			}
		}
		if item.ProductID == "" || strings.Contains(item.ProductID, " ") {
			return ErrNotValidIngredientID
		}

		if item.Quantity < 1 {
			return ErrNotValidQuantity
		}
	}
	return nil
}

func (s *orderService) AddOrder(o models.Order) error {
	if exists, err := s.OrderRepository.OrderExists(o); err != nil {
		return err
	} else if exists {
		return ErrNotUniqueOrder
	}

	enoughItems, err := s.IsInventorySufficient(o.Items)
	if err != nil || !enoughItems {
		return err
	}

	// Order validation
	if err := ValidateOrder(o); err != nil {
		return err
	}

	o.Status = "Open"
	o.CreatedAt = time.Now().Format(time.RFC3339)

	if _, err := s.OrderRepository.AddOrder(o); err != nil {
		return err
	}
	return nil
}

func (s *orderService) RetrieveOrders() ([]byte, error) {
	orders, err := s.OrderRepository.GetAllOrders()
	if err != nil {
		return nil, err
	}

	data, err := json.MarshalIndent(orders, "", " ")
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s *orderService) RetrieveOrder(id string) ([]byte, error) {
	var order models.Order
	order, err := s.OrderRepository.GetOrderById(id)
	if err != nil {
		if err.Error() == "EOF" {
			return nil, ErrNoOrder
		}
		return nil, err
	}

	data, err := json.MarshalIndent(order, "", " ")
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s *orderService) UpdateOrder(id string, order models.Order) error {
	// TODO: Validate the order update
	// TODO: Call RewriteOrder method from repository ->  err := s.OrderRepository.RewriteOrder(id, order)
	return nil
}

func (s *orderService) DeleteOrder(id string) error {
	return s.OrderRepository.DeleteOrderById(id)
}

func (s *orderService) CloseOrder(id string) error {
	// TODO: Когда заказ закрывается через /orders/{id}/close, система считает, что заказ выполнен, и обновляет инвентарь, вычитая количество ингредиентов, необходимых для его выполнения.
	// TODO: После успешного вычитания ингредиентов заказ считается закрытым( "status": "open", -> "status": "closed",), и он больше не будет доступен для изменений (Изменить Update, проверять статус closed or open).
	// ? TODO: Закрытие также означает, что заказ включается в итоговую статистику для расчетов выручки и популярных позиций.

	// TODO: Call UpdateOrder or DeleteOrder from repository if needed
	return nil
}

func (s *orderService) IsInventorySufficient(orderItems []models.OrderItem) (bool, error) {
	inventoryItem, err := s.InventoryRepository.GetAllItems()
	if err != nil {
		return false, err
	}

	for _, orderItem := range orderItems {
		exists, err := s.InventoryRepository.ItemExistsById(orderItem.ProductID)
		if err != nil {
			return false, err
		}

		if !exists {
			return false, ErrOrderProductNotFound
		}

		for _, inventoryItem := range inventoryItem {
			if orderItem.ProductID == inventoryItem.IngredientID {
				if orderItem.Quantity > int(inventoryItem.Quantity) {
					return false, ErrNotEnoughInventoryQuantity
				}
			}
		}
	}

	return true, nil
}
