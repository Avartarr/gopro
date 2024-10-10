package main

import (
	"gopro/models"
	"gopro/routes"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	// "os"
)

func main() {
	// r := gin.Default()
	// r.GET("/", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "Hello, Gin!",
	// 	})
	// })
	// r.Run()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading environment")
	}

	models.ConnectDatabase()

	r := gin.Default()

	routes.SetupRoutes(r)

	r.Run(":8080")
}
