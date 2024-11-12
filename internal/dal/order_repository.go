package dal

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"hot-coffee/internal/utils"
	"hot-coffee/models"
)

type OrderRepository interface {
	AddOrder(order models.Order) (models.Order, error)
	GetAllOrders() ([]models.Order, error)
	GetOrderById(id string) (models.Order, error)
	DeleteOrderById(id string) (models.Order, error)
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

	orders = append(orders, order)

	err = r.SaveOrders(orders)
	if err != nil {
		return models.Order{}, err
	}

	return order, nil
}

func (r *orderRepository) GetAllOrders() ([]models.Order, error) {
	orders := []models.InventoryItem{}

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

	return []models.Order{}, nil
}

func (r *orderRepository) GetOrderById(id string) (models.Order, error) {
	// TODO: getAllOrders and search for order by id
	// TODO: Unmarshal JSON to Order
	return models.Order{}, nil
}

func (r *orderRepository) DeleteOrderById(id string) (models.Order, error) {
	// TODO: Get all orders, find order by id and delete it
	// TODO: Marshal orders to JSON and write new file(ioutil.WriteFile, 0644)
	// TODO: Return the deleted order
	return models.Order{}, nil
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
	// TODO: Get all orders, find order by id and replace it with newOrder
	// TODO: Marshal orders to JSON and write new file(ioutil.WriteFile, 0644)
	return nil
}
