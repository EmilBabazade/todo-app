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

	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))
	mux.HandleFunc("GET /{$}", app.home)
	mux.HandleFunc("GET /todo/{id}", app.getTodo)
	mux.HandleFunc("PATCH /todo/{id}", app.patchTodo)
	mux.HandleFunc("DELETE /todo/{id}", app.deleteTodo)
	mux.HandleFunc("GET /todo/create", app.getTodoCreate)
	mux.HandleFunc("POST /todo/create", app.postTodoCreate)
	mux.HandleFunc("/", app.notFound)

	app.logger.Info("Serving at", "addr", *addr)

	err := http.ListenAndServe(*addr, mux)
	if err != nil {
		app.logger.Error(err.Error())
		os.Exit(1)
	}
}
