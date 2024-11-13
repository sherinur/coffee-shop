package service

import (
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

func (rs *reportService) GetTotalSales() (float64, error) {
	var totalSales float64

	// orders, err := rs.orderRepository.GetAllOrders()
	// if err != nil {
	// 	return 0, err
	// }
	// for _, order := range orders {
	// 	for _, item := range order.Items {
	// 		totalSales += item.Price * float64(item.Quantity)
	// 	}
	// }
	// TODO: Implement logic for total sales
	// TODO: Write comments

	return totalSales, nil
}

func (rs *reportService) GetPopularItems() ([]models.MenuItem, error) {
	// TODO:  Implement logic for popular items
	// TODO: Write comments

	return []models.MenuItem{}, nil
}
