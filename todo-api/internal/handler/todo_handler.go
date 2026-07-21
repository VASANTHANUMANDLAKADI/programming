package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/todo-api/internal/model"
	"github.com/todo-api/internal/storage"
)

type TodoHandler struct {
	storage *storage.TodoStorage
}

func NewTodoHandler(storage *storage.TodoStorage) *TodoHandler {
	return &TodoHandler{
		storage: storage,
	}
}

func (h *TodoHandler) CreateTodo(w http.ResponseWriter, r *http.Request) {
	var req model.CreateTodoRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if strings.TrimSpace(req.Title) == "" {
		http.Error(w, "title is required", http.StatusBadRequest)
		return
	}

	todo := model.Todo{
		Title:       req.Title,
		Description: req.Description,
	}

	createdTodo := h.storage.AddTodo(todo)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(createdTodo)
}