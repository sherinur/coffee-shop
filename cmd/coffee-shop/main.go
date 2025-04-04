package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"coffee-shop/internal/app"
	"coffee-shop/internal/transport/http/server"
	"coffee-shop/internal/utils"
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

	ctx := context.Background()

	cfg := server.NewConfig(configPath, ":"+port, dir)

	application, err := app.New(ctx, cfg)
	if err != nil {
		fmt.Println("failed to setup an application:", err)
		os.Exit(1)
	}

	err = application.Run()
	if err != nil {
		fmt.Println("error of application:", err)
		os.Exit(1)
	}
}
