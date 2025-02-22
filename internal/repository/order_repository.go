package repository

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"coffee-shop/internal/utils"
	"coffee-shop/models"
)

type OrderRepository interface {
	AddOrder(order models.Order) (models.Order, error)
	GetAllOrders() ([]models.Order, error)
	GetOrdersByStatus(status string) ([]models.Order, error)
	GetClosedOrders() ([]models.Order, error)
	GetOpenOrders() ([]models.Order, error)
	GetOrderById(id string) (models.Order, error)
	DeleteOrderById(id string) error
	SaveOrders(orders []models.Order) error
	OrderExists(o models.Order) (bool, error)
	RewriteOrder(id string, newOrder models.Order) error
}

type orderRepository struct {
	filePath string
}

func NewOrderRepository(filePath string) *orderRepository {
	return &orderRepository{filePath: filePath}
}

func (r *orderRepository) AddOrder(order models.Order) (models.Order, error) {
	orders, err := r.GetAllOrders()
	if err != nil {
		return models.Order{}, err
	}

	ordersID := []string{}

	for _, order := range orders {
		ordersID = append(ordersID, order.ID)
	}

	if order.ID == "" {
		order.ID = utils.GenerateNewID(ordersID, "orders")
	}

	orders = append(orders, order)

	err = r.SaveOrders(orders)
	if err != nil {
		return models.Order{}, err
	}

	return order, nil
}

func (r *orderRepository) GetAllOrders() ([]models.Order, error) {
	orders := []models.Order{}

	exists, err := utils.FileExists(r.filePath)
	if err != nil {
		return []models.Order{}, err
	}
	if !exists {
		return []models.Order{}, nil
	}

	file, err := os.Open(r.filePath)
	if err != nil {
		return []models.Order{}, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)

	if utils.FileEmpty(file) {
		return []models.Order{}, nil
	}

	err = decoder.Decode(&orders)
	if err != nil {
		return []models.Order{}, err
	}

	return orders, nil
}

func (r *orderRepository) GetOrderById(id string) (models.Order, error) {
	orders, err := r.GetAllOrders()
	if err != nil {
		return models.Order{}, err
	}

	for _, order := range orders {
		if order.ID == id {
			return order, nil
		}
	}

	return models.Order{}, errors.New("order not found")
}

func (r *orderRepository) DeleteOrderById(id string) error {
	orders, err := r.GetAllOrders()
	if err != nil {
		if err.Error() == "EOF" {
			return errors.New("order not found")
		}
		return err
	}

	isFound := false
	for i, order := range orders {
		if order.ID == id {
			orders = append(orders[:i], orders[i+1:]...)
			isFound = true
			break
		}
	}

	if !isFound {
		return errors.New("order not found")
	}

	err = r.SaveOrders(orders)
	if err != nil {
		return err
	}

	return nil
}

func (r *orderRepository) SaveOrders(orders []models.Order) error {
	// Checking the existence of a directory for a file
	dir := filepath.Dir(r.filePath)
	err := utils.CreateDir(dir)
	if err != nil {
		return err
	}

	// Checking file write permissions
	exists, err := utils.FileExists(r.filePath)
	if err != nil {
		return err
	}

	if !exists {
		// If the file does not exist, create it
		err := utils.CreateFile(r.filePath)
		if err != nil {
			return fmt.Errorf("error creating file %s: %w", r.filePath, err)
		}
	}

	jsonData, err := json.MarshalIndent(orders, "", " ")
	if err != nil {
		return err
	}

	err = os.WriteFile(r.filePath, jsonData, 0o644)
	if err != nil {
		return err
	}

	return nil
}

func (r *orderRepository) OrderExists(o models.Order) (bool, error) {
	orders, err := r.GetAllOrders()
	if err != nil {
		return false, err
	}

	for _, order := range orders {
		if order.ID == o.ID {
			return true, nil
		}
	}

	return false, nil
}

func (r *orderRepository) RewriteOrder(id string, newOrder models.Order) error {
	orders, err := r.GetAllOrders()
	if err != nil {
		return err
	}

	for i, order := range orders {
		if order.ID == id {
			orders[i] = newOrder
			break
		}
	}

	err = r.SaveOrders(orders)
	if err != nil {
		return err
	}

	return nil
}

func (r *orderRepository) GetOrdersByStatus(status string) ([]models.Order, error) {
	orders, err := r.GetAllOrders()
	if err != nil {
		return []models.Order{}, err
	}

	closedOrders := []models.Order{}
	for _, order := range orders {
		if order.Status == status {
			closedOrders = append(closedOrders, order)
		}
	}

	return closedOrders, nil
}

func (r *orderRepository) GetClosedOrders() ([]models.Order, error) {
	return r.GetOrdersByStatus("closed")
}

func (r *orderRepository) GetOpenOrders() ([]models.Order, error) {
	return r.GetOrdersByStatus("open")
}
