package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")
	servePage(w, r, "./ui/html/pages/home.tmpl.html")
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
	servePage(w, r, "./ui/html/pages/notFound.tmpl.html")
}

func servePage(w http.ResponseWriter, r *http.Request, file string) {
	files := []string{
		"./ui/html/base.tmpl.html",
		"./ui/html/partials/nav.tmpl.html",
		file,
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		internalServerError(w, err)
		return
	}

	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		internalServerError(w, err)
		return
	}
}

func internalServerError(w http.ResponseWriter, err error) {
	log.Print(err.Error())
	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
}
