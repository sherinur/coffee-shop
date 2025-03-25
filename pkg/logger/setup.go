package logger

import (
	"fmt"
	"io"
	"log/slog"
	"os"
)

const (
	EnvLocal = "local"
	EnvDev   = "dev"
	EnvProd  = "prod"
)

type LoggerOptions struct {
	Env         string
	LogFilepath string
}

func SetupLogger(opts *LoggerOptions) *slog.Logger {
	writer := getWriter(opts)
	handler := getHandler(opts, *writer)
	return slog.New(handler)
}

func getHandler(opts *LoggerOptions, writer io.Writer) slog.Handler {
	level := getLogLevel(opts.Env)
	handlerOpts := &slog.HandlerOptions{Level: level}

	switch opts.Env {
	case EnvLocal:
		return slog.NewTextHandler(writer, handlerOpts)
	default:
		return slog.NewJSONHandler(writer, handlerOpts)
	}
}

func getWriter(opts *LoggerOptions) *io.Writer {
	var writer io.Writer
	if opts.LogFilepath != "" {
		file := mustOpen(opts.LogFilepath)
		writer = io.MultiWriter(os.Stdout, file)
	} else {
		writer = io.MultiWriter(os.Stdout)
	}
	return &writer
}

func getLogLevel(env string) slog.Level {
	switch env {
	case EnvProd:
		return slog.LevelInfo
	default:
		return slog.LevelDebug
	}
}

func mustOpen(filepath string) *os.File {
	err := os.Mkdir("logs", 0777)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	file, err := os.OpenFile(filepath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o666)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return file
}
