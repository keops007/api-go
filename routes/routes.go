package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"go-api/handlers"
	"go-api/middleware"
	"go-api/repository"
	"go-api/services"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	// Repositories — stiu sa vorbeasca cu baza de date
	userRepo := repository.NewUserRepository(db)
	shoppingRepo := repository.NewShoppingRepository(db)

	// Services — contin logica de business, nu stiu nimic despre HTTP
	authSvc := services.NewAuthService(userRepo)
	profileSvc := services.NewProfileService(userRepo)
	shoppingSvc := services.NewShoppingService(shoppingRepo)

	// Handlers — primesc request HTTP, apeleaza services, returneaza JSON
	authHandler := handlers.NewAuthHandler(authSvc)
	profileHandler := handlers.NewProfileHandler(profileSvc)
	shoppingHandler := handlers.NewShoppingHandler(shoppingSvc)

	// Fisiere statice
	r.Static("/uploads", "./uploads")
	r.StaticFile("/logo.png", "./logo.png")

	// Pagini HTML
	r.GET("/", func(c *gin.Context) { c.File("./static/index.html") })
	r.GET("/register", func(c *gin.Context) { c.File("./static/register.html") })
	r.GET("/login", func(c *gin.Context) { c.File("./static/login.html") })
	r.GET("/profile", func(c *gin.Context) { c.File("./static/profile.html") })
	r.GET("/shopping", func(c *gin.Context) { c.File("./static/shopping.html") })

	// API public
	r.POST("/register", authHandler.Register)
	r.POST("/login", authHandler.Login)

	// API protejat cu JWT
	api := r.Group("/api")
	api.Use(middleware.AuthMiddleware())
	{
		api.GET("/profile", profileHandler.GetProfile)
		api.POST("/profile/avatar", profileHandler.UploadAvatar)
		api.GET("/shopping", shoppingHandler.GetItems)
		api.POST("/shopping", shoppingHandler.AddItem)
		api.PATCH("/shopping/:id/done", shoppingHandler.ToggleItem)
		api.DELETE("/shopping/:id", shoppingHandler.DeleteItem)
	}
}
