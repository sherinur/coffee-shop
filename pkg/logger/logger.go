package logger

import (
	"log"
	"log/slog"
	"net/http"
)

// TODO: Rewrite log to slog.
// TODO: Check what type of logs we can print in this project and оставить разрешенные

// ? TODO: Save logs to the ./logs/triple-s.log path (OPTIONAL)
type iLogger interface {
	PrintfInfoMsg(string, ...interface{})
	PrintfDebugMsg(string, ...interface{})
	PrintfErrorMsg(string, ...interface{})
	PrintfWarnMsg(string, ...interface{})
	LogRequestMiddleware(http.Handler) http.Handler
}

type Logger struct {
	debugMode    bool
	bracketsMode bool
}

func NewLogger(debugMode bool, bracketsMode bool) *Logger {
	return &Logger{
		debugMode:    debugMode,
		bracketsMode: bracketsMode,
	}
}

func printfMsg(level string, mes string, args ...interface{}) {
	log.Printf(level+" "+mes, args...)
}

func (l *Logger) PrintInfoMsg(mes string, args ...interface{}) {
	if l.bracketsMode {
		printfMsg("[INFO]", mes, args...)
		return
	}
	slog.Info(mes, args...)
}

func (l *Logger) PrintDebugMsg(mes string, args ...interface{}) {
	if l.debugMode {
		printfMsg("[DEBUG]", mes, args...)
	}
}

func (l *Logger) PrintErrorMsg(mes string, args ...interface{}) {
	if l.bracketsMode {
		printfMsg("[ERROR]", mes, args...)
		return
	}
	slog.Error(mes, args...)
}

func (l *Logger) PrintWarnMsg(mes string, args ...interface{}) {
	if l.bracketsMode {
		printfMsg("[WARN]", mes, args...)
		return
	}
	slog.Warn(mes, args...)
}

func (l *Logger) LogRequestMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[INFO] Request %s %s from %s", r.Method, r.URL.Path, r.RemoteAddr)
		next.ServeHTTP(w, r)
	})
}
