package app

import (
	"coffee-shop/internal/repository/postgres"
	"coffee-shop/internal/transport/http/server"
	"coffee-shop/pkg/logger"
	"context"
	"database/sql"
	"fmt"
	"log/slog"
)

const serviceName = "coffee-shop"
const connStr = "user=latte password=latte dbname=frappuccino sslmode=disable"

type App struct {
	httpServer *server.Server
	log        *slog.Logger
}

func New(ctx context.Context, cfg *server.Config) (*App, error) {
	log := logger.SetupLogger(&logger.LoggerOptions{Env: cfg.Env, LogFilepath: cfg.Log_file})
	log.Info("logger is initialized successfully")

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	// TODO: Repository
	postgres.NewInventory(db)
	postgres.NewMenu(db)
	postgres.NewOrder(db)

	// TODO: UseCase

	// TODO: http service

	return &App{
		httpServer: server.New(cfg, log),
		log:        log,
	}, nil
}

func (a *App) Close() {
	err := a.httpServer.Shutdown()
	if err != nil {
		a.log.Error("failed to shutdown the service:", err)
	}
}

func (a *App) Run() error {
	// TODO: Implement the gracefull shutdown

	a.log.Info(fmt.Sprintf("starting the %v service", serviceName))
	err := a.httpServer.Start()
	if err != nil {
		return err
	}

	return nil
}
