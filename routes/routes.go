package routes

import (
	"taskflow/controllers"

	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/tasks", controllers.CreateTask).Methods("POST")
	router.HandleFunc("/tasks", controllers.GetTasks).Methods("GET")
	router.HandleFunc("/tasks/{id}", controllers.GetTask).Methods("GET")
	router.HandleFunc("/tasks/{id}", controllers.DeleteTask).Methods("DELETE")

	return router
}
