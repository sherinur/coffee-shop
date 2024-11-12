package main

import (
	"flag"
	"fmt"
	"os"

	"hot-coffee/internal/server"

	. "hot-coffee/internal/utils"
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

	flag.Usage = CustomUsage
}

func validate() error {
	err := ValidatePort(port)
	if err != nil {
		return err
	}

	err = ValidateDir(dir)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	flag.Parse()

	validate()

	port = ":" + port

	cfg := server.NewConfig(configPath, port, dir)

	apiServer := server.New(cfg)
	err := apiServer.Start()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
