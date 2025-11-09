package services

import (
	"fmt"
	"io"
	"nebula/backend/models"
	"nebula/backend/providers"
	"net/http"
)

var alreadyproviders = map[string]models.Provider{
	"openai":    &providers.OpenAIProvider{},
	"anthropic": &providers.AnthropicProvider{},
	// "google":    &providers.GoogleProvider{},
	// "xai":       &providers.xAIProvider{},
	"deepseek": &providers.DeepSeekProvider{},
}

func SendMessage(request models.SendMessage) (string, error) {
	provider, ok := alreadyproviders[request.Provider.Name]
	if !ok {
		return "", fmt.Errorf("provider not found: %s", request.Provider.Name)
	}

	// Get HTTP request from provider
	httpRequest, err := provider.SendMessage(request.Prompt, request.Model, request.APIKey)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %v", err)
	}

	// Execute the HTTP request
	client := &http.Client{}
	response, err := client.Do(httpRequest)
	if err != nil {
		return "", fmt.Errorf("failed to send request: %v", err)
	}
	defer response.Body.Close()

	// Read response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response: %v", err)
	}

	return string(body), nil
}
