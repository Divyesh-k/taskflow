package services

import (
	"taskflow/models"
	"taskflow/repository"
)

// TaskService handles business logic for tasks.
// It depends on the TaskRepositoryInterface so it can be tested with a mock repo.
type TaskService struct {
	repo repository.TaskRepositoryInterface
}

// NewTaskService creates a new TaskService with the given repository.
func NewTaskService(repo repository.TaskRepositoryInterface) *TaskService {
	return &TaskService{repo: repo}
}

// CreateTask saves a new task linked to the given user.
func (s *TaskService) CreateTask(task *models.Task) error {
	return s.repo.Create(task)
}

// GetTasks returns all tasks belonging to the authenticated user.
func (s *TaskService) GetTasks(userID uint) ([]models.Task, error) {
	return s.repo.GetAll(userID)
}

// GetTaskByID returns a single task by its ID.
func (s *TaskService) GetTaskByID(id uint) (*models.Task, error) {
	return s.repo.GetById(id)
}

// UpdateTask persists changes to an existing task.
func (s *TaskService) UpdateTask(task *models.Task) error {
	return s.repo.Update(task)
}

// DeleteTask removes a task by its ID.
func (s *TaskService) DeleteTask(id uint) error {
	return s.repo.Delete(id)
}
