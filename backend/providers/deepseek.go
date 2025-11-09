package providers

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type DeepSeekProvider struct{}

func (d *DeepSeekProvider) GetAllModels(apiKey string) ([]string, error) {
	// Implement the logic to fetch all models from DeepSeek's API
	return []string{"deepseek-model-1", "deepseek-model-2"}, nil
}

func (d *DeepSeekProvider) SendMessage(prompt string, model string, apiKey string) (*http.Request, error) {
	payload := map[string]interface{}{
		"model": model,
		"messages": []map[string]interface{}{
			{"role": "user", "content": prompt},
			{"role": "system", "content": "You are a helpful assistant."},
		},
	}

	data, errorHandler := json.Marshal(payload)
	if errorHandler != nil {
		return nil, errorHandler
	}

	request, errorHandler := http.NewRequest("POST", "https://api.deepseek.com/chat/completions", bytes.NewBuffer(data))
	if errorHandler != nil {
		return nil, errorHandler
	}

	request.Header.Set("Authorization", "Bearer "+apiKey)
	request.Header.Set("Content-Type", "application/json")

	return request, nil
}
