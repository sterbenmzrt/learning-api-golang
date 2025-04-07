package storage

import (
	"errors"
	"todo-api/model"
)

var todos []model.Todo
var nextID = 1

func GetAllTodos() []model.Todo {
	return todos
}

func GetTodoByID(id int) (*model.Todo, error) {
	for _, t := range todos {
		if t.ID == id {
			return &t, nil
		}
	}
	return nil, errors.New("todo not found")
}

func AddTodo(todo model.Todo) model.Todo {
	todo.ID = nextID
	nextID++
	todos = append(todos, todo)
	return todo
}

func UpdateTodo(id int, updated model.Todo) error {
	for i, t := range todos {
		if t.ID == id {
			updated.ID = id
			todos[i] = updated
			return nil
		}
	}
	return errors.New("todo not found")
}

func DeleteTodo(id int) error {
	for i, t := range todos {
		if t.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			return nil
		}
	}
	return errors.New("todo not found")
}
