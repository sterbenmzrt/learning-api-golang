package main

import (
	"learning-api-golang/handler"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/todos", handler.HandleTodos)
	http.HandleFunc("/todos/", handler.HandleTodoByID)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("🚀 ToDo API is running!"))
	})

	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
