package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"taskflow/dto"
	"taskflow/models"
	"taskflow/services"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

// TaskHandler handles HTTP requests for task operations.
type TaskHandler struct {
	service  *services.TaskService
	validate *validator.Validate
}

// NewTaskHandler creates a new TaskHandler with the required TaskService.
func NewTaskHandler(service *services.TaskService) *TaskHandler {
	return &TaskHandler{
		service:  service,
		validate: validator.New(),
	}
}

// CreateTask handles POST /tasks — validates input and creates a task for the authenticated user.
func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(uint)

	var input dto.TaskCreateDTO
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate input against struct tags (required, min, oneof, etc.)
	if err := h.validate.Struct(input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	task := models.Task{
		Title:       input.Title,
		Description: input.Description,
		Status:      input.Status,
		UserId:      userID,
	}

	if err := h.service.CreateTask(&task); err != nil {
		http.Error(w, "Failed to create task", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

// GetTasks handles GET /tasks — returns all tasks for the authenticated user only.
func (h *TaskHandler) GetTasks(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(uint)

	tasks, err := h.service.GetTasks(userID)
	if err != nil {
		http.Error(w, "Failed to retrieve tasks", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tasks)
}

// UpdateTask handles PUT /tasks/{id} — reads ID from URL, updates task for the authenticated user.
// BUG FIX: Previously the ID came only from the body (often zero), causing the wrong row to be updated.
func (h *TaskHandler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(uint)

	// Read task ID from URL parameter
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	var input dto.TaskUpdateDTO
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.validate.Struct(input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Fetch the existing task to make sure it belongs to this user
	task, err := h.service.GetTaskByID(uint(id))
	if err != nil {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}
	if task.UserId != userID {
		http.Error(w, "Forbidden: you don't own this task", http.StatusForbidden)
		return
	}

	// Apply partial updates (only overwrite fields that were provided)
	if input.Title != "" {
		task.Title = input.Title
	}
	if input.Description != "" {
		task.Description = input.Description
	}
	if input.Status != "" {
		task.Status = input.Status
	}

	if err := h.service.UpdateTask(task); err != nil {
		http.Error(w, "Failed to update task", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(task)
}

// DeleteTask handles DELETE /tasks/{id} — removes a task by ID.
func (h *TaskHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	if err := h.service.DeleteTask(uint(id)); err != nil {
		http.Error(w, "Failed to delete task", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Task deleted successfully"})
}
