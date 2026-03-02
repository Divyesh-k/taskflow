package controllers

import (
	"encoding/json"
	"net/http"
	"taskflow/config"
	"taskflow/models"

	"github.com/gorilla/mux"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	id := mux.Vars(r)["id"]
	config.DB.Preload("Tasks").First(&user, id)
	json.NewEncoder(w).Encode(user)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	json.NewDecoder(r.Body).Decode(&user)
	config.DB.Create(&user)
	json.NewEncoder(w).Encode(user)
}
