package handlers

import (
	"nebula/backend/models"
	"nebula/backend/providers"
	"nebula/backend/services"

	"github.com/gin-gonic/gin"
)

var alreadyproviders = map[string]models.Provider{
	"openai":    &providers.OpenAIProvider{},
	"anthropic": &providers.AnthropicProvider{},
	// "google":    &providers.GoogleProvider{},
	// "xai":       &providers.xAIProvider{},
	"deepseek": &providers.DeepSeekProvider{},
}

func SendPromptHandler(c *gin.Context) {
	var request models.SendMessage

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	response, errorHandler := services.SendMessage(request)
	if errorHandler != nil {
		c.JSON(500, gin.H{"error": errorHandler.Error()})
		return
	}

	c.JSON(200, gin.H{"response": response})
}

func GetModelsHandler(c *gin.Context) {
	var request models.GetModels

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}
}
