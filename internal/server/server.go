package server

import (
	"net/http"

	"hot-coffee/pkg/logger"
)

type Server struct {
	config *Config
	logger *logger.Logger
	mux    *http.ServeMux
}

// New server
func New(config *Config) *Server {
	s := &Server{
		config: config,
		logger: logger.NewLogger(true, true),
		mux:    http.NewServeMux(),
	}

	s.registerRoutes()
	return s
}

// TODO: Продолжить по видео REST API на Golang

// Start the server
func (s *Server) Start() error {
	s.logger.PrintInfoMsg("Starting server on port " + s.config.port)
	s.logger.PrintInfoMsg("Path to the directory set: " + s.config.data_directory)
	s.logger.PrintInfoMsg("Path to the config set: " + s.config.cfg_file)

	mux := s.RequestMiddleware(s.mux)

	return http.ListenAndServe(s.config.port, mux)
}
