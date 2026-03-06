package services

import (
	"errors"
	"taskflow/dto"
	"taskflow/models"
	"taskflow/repository"
	"taskflow/utils"
)

// AuthService handles user registration and login logic.
type AuthService struct {
	userRepo *repository.UserRepository
}

// NewAuthService creates a new AuthService with the given UserRepository.
func NewAuthService(repo *repository.UserRepository) *AuthService {
	return &AuthService{repo}
}

// Register hashes the password and creates a new user in the database.
func (s *AuthService) Register(input dto.RegisterDTO) error {
	hash, err := utils.HashPassword(input.Password)
	if err != nil {
		return err
	}

	user := models.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: hash,
		Role:     input.Role, // role comes from validated DTO ("admin" or "user")
	}

	return s.userRepo.CreateUser(&user)
}

// Login verifies credentials and returns an access token + refresh token pair on success.
func (s *AuthService) Login(input dto.LoginDTO) (string, string, error) {
	user, err := s.userRepo.FindByEmail(input.Email)
	if err != nil {
		return "", "", errors.New("invalid credentials")
	}

	if !utils.CheckPassword(input.Password, user.Password) {
		return "", "", errors.New("invalid credentials")
	}

	return utils.GenerateTokens(user.ID, user.Role)
}

// GenerateTokensForUser issues a fresh access+refresh token pair for a user by ID.
// Used by the RefreshToken handler to issue new tokens without re-authenticating.
func (s *AuthService) GenerateTokensForUser(userID uint) (string, string, error) {
	user, err := s.userRepo.FindByID(userID)
	if err != nil {
		return "", "", errors.New("user not found")
	}
	return utils.GenerateTokens(user.ID, user.Role)
}
