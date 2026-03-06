package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"taskflow/dto"
	"taskflow/services"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt"
)

// AuthHandler handles registration, login, and token refresh.
type AuthHandler struct {
	authService *services.AuthService
	Validate    *validator.Validate
}

// NewAuthHandler creates a new AuthHandler with the required AuthService.
func NewAuthHandler(service *services.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: service,
		Validate:    validator.New(),
	}
}

// Register handles POST /auth/register — validates input and creates a new user account.
func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var input dto.RegisterDTO

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.Validate.Struct(input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.authService.Register(input); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User registered successfully"})
}

// Login handles POST /auth/login — validates credentials and returns access + refresh tokens.
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var input dto.LoginDTO

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.Validate.Struct(input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	access, refresh, err := h.authService.Login(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"access_token":  access,
		"refresh_token": refresh,
	})
}

// RefreshToken handles POST /auth/refresh — validates a refresh token and issues a new access token.
// The RefreshTokenDTO was defined in dto/auth_dto.go but had no corresponding handler until now.
func (h *AuthHandler) RefreshToken(w http.ResponseWriter, r *http.Request) {
	var input dto.RefreshTokenDTO

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.Validate.Struct(input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Parse and validate the refresh token
	token, err := jwt.Parse(input.RefreshToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil || !token.Valid {
		http.Error(w, "Invalid or expired refresh token", http.StatusUnauthorized)
		return
	}

	claims := token.Claims.(jwt.MapClaims)
	userIDFloat, ok := claims["user_id"].(float64)
	if !ok {
		http.Error(w, "Invalid token claims", http.StatusUnauthorized)
		return
	}
	userID := uint(userIDFloat)

	// Issue a new access + refresh token pair
	access, refresh, err := h.authService.GenerateTokensForUser(userID)
	if err != nil {
		http.Error(w, "Failed to generate tokens", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"access_token":  access,
		"refresh_token": refresh,
	})
}
