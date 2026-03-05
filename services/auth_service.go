package services

import (
	"errors"
	"taskflow/dto"
	"taskflow/models"
	"taskflow/repository"
	"taskflow/utils"
)

type AuthService struct {
	userRepo *repository.UserRepository
}

func NewAuthService(repo *repository.UserRepository) *AuthService {
	return &AuthService{repo}
}

func (s *AuthService) Register(input dto.RegisterDTO) error {
	hash, err := utils.HashPassword(input.Password)
	if err != nil {
		return err
	}

	user := models.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: hash,
		Role:     "user",
	}

	return s.userRepo.CreateUser(&user)
}

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
