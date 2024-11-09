package dal

import (
	"hot-coffee/models"
)

type OrderRepository interface {
	AddOrder(order models.Order) (models.Order, error)
	DeleteOrderById(id string) (models.Order, error)
	GetOrderById(id string) (models.Order, error)
	GetAllOrders() ([]models.Order, error)
}

type orderRepository struct {
	filePath string
}

func NewOrderRepository(filePath string) OrderRepository {
	return &orderRepository{filePath: filePath}
}

func (r *orderRepository) AddOrder(order models.Order) (models.Order, error) {
	// TODO: getAllOrders and append new order
	// TODO: Marshal orders to JSON and write new file(ioutil.WriteFile, 0644)
	return order, nil
}

func (r *orderRepository) DeleteOrderById(id string) (models.Order, error) {
	// TODO: Get all orders, find order by id and delete it
	// TODO: Marshal orders to JSON and write new file(ioutil.WriteFile, 0644
	// TODO: Return the deleted order
	return models.Order{}, nil
}

func (r *orderRepository) GetOrderById(id string) (models.Order, error) {
	// TODO: getAllOrders and search for order by id
	// TODO: Unmarshal JSON to Order
	return models.Order{}, nil
}

func (r *orderRepository) GetAllOrders() ([]models.Order, error) {
	// TODO: open and read file.  return error if file is not found
	// TODO: decode  json and return slice of orders

	return []models.Order{}, nil
}
