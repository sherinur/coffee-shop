package server

import (
	"fmt"
	"net/http"

	"hot-coffee/internal/handler"
)

func (s *Server) registerRoutes() {
	// // basic routes
	// s.mux.HandleFunc("GET /health", s.HandleHealth)

	// Order routes
	s.mux.HandleFunc("POST /orders", handler.CreateOrder)
	s.mux.HandleFunc("GET /orders", handler.RetrieveOrders)
	s.mux.HandleFunc("GET /orders/{id}", handler.RetrieveOrder)
	s.mux.HandleFunc("PUT /orders/{id}", handler.UpdateOrder)
	s.mux.HandleFunc("DELETE /orders/{id}", handler.DeleteOrder)
	s.mux.HandleFunc("POST	 /orders/{id}/close", handler.CloseOrder)

	// Menu routes

	// Inventory routes

	// Aggregation routes
}

func (s *Server) RequestMiddleware(next http.Handler) http.Handler {
	allowedMethods := map[string]bool{
		http.MethodGet:    true,
		http.MethodPost:   true,
		http.MethodPut:    true,
		http.MethodDelete: true,
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s.logger.PrintfInfoMsg(fmt.Sprintf("Request %s %s from %s", r.Method, r.URL.Path, r.RemoteAddr))

		if !allowedMethods[r.Method] {
			return
		}

		next.ServeHTTP(w, r)
	})
}
