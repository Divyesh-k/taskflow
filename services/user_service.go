package services

import (
	"taskflow/models"
	"taskflow/repository"
)

// UserService handles business logic for user profile management.
type UserService struct {
	repo *repository.UserRepository
}

// NewUserService creates a new UserService with the given repository.
func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

// GetUser returns a single user by their ID (e.g. for profile view).
func (s *UserService) GetUser(id uint) (*models.User, error) {
	return s.repo.FindByID(id)
}

// GetAllUsers returns all users with their tasks (admin only).
func (s *UserService) GetAllUsers() ([]models.User, error) {
	return s.repo.GetAllUsers()
}

// UpdateUser updates the name and/or email of an existing user.
func (s *UserService) UpdateUser(id uint, name string, email string) (*models.User, error) {
	user, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	// Only update fields that were provided (non-empty)
	if name != "" {
		user.Name = name
	}
	if email != "" {
		user.Email = email
	}

	if err := s.repo.UpdateUser(user); err != nil {
		return nil, err
	}
	return user, nil
}

// DeleteUser removes a user account by their ID.
func (s *UserService) DeleteUser(id uint) error {
	return s.repo.DeleteUser(id)
}
