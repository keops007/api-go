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
	r.GET("/shopping", func(c *gin.Context) { c.File("./static/shopping.html") })

	// API public
	r.POST("/register", handlers.Register)
	r.POST("/login", handlers.Login)

	// API protejat cu JWT
	api := r.Group("/api")
	api.Use(middleware.AuthMiddleware())
	{
		api.GET("/profile", handlers.GetProfile)
		api.POST("/profile/avatar", handlers.UploadAvatar)
		api.GET("/shopping", handlers.GetShoppingItems)
		api.POST("/shopping", handlers.AddShoppingItem)
		api.PATCH("/shopping/:id/done", handlers.ToggleShoppingItem)
		api.DELETE("/shopping/:id", handlers.DeleteShoppingItem)
	}
}
