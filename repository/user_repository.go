package repository

import (
	"taskflow/models"

	"gorm.io/gorm"
)

// UserRepository handles all database operations for the User model.
type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

// CreateUser inserts a new user record into the database.
func (r *UserRepository) CreateUser(user *models.User) error {
	return r.DB.Create(user).Error
}

// FindByEmail retrieves a user by their email address (used for login).
func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	if err := r.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// FindByID retrieves a single user by their primary key (used for profile, etc.).
func (r *UserRepository) FindByID(id uint) (*models.User, error) {
	var user models.User
	if err := r.DB.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// GetAllUsers returns all users with their tasks preloaded (admin only).
func (r *UserRepository) GetAllUsers() ([]models.User, error) {
	var users []models.User
	if err := r.DB.Preload("Tasks").Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// UpdateUser saves changes to an existing user record.
func (r *UserRepository) UpdateUser(user *models.User) error {
	return r.DB.Save(user).Error
}

// DeleteUser soft-deletes a user by their ID (GORM uses soft delete via DeletedAt).
func (r *UserRepository) DeleteUser(id uint) error {
	return r.DB.Delete(&models.User{}, id).Error
}
