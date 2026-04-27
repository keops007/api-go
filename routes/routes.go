package routes

import (
	"github.com/gin-gonic/gin"
	"go-api/handlers"
	"go-api/middleware"
)

func SetupRoutes(r *gin.Engine) {
	// Fisiere statice
	r.Static("/uploads", "./uploads")

	// Pagini HTML
	r.GET("/register", func(c *gin.Context) { c.File("./static/register.html") })
	r.GET("/login", func(c *gin.Context) { c.File("./static/login.html") })
	r.GET("/profile", func(c *gin.Context) { c.File("./static/profile.html") })

	// API public
	r.POST("/register", handlers.Register)
	r.POST("/login", handlers.Login)

	// API protejat cu JWT
	api := r.Group("/api")
	api.Use(middleware.AuthMiddleware())
	{
		api.GET("/profile", handlers.GetProfile)
		api.POST("/profile/avatar", handlers.UploadAvatar)
	}
}
