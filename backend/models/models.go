package models

import (
	"net/http"
)

type Provider interface {
	GetAllModels(apiKey string) ([]string, error)
	SendMessage(prompt string, model string, apiKey string) (*http.Request, error)
}

type AIProvider struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	URL     string `json:"url"`
	APIKey  string `json:"api_key"`
}

type GetModels struct {
	APIKey   string     `json:"api_key"`
	Provider AIProvider `json:"provider"`
}

type SendMessage struct {
	APIKey   string     `json:"api_key"`
	Prompt   string     `json:"prompt"`
	Model    string     `json:"model"`
	Provider AIProvider `json:"provider"`
}
