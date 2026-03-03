package controllers

import (
	"encoding/json"
	"net/http"
	"taskflow/config"
	"taskflow/models"
	"taskflow/utils"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var user models.User

	json.NewDecoder(r.Body).Decode(&user)

	hashPassword, err := utils.HashPassword(user.Password)

	if err != nil {
		http.Error(w, "Error hashing Password", http.StatusInternalServerError)
	}

	user.Password = hashPassword

	config.DB.Create(&user)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)

}

func Login(w http.ResponseWriter, r *http.Request) {
	var input models.User
	var user models.User

	json.NewDecoder(r.Body).Decode(&input)

	config.DB.Where("email = ?", input.Email).First(&user)

	if user.ID == 0 {
		http.Error(w, "User not found", http.StatusUnauthorized)
		return
	}

	if !utils.CheckPassword(input.Password, user.Password) {
		http.Error(w, "Invalid Password", http.StatusUnauthorized)
		return
	}

	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
