package server

import (
	"fmt"
	"net/http"
)

func (s *Server) registerRoutes() {
	// // basic routes
	// s.mux.HandleFunc("GET /health", s.HandleHealth)

	// // bucket routes
	// s.mux.HandleFunc("GET /{BucketName}", s.HandleGetBucket)
	// s.mux.HandleFunc("GET /", s.HandleListBuckets)
	// s.mux.HandleFunc("PUT /{BucketName}", s.HandleCreateBucket)
	// s.mux.HandleFunc("DELETE /{BucketName}", s.HandleDeleteBucket)

	// // object routes
	// s.mux.HandleFunc("PUT /{BucketName}/{ObjectKey}", s.HandlePutObject)
	// s.mux.HandleFunc("GET /{BucketName}/{ObjectKey}", s.HandleGetObject)
	// s.mux.HandleFunc("DELETE /{BucketName}/{ObjectKey}", s.HandleDeleteObject)
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
