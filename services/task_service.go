package services

import (
	"taskflow/models"
	"taskflow/repository"
)

// TaskService handles business logic for tasks.
// It depends on the TaskRepositoryInterface so it can be tested with a mock repo.
type TaskService struct {
	repo  repository.TaskRepositoryInterface
	cache *TaskCache
}

// NewTaskService creates a new TaskService with the given repository.
func NewTaskService(repo repository.TaskRepositoryInterface) *TaskService {
	return &TaskService{repo: repo, cache: NewTaskCache()}
}

// CreateTask saves a new task linked to the given user.
func (s *TaskService) CreateTask(task *models.Task) error {
	err := s.repo.Create(task)
	if err == nil {
		s.cache.Set(*task)
	}
	return err
}

// GetTasks returns all tasks belonging to the authenticated user.
func (s *TaskService) GetTasks(userID uint) ([]models.Task, error) {
	tasks, err := s.repo.GetAll(userID)
	if err != nil {
		return nil, err
	}
	for _, task := range tasks {
		s.cache.Set(task)
	}
	return tasks, nil
}

// GetTaskByID returns a single task by its ID.
func (s *TaskService) GetTaskByID(id uint) (*models.Task, error) {
	// --- CACHE HIT ---
	if cached, ok := s.cache.Get(id); ok {
		// RLock was used inside Get(); other goroutines can read simultaneously.
		return &cached, nil
	}
	// --- CACHE MISS ---
	task, err := s.repo.GetById(id)
	if err != nil {
		return nil, err
	}
	// Store in cache so the next request doesn't hit the DB.
	s.cache.Set(*task)
	return task, nil
}

// UpdateTask persists changes to an existing task.
func (s *TaskService) UpdateTask(task *models.Task) error {
	err := s.repo.Update(task)
	if err == nil {
		s.cache.Set(*task)
	}
	return err
}

// DeleteTask removes a task by its ID.
func (s *TaskService) DeleteTask(id uint) error {
	err := s.repo.Delete(id)
	if err == nil {
		s.cache.Delete(id)
	}
	return err
}
