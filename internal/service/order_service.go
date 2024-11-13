package service

import (
	"encoding/json"

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
	// TODO: Define and implement validation rules for order
	// TODO: Validate order items using ValidateOrderItems()
	return nil
}

func ValidateOrderItems(items []models.OrderItem) error {
	// TODO: Define and implement validation rules for order items

	// TODO: Добавить правило чтобы не повторялись продукты в массиве (один ингредиент и количество сразу пишутся)
	return nil
}

func (s *orderService) AddOrder(o models.Order) error {
	// TODO: Пересмотреть добавление проверки уникальности заказа
	// if exists, err := s.OrderRepository.OrderExists(o); err != nil {
	// 	return err
	// } else if exists {
	// 	return ErrNotUniqueOrder
	// }

	// TODO: проверить есть ли в инвентаре достаточное количество ингредиентов для всех позиций заказа
	// TODO: Если каких-то ингредиентов недостаточно, заказ не обрабатывается, и возвращается сообщение об ошибке с указанием недостаточных ингредиентов.

	// Order validation
	if err := ValidateOrder(o); err != nil {
		return err
	}

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
