package handler

import (
	"encoding/json"
	"learning-api-golang/model"
	"learning-api-golang/storage"
	"learning-api-golang/utils"
	"net/http"
	"strconv"
	"strings"
)

func HandleTodos(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		todos := storage.GetAllTodos()
		utils.WriteJSON(w, http.StatusOK, todos)

	case http.MethodPost:
		var todo model.Todo
		err := json.NewDecoder(r.Body).Decode(&todo)
		if err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}
		created := storage.AddTodo(todo)
		utils.WriteJSON(w, http.StatusCreated, created)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func HandleTodoByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/todos/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		todo, err := storage.GetTodoByID(id)
		if err != nil {
			http.Error(w, "Not found", http.StatusNotFound)
			return
		}
		utils.WriteJSON(w, http.StatusOK, todo)

	case http.MethodPut:
		var todo model.Todo
		err := json.NewDecoder(r.Body).Decode(&todo)
		if err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}
		if err := storage.UpdateTodo(id, todo); err != nil {
			http.Error(w, "Not found", http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusNoContent)

	case http.MethodDelete:
		if err := storage.DeleteTodo(id); err != nil {
			http.Error(w, "Not found", http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusNoContent)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
