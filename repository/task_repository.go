package repository

import (
	"taskflow/config"
	"taskflow/models"
)

// TaskRepositoryInterface defines the contract for task data operations.
// Using an interface allows the service layer to be tested independently (mock-friendly).
type TaskRepositoryInterface interface {
	Create(task *models.Task) error
	GetAll(userId uint) ([]models.Task, error)
	GetById(id uint) (*models.Task, error)
	Update(task *models.Task) error
	Delete(id uint) error
}

// TaskRepository is the concrete implementation backed by GORM/PostgreSQL.
type TaskRepository struct{}

func (r *TaskRepository) Create(task *models.Task) error {
	return config.DB.Create(task).Error
}

func (r *TaskRepository) GetAll(userId uint) ([]models.Task, error) {
	var tasks []models.Task
	err := config.DB.Where("user_id = ?", userId).Find(&tasks).Error
	return tasks, err
}

func (r *TaskRepository) GetById(id uint) (*models.Task, error) {
	var task models.Task
	err := config.DB.First(&task, id).Error
	return &task, err
}

func (r *TaskRepository) Update(task *models.Task) error {
	return config.DB.Save(task).Error
}

func (r *TaskRepository) Delete(id uint) error {
	return config.DB.Delete(&models.Task{}, id).Error
}
