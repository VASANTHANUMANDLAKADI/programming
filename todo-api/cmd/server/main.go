package main

import (
	"log"
	"net/http"

	"github.com/todo-api/internal/handler"
	"github.com/todo-api/internal/storage"
)

func main() {
	// Create storage
	todoStorage := storage.NewTodoStorage()

	// Create handler
	todoHandler := handler.NewTodoHandler(todoStorage)

	// Register routes
	http.HandleFunc("/info", handler.InfoHandler)
	http.HandleFunc("/todos", todoHandler.CreateTodo)

	port := ":8080"

	log.Println("Server running on port", port)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(err)
	}
}