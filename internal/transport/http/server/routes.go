package server

import "coffee-shop/internal/transport/http/handler"

// endpoint prefix patterns
const inventoryPrefix = "/inventory"

func (s *Server) registerRoutes() {
	// Registering inventory routes
	s.setupInventoryRoutes()

	// // Registering  menu routes
	// s.registerMenuRoutes()

	// // Registering ordeer routes
	// s.registerOrderRoutes()

	// //  Registering report routes
	// s.registerReportRoutes()
}

func (s *Server) setupInventoryRoutes(handler *handler.InventoryHandler) {
	s.r.POST(inventoryPrefix, handler.AddInventoryItem)
	s.r.GET(inventoryPrefix, handler.GetInventoryItems)
	s.r.GET(inventoryPrefix+"/{id}", handler.GetInventoryItem)
	s.r.PUT(inventoryPrefix+"/{id}", handler.UpdateInventoryItem)
	s.r.DELETE(inventoryPrefix+"/{id}", handler.DeleteInventoryItem)
}

// func (s *Server) registerMenuRoutes() {
// 	// Interfaces
// 	menuRepository := repository.NewMenuRepository(s.config.menu_file)
// 	if menuRepository == nil {
// 		s.log.Error("Failed to create menu repository")
// 	}

// 	menuService := service.NewMenuService(menuRepository)
// 	if menuService == nil {
// 		s.log.Error("Failed to create menu service")
// 	}

// 	menuHandler := handler.NewMenuHandler(menuService, s.log)
// 	if menuHandler == nil {
// 		s.log.Error("Failed to create  handler")
// 	}

// 	// Routes
// 	s.mux.HandleFunc("POST /menu", menuHandler.AddMenuItem)
// 	s.mux.HandleFunc("GET /menu", menuHandler.GetMenuItems)
// 	s.mux.HandleFunc("GET /menu/{id}", menuHandler.GetMenuItem)
// 	s.mux.HandleFunc("PUT /menu/{id}", menuHandler.UpdateMenuItem)
// 	s.mux.HandleFunc("DELETE /menu/{id}", menuHandler.DeleteMenuItem)

// 	// logging
// 	s.log.Info("Menu routes is registered successfully")
// }

// func (s *Server) registerOrderRoutes() {
// 	// Orders
// 	orderRepository := repository.NewOrderRepository(s.config.order_file)
// 	if orderRepository == nil {
// 		s.log.Warn("Failed to create order repository")
// 	}

// 	menuRepository := repository.NewMenuRepository(s.config.menu_file)
// 	if menuRepository == nil {
// 		s.log.Error("Failed to create menu repository")
// 	}

// 	inventoryRepository := repository.NewInventoryRepository(s.config.inventory_file)
// 	if inventoryRepository == nil {
// 		s.log.Warn("Failed to create inventory repository")
// 	}

// 	reportRepository := repository.NewReportRepository(s.config.report_file)
// 	if reportRepository == nil {
// 		s.log.Warn("Failed to create report repository")
// 	}

// 	orderService := service.NewOrderService(orderRepository, menuRepository, inventoryRepository, reportRepository)
// 	if orderService == nil {
// 		s.log.Warn("Failed to create order service")
// 	}

// 	orderHandler := handler.NewOrderHandler(orderService, s.log)
// 	if orderHandler == nil {
// 		s.log.Warn("Failed to create order handler")
// 	}

// 	// Order routes
// 	s.mux.HandleFunc("POST /orders", orderHandler.CreateOrder)
// 	s.mux.HandleFunc("GET /orders", orderHandler.RetrieveOrders)
// 	s.mux.HandleFunc("GET /orders/{id}", orderHandler.RetrieveOrder)
// 	s.mux.HandleFunc("PUT /orders/{id}", orderHandler.UpdateOrder)
// 	s.mux.HandleFunc("DELETE /orders/{id}", orderHandler.DeleteOrder)
// 	s.mux.HandleFunc("POST /orders/{id}/close", orderHandler.CloseOrder)

// 	// logging
// 	s.log.Info("Order routes is registered successfully")
// }

// func (s *Server) registerReportRoutes() {
// 	// Interfaces
// 	orderRepository := repository.NewOrderRepository(s.config.order_file)
// 	if orderRepository == nil {
// 		s.log.Warn("Failed to create order repository")
// 	}

// 	menuRepository := repository.NewMenuRepository(s.config.menu_file)
// 	if menuRepository == nil {
// 		s.log.Error("Failed to create menu repository")
// 	}

// 	inventoryRepository := repository.NewInventoryRepository(s.config.inventory_file)
// 	if inventoryRepository == nil {
// 		s.log.Warn("Failed to create inventory repository")
// 	}

// 	reportRepository := repository.NewReportRepository(s.config.report_file)
// 	if reportRepository == nil {
// 		s.log.Warn("Failed to create report repository")
// 	}

// 	reportService := service.NewReportService(orderRepository, menuRepository, inventoryRepository, reportRepository)
// 	if reportService == nil {
// 		s.log.Warn("Failed to create report service")
// 	}

// 	reportHandler := handler.NewReportHandler(reportService, s.log)
// 	if reportService == nil {
// 		s.log.Warn("Failed to create report handler")
// 	}

// 	// Aggregation routes
// 	s.mux.HandleFunc("GET /reports/total-sales", reportHandler.GetTotalSales)
// 	s.mux.HandleFunc("GET /reports/popular-items", reportHandler.GetPopularItems)

// 	// logging
// 	s.log.Info("Report routes is registered successfully")
// }
