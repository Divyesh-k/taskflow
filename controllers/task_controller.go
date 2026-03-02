package controllers

import (
	"encoding/json"
	"net/http"
	"taskflow/config"
	"taskflow/models"

	"github.com/gorilla/mux"
)

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task

	json.NewDecoder(r.Body).Decode(&task)
	config.DB.Create(&task)
	json.NewEncoder(w).Encode(task)
}

func GetTasks(w http.ResponseWriter, r *http.Request) {
	var tasks []models.Task
	config.DB.Find(&tasks)
	json.NewEncoder(w).Encode(tasks)
}

func GetTask(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var task models.Task
	config.DB.First(&task, id)
	json.NewEncoder(w).Encode(task)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var task models.Task
	config.DB.Delete(&task, id)
	json.NewEncoder(w).Encode("Task deleted")
}
