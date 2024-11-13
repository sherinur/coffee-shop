package server

import (
	"fmt"
	"net/http"

	"hot-coffee/internal/dal"
	"hot-coffee/internal/handler"
	"hot-coffee/internal/service"
)

func (s *Server) registerRoutes() {
	// // basic routes
	// s.mux.HandleFunc("GET /health", s.HandleHealth)

	// ! 1)	Конфликт имен для интерфейса и структуры
	// !  	Ошибка: Название интерфейсов и структур одинаковое
	// !	(например, InventoryService и inventoryService). Это может быть запутывающим.
	// ! 	Рекомендация: Используйте префиксы или более уникальные имена для интерфейсов,
	// ! 	например, InventoryService (интерфейс) и InventoryServiceImpl (структура) или InventoryRepo и inventoryRepositoryImpl.

	// ! 2) Неиспользование интерфейсов в handler
	// ! 	Ошибка: В inventoryHandler интерфейс InventoryService используется напрямую вместо указателя на service.InventoryService.
	// !	Решение: Убедитесь, что InventoryService является интерфейсом,
	// ! 	а не указателем на конкретную реализацию, чтобы сохранить гибкость в тестировании и подмене реализации.

	// TODO: Refactor the architecture of route registering (Decouple registering for functions)
	/*
		func (s *Server) registerRoutes() {
			s.mux.HandleFunc("GET /health", s.HandleHealth)

			s.registerInventoryRoutes()

			s.registerMenuRoutes()

			s.registerOrderRoutes()

			s.registerReportRoutes()
		}
	*/

	// Inventory
	inventoryRepository := dal.NewInventoryRepository(s.config.data_directory + "/inventory.json")
	if inventoryRepository == nil {
		s.logger.PrintWarnMsg("Failed to create inventory repository")
	}

	inventoryService := service.NewInventoryService(inventoryRepository)
	if inventoryService == nil {
		s.logger.PrintWarnMsg("Failed to create inventory service")
	}

	inventoryHandler := handler.NewInventoryHandler(inventoryService, s.logger)
	if inventoryHandler == nil {
		s.logger.PrintWarnMsg("Failed to create inventory handler")
	}

	// Inventory routes
	s.mux.HandleFunc("POST /inventory", inventoryHandler.AddInventoryItem)
	s.mux.HandleFunc("GET /inventory", inventoryHandler.GetInventoryItems)
	s.mux.HandleFunc("GET /inventory/{id}", inventoryHandler.GetInventoryItem)
	s.mux.HandleFunc("PUT /inventory/{id}", inventoryHandler.UpdateInventoryItem)
	s.mux.HandleFunc("DELETE /inventory/{id}", inventoryHandler.DeleteInventoryItem)

	// Menu
	menuRepository := dal.NewMenuRepository(s.config.data_directory + "/menu_items.json")
	if menuRepository == nil {
		s.logger.PrintErrorMsg("Failed to create menu repository")
	}

	menuService := service.NewMenuService(menuRepository)
	if menuService == nil {
		s.logger.PrintErrorMsg("Failed to create menu service")
	}

	menuHandler := handler.NewMenuHandler(menuService, s.logger)
	if menuHandler == nil {
		s.logger.PrintErrorMsg("Failed to create  handler")
	}

	// Menu routes
	s.mux.HandleFunc("POST /menu", menuHandler.AddMenuItem)
	s.mux.HandleFunc("GET /menu", menuHandler.GetMenuItems)
	s.mux.HandleFunc("GET /menu/{id}", menuHandler.GetMenuItem)
	s.mux.HandleFunc("PUT /menu/{id}", menuHandler.UpdateMenuItem)
	s.mux.HandleFunc("DELETE /menu/{id}", menuHandler.DeleteMenuItem)

	// Orders
	orderRepository := dal.NewOrderRepository(s.config.data_directory + "/orders.json")
	if orderRepository == nil {
		s.logger.PrintWarnMsg("Failed to create order repository")
	}

	orderService := service.NewOrderService(orderRepository)
	if inventoryService == nil {
		s.logger.PrintWarnMsg("Failed to create order service")
	}

	orderHandler := handler.NewOrderHandler(orderService, s.logger)
	if inventoryHandler == nil {
		s.logger.PrintWarnMsg("Failed to create order handler")
	}

	// Order routes
	s.mux.HandleFunc("POST /orders", orderHandler.CreateOrder)
	s.mux.HandleFunc("GET /orders", orderHandler.RetrieveOrders)
	s.mux.HandleFunc("GET /orders/{id}", orderHandler.RetrieveOrder)
	s.mux.HandleFunc("PUT /orders/{id}", orderHandler.UpdateOrder)
	s.mux.HandleFunc("DELETE /orders/{id}", orderHandler.DeleteOrder)
	s.mux.HandleFunc("POST /orders/{id}/close", orderHandler.CloseOrder)

	// Aggregations
	reportService := service.NewReportService(orderRepository, menuRepository, inventoryRepository)
	if reportService == nil {
		s.logger.PrintWarnMsg("Failed to create report service")
	}

	reportHandler := handler.NewReportHandler(reportService, s.logger)
	if reportService == nil {
		s.logger.PrintWarnMsg("Failed to create report handler")
	}

	// Aggregation routes
	s.mux.HandleFunc("GET /reports/total-sales", reportHandler.GetTotalSales)
	s.mux.HandleFunc("GET /reports/popular-items", reportHandler.GetPopularItems)
}

func (s *Server) RequestMiddleware(next http.Handler) http.Handler {
	allowedMethods := map[string]bool{
		http.MethodGet:    true,
		http.MethodPost:   true,
		http.MethodPut:    true,
		http.MethodDelete: true,
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s.logger.PrintInfoMsg(fmt.Sprintf("Request %s %s from %s", r.Method, r.URL.Path, r.RemoteAddr))
		if !allowedMethods[r.Method] {
			return
		}

		next.ServeHTTP(w, r)
	})
}
