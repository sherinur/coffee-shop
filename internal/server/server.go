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
func New(config *Config, LOGGER *logger.Logger) *Server {
	s := &Server{
		config: config,
		logger: LOGGER,
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

	// TODO: Провести проверку всех зависимостей (например, подключение к базе данных)
	// if !checkDependencies() {
	//     return fmt.Errorf("dependencies are not satisfied")
	// }

	//  TODO: Использовать http.Server для более гибкой настройки сервера (таймауты, shutdown)
	// server := &http.Server{
	// 	Addr:    s.config.port,
	// 	Handler: s.RequestMiddleware(s.mux),
	// TODO: Установить таймауты для соединений и запросов
	// ReadTimeout:  10 * time.Second,
	// WriteTimeout: 10 * time.Second,
	// IdleTimeout:  120 * time.Second,
	// }

	// TODO: Использовать горутины для асинхронного запуска сервера (обработка запросов)
	// go func() {
	// 	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
	// 		s.logger.PrintErrorMsg("Server error: " + err.Error())
	// 	}
	// }()

	// TODO: Обработать сигналы для корректного завершения работы (graceful shutdown)
	// signalChannel := make(chan os.Signal, 1)
	// signal.Notify(signalChannel, syscall.SIGINT, syscall.SIGTERM)
	// sig := <-signalChannel
	// s.logger.PrintInfoMsg(fmt.Sprintf("Received signal: %s. Shutting down...", sig))
	// return s.Shutdown(server)

	mux := s.RequestMiddleware(s.mux)

	return http.ListenAndServe(s.config.port, mux)
}

// Shutdown the server
func (s *Server) Shutdown() error {
	s.logger.PrintInfoMsg("Stopping the server")

	// TODO: Write server graceful shutdown method

	s.logger.PrintInfoMsg("Server gracefully stopped")
	return nil
}
