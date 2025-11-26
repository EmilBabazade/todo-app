package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

type application struct {
	logger *slog.Logger
}

func main() {
	addr := flag.String("addr", "127.0.0.1:5000", "HTTP network address")
	flag.Parse()

	logHandler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug,
	})
	logger := slog.New(logHandler)

	app := &application{
		logger: logger,
	}

	app.logger.Info("Serving at", "addr", *addr)
	err := http.ListenAndServe(*addr, app.routes())
	if err != nil {
		app.logger.Error(err.Error())
		os.Exit(1)
	}
}
