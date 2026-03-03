package routes

import (
	"net/http"
	"taskflow/controllers"
	"taskflow/middleware"

	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	router := mux.NewRouter()

	router.Handle("/admin", middleware.RequireRole("admin")(http.HandlerFunc(controllers.AdminDashboard))).Methods("GET")
	router.Handle("/tasks", middleware.JWTAuth(http.HandlerFunc(controllers.CreateTask))).Methods("POST")
	router.Handle("/tasks", middleware.JWTAuth(http.HandlerFunc(controllers.GetTasks))).Methods("GET")
	router.Handle("/tasks/{id}", middleware.JWTAuth(http.HandlerFunc(controllers.GetTask))).Methods("GET")
	router.Handle("/tasks/{id}", middleware.JWTAuth(http.HandlerFunc(controllers.DeleteTask))).Methods("DELETE")

	router.HandleFunc("/users/{id}", controllers.GetUser).Methods("GET")
	router.HandleFunc("/users", controllers.CreateUser).Methods("POST")

	return router
}
