package main

import "net/http"

func (app *application) routes() *http.ServeMux {
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

	return mux
}
