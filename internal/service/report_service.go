package service

import (
	"fmt"

	"hot-coffee/internal/dal"
	"hot-coffee/models"
)

type ReportService interface {
	GetTotalSales() (float64, error)
	GetPopularItems() ([]models.MenuItem, error)
}

type reportService struct {
	orderRepository     dal.OrderRepository
	menuReposipory      dal.MenuRepository
	inventoryRepository dal.InventoryRepository
}

func NewReportService(o dal.OrderRepository, m dal.MenuRepository, i dal.InventoryRepository) *reportService {
	if o == nil || m == nil || i == nil {
		return nil
	}
	return &reportService{orderRepository: o, menuReposipory: m, inventoryRepository: i}
}

// TODO: Write comments
func (rs *reportService) GetTotalSales() (float64, error) {
	var totalSales float64

	// Retriving successfull (closed) orders from the repo
	orders, err := rs.orderRepository.GetOrdersByStatus("closed")
	if err != nil {
		return totalSales, err
	}

	// Making a slice of ordered items
	menuItems, err := rs.menuReposipory.GetAllMenuItems()
	if err != nil {
		return totalSales, err
	}
	menuMap := make(map[string]models.MenuItem)
	for _, item := range menuItems {
		menuMap[item.ID] = item
	}

	// Iterating and summing the prices of ordered items
	for _, order := range orders {
		for _, item := range order.Items {
			menuItem, exists := menuMap[item.ProductID]
			if !exists {
				return totalSales, fmt.Errorf("menu item with ID %s not found", item.ProductID)
			}
			totalSales += menuItem.Price * float64(item.Quantity)
		}
	}

	return totalSales, nil
}

// TODO: Write comments

func (rs *reportService) GetPopularItems() ([]models.MenuItem, error) {
	// TODO:  Implement logic for popular items

	return []models.MenuItem{}, nil
}
