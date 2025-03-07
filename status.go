package obsidian

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// StatusService handles server status operations
type StatusService struct {
	client *Client
}

// GetStatus returns basic details about the server and authentication status.
// This is the only API request that does not require authentication.
func (s *StatusService) GetStatus(ctx context.Context) (*ServerInfo, error) {
	req, err := http.NewRequest("GET", s.client.baseURL.String(), nil)
	if err != nil {
		return nil, err
	}

	// Don't use client.do here as this endpoint doesn't require authentication
	if ctx == nil {
		ctx = context.Background()
	}
	req = req.WithContext(ctx)

	resp, err := s.client.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var info ServerInfo
	if err := json.NewDecoder(resp.Body).Decode(&info); err != nil {
		return nil, err
	}

	return &info, nil
}
