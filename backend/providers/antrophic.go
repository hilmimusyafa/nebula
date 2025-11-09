package providers

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type AnthropicProvider struct{}

func (a *AnthropicProvider) GetAllModels(apiKey string) ([]string, error) {
	// Anthropic doesn't have a public models endpoint, return known models
	models := []string{
		"claude-3-opus-20240229",
		"claude-3-sonnet-20240229",
		"claude-3-haiku-20240307",
		"claude-2.1",
		"claude-2.0",
		"claude-instant-1.2",
	}
	return models, nil
}

func (a *AnthropicProvider) SendMessage(prompt string, model string, apiKey string) (*http.Request, error) {
	payload := map[string]interface{}{
		"model":      model,
		"max_tokens": 1000,
		"messages": []map[string]string{
			{
				"role":    "user",
				"content": prompt,
			},
		},
	}

	data, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", "https://api.anthropic.com/v1/messages", bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	req.Header.Set("x-api-key", apiKey)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("anthropic-version", "2023-06-01")

	return req, nil
}
