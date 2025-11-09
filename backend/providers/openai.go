package providers

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type OpenAIProvider struct{}

func (o *OpenAIProvider) GetAllModels(apiKey string) ([]string, error) {
	var models []string

	request, errorHandler := http.NewRequest("GET", "https://api.openai.com/v1/models", nil)
	if errorHandler != nil {
		return nil, errorHandler
	}

	request.Header.Set("Authorization", "Bearer "+apiKey)
	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, errorHandler := client.Do(request)
	if errorHandler != nil {
		return nil, errorHandler
	}

	defer response.Body.Close()

	body, errerrorHandler := io.ReadAll(response.Body)
	if errerrorHandler != nil {
		return nil, errerrorHandler
	}

	var result map[string]interface{}
	errorHandler = json.Unmarshal(body, &result)
	if errorHandler != nil {
		return nil, errorHandler
	}

	if data, ok := result["data"].([]interface{}); ok {
		for _, item := range data {
			if model, ok := item.(map[string]interface{})["id"].(string); ok {
				models = append(models, model)
			}
		}
	}

	return models, nil
}

func (o *OpenAIProvider) SendMessage(prompt string, model string, apiKey string) (*http.Request, error) {
	payload := map[string]interface{}{
		"model": model,
		"messages": []map[string]string{
			{
				"role":    "user",
				"content": prompt,
			},
		},
	}

	data, errorHandler := json.Marshal(payload)
	if errorHandler != nil {
		return nil, errorHandler
	}

	request, errorHandler := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(data))
	if errorHandler != nil {
		return nil, errorHandler
	}

	request.Header.Set("Authorization", "Bearer "+apiKey)
	request.Header.Set("Content-Type", "application/json")

	return request, nil
}
