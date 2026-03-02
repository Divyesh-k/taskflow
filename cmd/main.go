package cmd

import (
	"log"
	"net/http"
	"taskflow/config"
	"taskflow/models"
	"taskflow/routes"
)

func main() {
	config.ConnectDB()

	config.DB.AutoMigrate(&models.Task{})

	r := routes.SetupRoutes()

	log.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", r)
}
