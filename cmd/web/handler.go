package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")
	app.page(w, "./ui/html/pages/home.tmpl.html", nil)
}

// view a todo
func (app *application) getTodo(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 0)
	if err != nil || id < 1 {
		app.logger.Error("invalid pathvalue", "error", err)
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Todo with id %d", id)
}

// edit a todo patch
func (app *application) patchTodo(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 0)
	if err != nil || id < 1 {
		app.logger.Error("invalid pathvalue", "error", err)
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Editing todo with id %d", id)
}

// delete a todo
func (app *application) deleteTodo(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 0)
	if err != nil || id < 1 {
		app.logger.Error("invalid pathvalue", "error", err)
		http.NotFound(w, r)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// create a todo page
func (app *application) getTodoCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create a todo"))
}

// create a todo
func (app *application) postTodoCreate(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Creating a todo"))
}

func (app *application) notFound(w http.ResponseWriter, r *http.Request) {
	app.page(w, "./ui/html/pages/notFound.tmpl.html", nil)
}

func (app *application) page(w http.ResponseWriter, file string, data any) {
	files := []string{
		"./ui/html/base.tmpl.html",
		"./ui/html/partials/nav.tmpl.html",
		file,
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.logger.Error("Invalid files", "error", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = ts.ExecuteTemplate(w, "base", data)
	if err != nil {
		app.logger.Error("Invalid templates", "error", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
