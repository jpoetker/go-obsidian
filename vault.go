package obsidian

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"path"
	"strings"
)

// VaultService handles operations on the vault
type VaultService struct {
	client *Client
}

// GetFile retrieves the content of a file from the vault
func (s *VaultService) GetFile(ctx context.Context, filepath string) (*NoteJson, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/vault/%s", s.client.baseURL.String(), path.Clean(filepath)), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/vnd.olrapi.note+json")

	resp, err := s.client.do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return nil, fmt.Errorf("file not found: %s", filepath)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var note NoteJson
	if err := json.NewDecoder(resp.Body).Decode(&note); err != nil {
		return nil, err
	}

	return &note, nil
}

// ListDirectory lists the contents of a directory in the vault
func (s *VaultService) ListDirectory(ctx context.Context, dirpath string) (*DirectoryListing, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/vault/%s", s.client.baseURL.String(), path.Clean(dirpath)), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return nil, fmt.Errorf("directory not found: %s", dirpath)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var listing DirectoryListing
	if err := json.NewDecoder(resp.Body).Decode(&listing); err != nil {
		return nil, err
	}

	return &listing, nil
}

// CreateOrUpdateFile creates a new file or updates an existing one in the vault
func (s *VaultService) CreateOrUpdateFile(ctx context.Context, filepath string, content string) error {
	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/vault/%s", s.client.baseURL.String(), path.Clean(filepath)), nil)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "text/markdown")
	req.Body = io.NopCloser(strings.NewReader(content))

	resp, err := s.client.do(ctx, req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return nil
}

// DeleteFile deletes a file from the vault
func (s *VaultService) DeleteFile(ctx context.Context, filepath string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/vault/%s", s.client.baseURL.String(), path.Clean(filepath)), nil)
	if err != nil {
		return err
	}

	resp, err := s.client.do(ctx, req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return nil
}

// AppendToFile appends content to an existing file
func (s *VaultService) AppendToFile(ctx context.Context, filepath string, content string) error {
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/vault/%s", s.client.baseURL.String(), path.Clean(filepath)), nil)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "text/markdown")
	req.Body = io.NopCloser(strings.NewReader(content))

	resp, err := s.client.do(ctx, req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return nil
}
