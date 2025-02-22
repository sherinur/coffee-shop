package server

import (
	"log/slog"

	"coffee-shop/internal/utils"
	"god"
)

type Server struct {
	config *Config
	log    *slog.Logger
	r      *god.Router
}

// New server
func New(config *Config, logger *slog.Logger) *Server {
	s := &Server{
		config: config,
		log:    logger,
		r:      god.Default(),
	}

	s.registerRoutes()
	return s
}

// Start the server
func (s *Server) Start() error {
	s.log.Info("Path to the directory set: " + s.config.data_directory)
	s.log.Info("Path to the config set: " + s.config.cfg_file)

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

	utils.CreateFile(s.config.inventory_file)
	utils.CreateFile(s.config.menu_file)
	utils.CreateFile(s.config.order_file)
	utils.CreateFile(s.config.report_file)

	s.log.Info("Starting server on port " + s.config.port)
	return s.r.Run(s.config.port)
}

// Shutdown the server
func (s *Server) Shutdown() error {
	s.log.Info("Stopping the server")

	// TODO: Write server graceful shutdown method

	s.log.Info("Server gracefully stopped")
	return nil
}
