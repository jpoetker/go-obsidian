// Package obsidian provides a client for interacting with the Obsidian Local REST API
package obsidian

import (
	"context"
	"net/http"
	"net/url"
)

// Client represents an Obsidian REST API client
type Client struct {
	baseURL    *url.URL
	httpClient *http.Client
	apiKey     string

	// Services
	Status   *StatusService
	Vault    *VaultService
	Search   *SearchService
	Commands *CommandService
}

// ClientOption is a function that modifies the client configuration
type ClientOption func(*Client)

// WithHTTPClient sets a custom HTTP client
func WithHTTPClient(httpClient *http.Client) ClientOption {
	return func(c *Client) {
		c.httpClient = httpClient
	}
}

// WithBaseURL sets a custom base URL for the API
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) {
		if parsedURL, err := url.Parse(baseURL); err == nil {
			c.baseURL = parsedURL
		}
	}
}

// NewClient creates a new Obsidian API client
func NewClient(apiKey string, options ...ClientOption) *Client {
	defaultURL, _ := url.Parse("https://127.0.0.1:27124")

	client := &Client{
		baseURL:    defaultURL,
		httpClient: http.DefaultClient,
		apiKey:     apiKey,
	}

	// Apply any custom options
	for _, option := range options {
		option(client)
	}

	// Initialize services
	client.Status = &StatusService{client: client}
	client.Vault = &VaultService{client: client}
	client.Search = &SearchService{client: client}
	client.Commands = &CommandService{client: client}

	return client
}

// do performs an HTTP request and returns the response
func (c *Client) do(ctx context.Context, req *http.Request) (*http.Response, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	req = req.WithContext(ctx)
	req.Header.Set("Authorization", "Bearer "+c.apiKey)

	return c.httpClient.Do(req)
}
