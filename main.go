package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Wekcome to todo app"))
}

// todos list
func getTodos(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("List of todos..."))
}

// view a todo
func getTodo(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 0)
	if err != nil || id < 1 {
		log.Print(err)
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Todo with id %d", id)
}

// edit a todo patch
func patchTodo(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 0)
	if err != nil || id < 1 {
		log.Print(err)
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Editing todo with id %d", id)
}

// delete a todo
func deleteTodo(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 0)
	if err != nil || id < 1 {
		log.Print(err)
		http.NotFound(w, r)
		return
	}
	w.WriteHeader(http.StatusNoContent)
	fmt.Fprintf(w, "Deleting todo with id %d", id)
}

// create a todo page
func getTodoCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create a todo"))
}

// create a todo
func postTodoCreate(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Creating a todo"))
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("404 that don't exist boy"))
}

func main() {
	mux := http.NewServeMux()
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
