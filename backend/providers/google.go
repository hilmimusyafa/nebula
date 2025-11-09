package providers

type GoogleProvider struct{}

func (g *GoogleProvider) SendMessage(prompt string, model string, apiKey string) (string, error) {
	// Implement the logic to send a message to Google's API
	return "", nil
}