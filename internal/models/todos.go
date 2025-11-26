package models

import (
	"database/sql"
	"time"
)

type Todo struct {
	Id        int
	Title     string
	Content   string
	Created   time.Time
	Deadline  time.Time
	Completed bool
}

type TodoModel struct {
	DB *sql.DB
}

func (model *TodoModel) GetMany() ([]Todo, error) {
	return []Todo{}, nil
}

func (model *TodoModel) Get(id int) (Todo, error) {
	return Todo{}, nil
}

func (model *TodoModel) Insert(title string, content string, deadline time.Time) error {
	return nil
}

func (model *TodoModel) Complete(id int) error {
	return nil
}

func (model *TodoModel) Update(id int, title string, content string, deadline time.Time, completed bool) error {
	return nil
}

func (model *TodoModel) Delete(id int) error {
	return nil
}
