package main

import (
	"os"

	"go-api/config"
	"go-api/models"
	"go-api/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
    godotenv.Load()
    os.MkdirAll("./uploads", 0755)

    // Connect to database
    config.ConnectDB()

    // Auto migrate models
    config.DB.AutoMigrate(&models.User{}, &models.ShoppingItem{})

    r := gin.Default()

    // ❗ THIS IS CRITICAL (your problem most likely here)
    routes.SetupRoutes(r)

    r.Run(":8081")
}
