package security

import (
	"context"
	"net/http"
)

// APIKeyProvider handles API key authentication for the Obsidian Local REST API
type APIKeyProvider struct {
	apiKey string
}

// NewAPIKeyProvider creates a new API key security provider
func NewAPIKeyProvider(apiKey string) *APIKeyProvider {
	return &APIKeyProvider{
		apiKey: apiKey,
	}
}

// Intercept adds the API key to the request header
func (p *APIKeyProvider) Intercept(ctx context.Context, req *http.Request) error {
	req.Header.Set("Authorization", "Bearer "+p.apiKey)
	return nil
}
