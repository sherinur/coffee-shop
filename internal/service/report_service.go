package service

import (
	"sort"

	"coffee-shop/internal/dal"
	"coffee-shop/models"
)

type ReportService interface {
	GetTotalSales() (models.TotalSales, error)
	GetPopularItems() ([]models.MenuItem, error)
}

type reportService struct {
	orderRepository     dal.OrderRepository
	menuReposipory      dal.MenuRepository
	inventoryRepository dal.InventoryRepository
	reportRepository    dal.ReportRepository
}

func NewReportService(o dal.OrderRepository, m dal.MenuRepository, i dal.InventoryRepository, r dal.ReportRepository) *reportService {
	if o == nil || m == nil || i == nil || r == nil {
		return nil
	}
	return &reportService{orderRepository: o, menuReposipory: m, inventoryRepository: i, reportRepository: r}
}

func (rs *reportService) GetTotalSales() (models.TotalSales, error) {
	return rs.reportRepository.GetTotalSales()
}

func (rs *reportService) GetPopularItems() ([]models.MenuItem, error) {
	orders, err := rs.orderRepository.GetClosedOrders()
	if err != nil {
		return nil, err
	}

	frequencyMap := make(map[string]int)
	for _, order := range orders {
		for _, item := range order.Items {
			frequencyMap[item.ProductID] += item.Quantity
		}
	}

	type menuItemCount struct {
		ID    string
		Count int
	}

	menuItemsCount := []menuItemCount{}
	for id, count := range frequencyMap {
		menuItemsCount = append(menuItemsCount, menuItemCount{ID: id, Count: count})
	}

	sort.Slice(menuItemsCount, func(i, j int) bool {
		return menuItemsCount[i].Count > menuItemsCount[j].Count
	})

	topMenuItemsCount := menuItemsCount
	if len(menuItemsCount) > 10 {
		topMenuItemsCount = menuItemsCount[:10]
	}

	popularItems := []models.MenuItem{}
	for _, menuItem := range topMenuItemsCount {
		item, err := rs.menuReposipory.GetMenuItemById(menuItem.ID)
		if err != nil {
			return nil, err
		}
		popularItems = append(popularItems, item)
	}

	return popularItems, nil
}
