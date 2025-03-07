package obsidian

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// CommandService handles command operations
type CommandService struct {
	client *Client
}

// Command represents an Obsidian command
type Command struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// CommandList represents a list of available commands
type CommandList struct {
	Commands []Command `json:"commands"`
}

// List returns all available commands
func (s *CommandService) List(ctx context.Context) ([]Command, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/commands", s.client.baseURL.String()), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var list CommandList
	if err := json.NewDecoder(resp.Body).Decode(&list); err != nil {
		return nil, err
	}

	return list.Commands, nil
}

// Execute executes a command by its ID
func (s *CommandService) Execute(ctx context.Context, commandID string) error {
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/commands/%s", s.client.baseURL.String(), commandID), nil)
	if err != nil {
		return err
	}

	resp, err := s.client.do(ctx, req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return fmt.Errorf("command not found: %s", commandID)
	}

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return nil
}
