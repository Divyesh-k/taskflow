package main

import "fmt"

// import (
// 	"log"
// 	"net/http"
// 	"os"
// 	"taskflow/config"
// 	"taskflow/models"
// 	"taskflow/routes"

// 	"github.com/joho/godotenv"
// )

// func main() {

// 	// ✅ start env globally
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatal("No .env file found")
// 	}

// 	config.ConnectDB()

// 	config.DB.AutoMigrate(&models.Task{}, &models.Task{})

// 	router := routes.SetupRoutes()

// 	port := os.Getenv("API_PORT")

// 	log.Println("Server running on port", port)

// 	http.ListenAndServe(":"+port, router)
// }

func main() {
	var x interface{}

	value, ok := x.(int)
	if !ok {
		fmt.Println("Not an integer")
		return
	}
	fmt.Println(value)
}
