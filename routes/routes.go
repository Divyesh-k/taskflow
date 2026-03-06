package routes

import (
	"net/http"
	"taskflow/config"
	"taskflow/handlers"
	"taskflow/middleware"
	"taskflow/repository"
	"taskflow/services"

	"github.com/gorilla/mux"
)

// SetupRoutes builds and returns the router with all routes wired to their handlers.
// Architecture: Handler → Service → Repository (layered, testable)
func SetupRoutes() *mux.Router {
	router := mux.NewRouter()

	// ─── Dependency Injection ───────────────────────────────────────────────────
	userRepo := repository.NewUserRepository(config.DB)
	taskRepo := &repository.TaskRepository{}

	authService := services.NewAuthService(userRepo)
	taskService := services.NewTaskService(taskRepo)
	userService := services.NewUserService(userRepo)

	authHandler := handlers.NewAuthHandler(authService)
	taskHandler := handlers.NewTaskHandler(taskService)
	userHandler := handlers.NewUserHandler(userService)

	// ─── Auth Routes (public) ───────────────────────────────────────────────────
	router.HandleFunc("/auth/register", authHandler.Register).Methods(http.MethodPost)
	router.HandleFunc("/auth/login", authHandler.Login).Methods(http.MethodPost)
	// RefreshTokenDTO existed but had no route — now wired up
	router.HandleFunc("/auth/refresh", authHandler.RefreshToken).Methods(http.MethodPost)

	// ─── Task Routes (JWT protected) ────────────────────────────────────────────
	router.Handle("/tasks", middleware.JWTAuth(http.HandlerFunc(taskHandler.CreateTask))).Methods(http.MethodPost)
	router.Handle("/tasks", middleware.JWTAuth(http.HandlerFunc(taskHandler.GetTasks))).Methods(http.MethodGet)
	// PUT /tasks/{id} was completely missing — now added
	router.Handle("/tasks/{id}", middleware.JWTAuth(http.HandlerFunc(taskHandler.UpdateTask))).Methods(http.MethodPut)
	router.Handle("/tasks/{id}", middleware.JWTAuth(http.HandlerFunc(taskHandler.DeleteTask))).Methods(http.MethodDelete)

	// ─── User Profile Routes (JWT protected) ────────────────────────────────────
	router.Handle("/users/me", middleware.JWTAuth(http.HandlerFunc(userHandler.GetProfile))).Methods(http.MethodGet)
	router.Handle("/users/me", middleware.JWTAuth(http.HandlerFunc(userHandler.UpdateProfile))).Methods(http.MethodPut)
	router.Handle("/users/me", middleware.JWTAuth(http.HandlerFunc(userHandler.DeleteAccount))).Methods(http.MethodDelete)

	// ─── Admin Routes (JWT + role check) ────────────────────────────────────────
	// BUG FIX: Previously /admin used RequireRole without JWTAuth first, so
	// "role" was never set in context → always panicked or returned unauthorized.
	// Now JWTAuth runs first, then RequireRole checks the context value.
	adminMiddleware := middleware.JWTAuth(middleware.RequireRole("admin")(http.HandlerFunc(userHandler.AdminGetAllUsers)))
	router.Handle("/admin/users", adminMiddleware).Methods(http.MethodGet)

	return router
}
