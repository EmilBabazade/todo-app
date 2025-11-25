package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /todo", getTodos)
	mux.HandleFunc("GET /todo/{id}", getTodo)
	mux.HandleFunc("PATCH /todo/{id}", patchTodo)
	mux.HandleFunc("DELETE /todo/{id}", deleteTodo)
	mux.HandleFunc("GET /todo/create", getTodoCreate)
	mux.HandleFunc("POST /todo/create", postTodoCreate)
	mux.HandleFunc("/", notFound)

	url := "127.0.0.1:5000"
	fmt.Printf("Serving at %v", url)

	err := http.ListenAndServe(url, mux)
	if err != nil {
		log.Fatal(err.Error())
	}
}
