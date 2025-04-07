package main

import (
	"log"
	"net/http"
	"todo-api/handler"
)

func main() {
	http.HandleFunc("/todos", handler.HandleTodos)
	http.HandleFunc("/todos/", handler.HandleTodoByID)

	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
