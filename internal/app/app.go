package app

import (
	"coffee-shop/internal/repository/postgres"
	"coffee-shop/internal/transport/http/server"
	"coffee-shop/pkg/logger"
	"context"
	"database/sql"
)

const serviceName = "coffee-shop"

type App struct {
	httpServer *server.Server
}

func New(ctx context.Context, cfg server.Config) (App, error) {
	log := logger.SetupLogger(&logger.LoggerOptions{Env: cfg.Env, LogFilepath: cfg.Log_file})
	log.Info("Logger is initialized successfully")
	return App{server.New(&cfg, log)}, nil
}

func (a *App) Run() error {
	connStr := "user=latte password=latte dbname=frappuccino sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}

	postgres.NewInventory(db)
	postgres.NewMenu(db)
	postgres.NewOrder(db)

	a.httpServer.Start()
	if err != nil {
		return err
	}

	return nil
}
