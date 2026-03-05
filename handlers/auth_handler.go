package handlers

import (
	"encoding/json"
	"net/http"
	"taskflow/dto"
	"taskflow/services"

	"github.com/go-playground/validator/v10"
)

type AuthHandler struct {
	authService *services.AuthService
	Validate    *validator.Validate
}

func NewAuthHandler(service *services.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: service,
		Validate:    validator.New(),
	}
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var input dto.RegisterDTO

	json.NewDecoder(r.Body).Decode(&input)

	if err := h.Validate.Struct(input); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	err := h.authService.Register(input)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var input dto.LoginDTO

	json.NewDecoder(r.Body).Decode(&input)

	access, refresh, err := h.authService.Login(input)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"access_token":  access,
		"refresh_token": refresh,
	})
}
