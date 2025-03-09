package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestNewAuthenticatedClient(t *testing.T) {
	tests := []struct {
		name        string
		server      string
		apikey      string
		wantErr     bool
		errContains string
	}{
		{
			name:    "valid initialization",
			server:  "http://localhost:27124",
			apikey:  "test-api-key",
			wantErr: false,
		},
		{
			name:        "empty api key",
			server:      "http://localhost:27124",
			apikey:      "",
			wantErr:     true,
			errContains: "apikey is required",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, err := NewAuthenticatedClient(tt.server, tt.apikey)

			if tt.wantErr {
				if err == nil {
					t.Errorf("NewAuthenticatedClient() error = nil, wantErr %v", tt.wantErr)
					return
				}
				if tt.errContains != "" && !strings.Contains(err.Error(), tt.errContains) {
					t.Errorf("NewAuthenticatedClient() error = %v, want error containing %v", err, tt.errContains)
				}
				return
			}

			if err != nil {
				t.Errorf("NewAuthenticatedClient() unexpected error = %v", err)
				return
			}

			if client == nil {
				t.Error("NewAuthenticatedClient() returned nil client")
			}
		})
	}
}

func TestNewAuthenticedClientWithResponses(t *testing.T) {
	tests := []struct {
		name        string
		server      string
		apikey      string
		wantErr     bool
		errContains string
	}{
		{
			name:    "valid initialization",
			server:  "http://localhost:27124",
			apikey:  "test-api-key",
			wantErr: false,
		},
		{
			name:        "empty api key",
			server:      "http://localhost:27124",
			apikey:      "",
			wantErr:     true,
			errContains: "apikey is required",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, err := NewAuthenticedClientWithResponses(tt.server, tt.apikey)

			if tt.wantErr {
				if err == nil {
					t.Errorf("NewAuthenticedClientWithResponses() error = nil, wantErr %v", tt.wantErr)
					return
				}
				if tt.errContains != "" && !strings.Contains(err.Error(), tt.errContains) {
					t.Errorf("NewAuthenticedClientWithResponses() error = %v, want error containing %v", err, tt.errContains)
				}
				return
			}

			if err != nil {
				t.Errorf("NewAuthenticedClientWithResponses() unexpected error = %v", err)
				return
			}

			if client == nil {
				t.Error("NewAuthenticedClientWithResponses() returned nil client")
			}
		})
	}
}

func TestClientWithServer(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify the API key is set in the Authorization header
		authHeader := r.Header.Get("Authorization")
		expectedHeader := "Bearer test-api-key"
		if authHeader != expectedHeader {
			t.Errorf("Expected Authorization header '%s', got '%s'", expectedHeader, authHeader)
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	client, err := NewAuthenticatedClient(server.URL, "test-api-key")
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	req, err := http.NewRequest("GET", server.URL, nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	// Add the security provider's request editor
	ctx := req.Context()
	for _, editor := range client.RequestEditors {
		err := editor(ctx, req)
		if err != nil {
			t.Fatalf("Failed to apply request editor: %v", err)
		}
	}

	resp, err := client.Client.Do(req)
	if err != nil {
		t.Fatalf("Failed to do request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status OK, got %v", resp.Status)
	}
}

// Helper function to check if a string contains another string
func contains(s, substr string) bool {
	return s != "" && substr != "" && s != substr && len(s) > len(substr) && s[len(s)-len(substr):] == substr
}
