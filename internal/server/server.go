package server

import (
	"hot-coffee/pkg/logger"
	"net/http"
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
		logger: logger.New(true),
		mux:    http.NewServeMux(),
	}

	s.registerRoutes()
	return s
}

// TODO: Продолжить по видео REST API на Golang

// Start the server
func (s *Server) Start() error {
	s.logger.PrintfInfoMsg("Starting server on port " + s.config.port)
	s.logger.PrintfInfoMsg("Path to the directory set: " + s.config.data_directory)
	s.logger.PrintfInfoMsg("Path to the config set: " + s.config.cfg_file)

	mux := s.RequestMiddleware(s.mux)

	return http.ListenAndServe(s.config.port, mux)
}
