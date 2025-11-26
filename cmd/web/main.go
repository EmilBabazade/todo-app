package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	addr := flag.String("addr", "127.0.0.1", "HTTP network address")

	logHandler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug,
	})
	logger := slog.New(logHandler)

	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))
	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /todo/{id}", getTodo)
	mux.HandleFunc("PATCH /todo/{id}", patchTodo)
	mux.HandleFunc("DELETE /todo/{id}", deleteTodo)
	mux.HandleFunc("GET /todo/create", getTodoCreate)
	mux.HandleFunc("POST /todo/create", postTodoCreate)
	mux.HandleFunc("/", notFound)

	logger.Info("Serving at", "addr", *addr)

	err := http.ListenAndServe(*addr, mux)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}
