package server

import (
	"fmt"
	"log"
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

	// Inventory
	inventoryRepository := dal.NewInventoryRepository(s.config.data_directory + "/inventory.json")
	if inventoryRepository == nil {
		s.logger.PrintErrorMsg("Failed to create inventory repository")
	}

	inventoryService := service.NewInventoryService(inventoryRepository)
	if inventoryService == nil {
		s.logger.PrintErrorMsg("Failed to create inventory service")
	}

	inventoryHandler := handler.NewInventoryHandler(inventoryService, s.logger)
	if inventoryHandler == nil {
		s.logger.PrintErrorMsg("Failed to create inventory handler")
	}

	s.mux.HandleFunc("POST /inventory", inventoryHandler.AddInventoryItem)
	s.mux.HandleFunc("GET /inventory", inventoryHandler.GetInventoryItems)
	s.mux.HandleFunc("GET /inventory/{id}", inventoryHandler.GetInventoryItem)
	s.mux.HandleFunc("PUT /inventory/{id}", inventoryHandler.UpdateInventoryItem)
	s.mux.HandleFunc("DELETE /inventory/{id}", inventoryHandler.DeleteInventoryItem)

	// Menu
	menuRepository := dal.NewMenuRepository(s.config.data_directory + "/menu.json")
	if menuRepository == nil {
		log.Fatal("Failed to create menu repository")
	}

	menuService := service.NewMenuService(menuRepository)
	if menuService == nil {
		log.Fatal("Failed to create menu service")
	}

	menuHandler := handler.NewMenuHandler(menuService, s.logger)
	if menuHandler == nil {
		log.Fatal("Failed to create  handler")
	}

	s.mux.HandleFunc("GET /menu", menuHandler.GetMenuItems)
	s.mux.HandleFunc("GET /menu/{id}", menuHandler.GetMenuItem)

	// Orders
	orderRepository := dal.NewOrderRepository(s.config.data_directory + "/orders.json")
	if orderRepository == nil {
		s.logger.PrintErrorMsg("Failed to create order repository")
	}

	orderService := service.NewOrderService(orderRepository)
	if inventoryService == nil {
		s.logger.PrintErrorMsg("Failed to create order service")
	}

	orderHandler := handler.NewOrderHandler(orderService, s.logger)
	if inventoryHandler == nil {
		s.logger.PrintErrorMsg("Failed to create order handler")
	}

	// Order routes
	s.mux.HandleFunc("POST /orders", orderHandler.CreateOrder)
	s.mux.HandleFunc("GET /orders", orderHandler.RetrieveOrders)
	s.mux.HandleFunc("GET /orders/{id}", orderHandler.RetrieveOrder)
	s.mux.HandleFunc("PUT /orders/{id}", orderHandler.UpdateOrder)
	s.mux.HandleFunc("DELETE /orders/{id}", orderHandler.DeleteOrder)
	s.mux.HandleFunc("POST /orders/{id}/close", orderHandler.CloseOrder)
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
