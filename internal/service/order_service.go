package service

import (
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
}

type orderService struct {
	OrderRepository dal.OrderRepository
}

func NewOrderService(repo dal.OrderRepository) *orderService {
	if repo == nil {
		return nil
	}
	return &orderService{OrderRepository: repo}
}

func ValidateOrder(o models.Order) error {
	// TODO: Define and implement validation rules for order
	// TODO: Validate order items using ValidateOrderItems()
	return nil
}

func ValidateOrderItems(items []models.OrderItem) error {
	// TODO: Define and implement validation rules for order items
	return nil
}

func (s *orderService) AddOrder(o models.Order) error {
	if exists, err := s.OrderRepository.OrderExists(o); err != nil {
		return err
	} else if exists {
		return ErrNotUniqueOrder
	}

	// Order validation
	if err := ValidateOrder(o); err != nil {
		return err
	}

	if _, err := s.OrderRepository.AddOrder(o); err != nil {
		return err
	}

	// TODO: Validate the order
	// TODO: Call AddOrder method from repository
	return nil
}

func (s *orderService) RetrieveOrders() ([]byte, error) {
	// TODO: Call GetAllOrders from repository
	// TODO: Marshal orders to JSON and return
	return nil, nil
}

func (s *orderService) RetrieveOrder(id string) ([]byte, error) {
	// TODO: Call GetOrderById from repository
	// TODO: Marshal the order to JSON and return
	return nil, nil
}

func (s *orderService) UpdateOrder(id string, item models.Order) error {
	// TODO: Validate the order update
	// TODO: Call RewriteOrder method from repository
	return nil
}

func (s *orderService) DeleteOrder(id string) error {
	// TODO: Call DeleteOrderById from repository
	return nil
}

func (s *orderService) CloseOrder(id string) error {
	// TODO: Implement logic to close the order
	// TODO: Call UpdateOrder or DeleteOrder from repository if needed
	return nil
}
