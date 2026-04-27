package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"go-api/config"
	"go-api/models"
	"go-api/routes"
)

func main() {
	godotenv.Load()
	os.MkdirAll("./uploads", 0755)

	config.ConnectDB()
	config.DB.AutoMigrate(&models.User{}, &models.ShoppingItem{})

	r := gin.Default()
	routes.SetupRoutes(r, config.DB)
	r.Run(":8081")
}
