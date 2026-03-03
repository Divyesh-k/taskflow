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

func Login(w http.ResponseWriter , r *http.Request) {
	
}
