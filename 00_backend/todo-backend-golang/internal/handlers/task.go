// internal/handlers/task.go
package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"todo-backend-golang/internal/models"
	"todo-backend-golang/internal/storage"

	"github.com/gorilla/mux"
)

// TaskHandler maneja las operaciones relacionadas con tareas
type TaskHandler struct {
	storage *storage.MemoryStorage
}

// NewTaskHandler crea un nuevo controlador de tareas
func NewTaskHandler(storage *storage.MemoryStorage) *TaskHandler {
	return &TaskHandler{storage: storage}
}

// GetAllTasks maneja GET /tasks
func (h *TaskHandler) GetAllTasks(w http.ResponseWriter, r *http.Request) {
	tasks := h.storage.GetAll()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

// GetTaskByID maneja GET /tasks/{id}
func (h *TaskHandler) GetTaskByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	task, err := h.storage.GetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

// CreateTask maneja POST /tasks
func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	createdTask := h.storage.Create(&task)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdTask)
}

// UpdateTask maneja PUT /tasks/{id}
func (h *TaskHandler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var updatedTask models.Task
	if err := json.NewDecoder(r.Body).Decode(&updatedTask); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	task, err := h.storage.Update(id, &updatedTask)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

// DeleteTask maneja DELETE /tasks/{id}
func (h *TaskHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	if err := h.storage.Delete(id); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
