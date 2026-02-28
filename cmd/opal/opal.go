package main

import (
	"flag"
	"log/slog"
	"net/http"
	"opal/internal/config"
	"opal/internal/librarymgmt"
	"opal/internal/usermgmt"
	"os"
	"time"

	"github.com/lmittmann/tint"
	// _ "net/http/pprof"
)

// TODO: use linker version numbers for releases
var version_number = "dev"

func main() {
	logLevel := flag.String("log", "info", "set log level (debug, info, warn, error)")
	flag.Parse()

	var level slog.Level
	switch *logLevel {
	case "debug":
		level = slog.LevelDebug
	case "warn":
		level = slog.LevelWarn
	case "error":
		level = slog.LevelError
	default:
		level = slog.LevelInfo
	}

	logger := slog.New(tint.NewHandler(os.Stdout, &tint.Options{
		Level:      level,
		TimeFormat: time.TimeOnly,
	}))

	slog.SetDefault(logger)

	slog.Info("Opal media server starting", "version", version_number)

	if *logLevel == "debug" {
		addr := "localhost:6060"
		go func() {
			slog.Debug("starting pprof debugger", "addr", addr)
			if err := http.ListenAndServe(addr, nil); err != nil && err != http.ErrServerClosed {
				slog.Error("pprof debugger failed", "error", err)
			}
		}()
	}

	//TODO: make configurable
	dbDir := "./db"
	if err := os.MkdirAll(dbDir, 0755); err != nil {
		slog.Error("failed to mkdir", "path", dbDir)
		os.Exit(1)
	}

	config.Init()
	usermgmt.Init()
	librarymgmt.Init()

	//httpserver.Start()
}
