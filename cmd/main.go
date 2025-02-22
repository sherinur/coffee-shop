package main

import (
	"flag"
	"fmt"
	"os"

	"coffee-shop/internal/server"
	"coffee-shop/internal/utils"
	"coffee-shop/pkg/logger"
)

// TODO: Test how flag parsing works, and find bugs

// ! Fix the bug when there are no dir or port or cfg path

// TODO: Validate the port correctly for the RESTful API (or REST API). Write in the utils package validatePort.go

// TODO: Validate the path to the data directory correctly (Client dir creating fix). Write in the utils package validateDir.go

// ? TODO: Validate the path to the cfg correctly. Write in the utils package validateCfg.go (OPTIONAL)

var (
	configPath string
	port       string
	dir        string
)

func init() {
	flag.StringVar(&port, "port", "8080", "Port number")
	flag.StringVar(&dir, "dir", "./data", "Path to the directory")
	flag.StringVar(&configPath, "cfg", "configs/server.yaml", "Path to the config file")

	flag.Usage = utils.CustomUsage
}

func main() {
	flag.Parse()

	err := utils.ValidatePort(port)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = utils.ValidateDir(dir)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	cfg := server.NewConfig(configPath, ":"+port, dir)

	log := logger.SetupLogger(&logger.LoggerOptions{Env: cfg.Env, LogFilepath: cfg.Log_file})
	if log == nil {
		fmt.Println("Logger initialization failed (Logger instance is nil)")
		os.Exit(1)
	}
	log.Info("Logger is initialized successfully")

	apiServer := server.New(cfg, log)
	err = apiServer.Start()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
