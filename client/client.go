package client

import (
	"errors"
	"net/url"

	"github.com/jpoetker/go-obsidian/security"
)

func NewAuthenticedClientWithResponses(server string, apikey string, opts ...ClientOption) (*ClientWithResponses, error) {
	client, err := NewAuthenticatedClient(server, apikey, opts...)
	if err != nil {
		return nil, err
	}
	return &ClientWithResponses{client}, nil
}

func NewAuthenticatedClient(server string, apikey string, opts ...ClientOption) (*Client, error) {
	if apikey == "" {
		return nil, errors.New("apikey is required")
	}

	// Validate server URL
	parsedURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}
	if parsedURL.Scheme == "" || parsedURL.Host == "" {
		return nil, errors.New("invalid server URL: missing scheme or host")
	}

	securityProvider := security.NewAPIKeyProvider(apikey)

	// Append the security provider to the existing options
	opts = append(opts, WithRequestEditorFn(securityProvider.Intercept))

	client, err := NewClient(server, opts...)
	if err != nil {
		return nil, err
	}
	return client, nil
}
