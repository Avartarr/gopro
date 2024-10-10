package routes

import (
	"gopro/controllers"
	"gopro/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	// router.GET("/users", controllers.GetUsers)
	// router.POST("/users", controllers.CreateUser)
	// router.GET("/users/:id", controllers.GetUserByID)

	router.POST("/login", controllers.Login)
	router.POST("/users", controllers.CreateUser) // Registration (unprotected)

	// Protected routes
	protected := router.Group("/")
	protected.Use(middleware.JWTAuthMiddleware()) // Apply JWT middleware here
	protected.GET("/users", controllers.GetUsers) // Protected route
}
