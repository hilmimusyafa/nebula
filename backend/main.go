package main

import (
	"nebula/backend/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Nebula Backend API is running!",
			"status":  "ok",
		})
	})

	router.GET("/chat")
	router.GET("/models")
	router.GET("/settings")
	router.GET("/about")

	router.POST("/send-message", handlers.SendPromptHandler)
	router.POST("/get-models", handlers.GetModelsHandler)

	router.Run(":8080")
}
