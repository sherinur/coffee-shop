package main

import (
	"flag"
	"fmt"
	"os"

	"hot-coffee/internal/server"
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

func CustomUsage() {
	fmt.Println(`Coffee Shop Management System

Usage:
  hot-coffee [--port <N>] [--dir <S>] [--cfg <S>]
  hot-coffee --help

Options:
  --help       Show this screen.
  --port N     Port number.
  --dir S      Path to the data directory.
- --cfg S    Path to the config file`)
}

func init() {
	flag.StringVar(&port, "port", "4400", "Port number")
	flag.StringVar(&dir, "dir", "./data", "Path to the directory")
	flag.StringVar(&configPath, "cfg", "configs/server.yaml", "Path to the config file")

	flag.Usage = CustomUsage
}

func main() {
	flag.Parse()

	port = ":" + port

	cfg := server.NewConfig(configPath, port, dir)

	apiServer := server.New(cfg)
	err := apiServer.Start()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
