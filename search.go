package obsidian

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// SearchService handles search operations
type SearchService struct {
	client *Client
}

// SimpleSearchOptions represents options for simple search
type SimpleSearchOptions struct {
	ContextLength int
}

// SimpleSearch performs a simple text search across the vault
func (s *SearchService) SimpleSearch(ctx context.Context, query string, opts *SimpleSearchOptions) ([]SimpleSearchResult, error) {
	u, err := url.Parse(fmt.Sprintf("%s/search/simple", s.client.baseURL.String()))
	if err != nil {
		return nil, err
	}

	q := u.Query()
	q.Set("query", query)
	if opts != nil && opts.ContextLength > 0 {
		q.Set("contextLength", fmt.Sprintf("%d", opts.ContextLength))
	}
	u.RawQuery = q.Encode()

	req, err := http.NewRequest("POST", u.String(), nil)
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

	var results []SimpleSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&results); err != nil {
		return nil, err
	}

	return results, nil
}

// SearchQuery represents a search query for the advanced search endpoint
type SearchQuery struct {
	Query     interface{} // Can be string for DQL or map[string]interface{} for JsonLogic
	QueryType string      // "dql" or "jsonlogic"
}

// Search performs an advanced search using either DQL or JsonLogic
func (s *SearchService) Search(ctx context.Context, query SearchQuery) ([]SearchResult, error) {
	var contentType string
	var body io.Reader

	switch query.QueryType {
	case "dql":
		contentType = "application/vnd.olrapi.dataview.dql+txt"
		body = strings.NewReader(query.Query.(string))
	case "jsonlogic":
		contentType = "application/vnd.olrapi.jsonlogic+json"
		jsonData, err := json.Marshal(query.Query)
		if err != nil {
			return nil, err
		}
		body = strings.NewReader(string(jsonData))
	default:
		return nil, fmt.Errorf("invalid query type: %s", query.QueryType)
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/search", s.client.baseURL.String()), body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", contentType)

	resp, err := s.client.do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var results []SearchResult
	if err := json.NewDecoder(resp.Body).Decode(&results); err != nil {
		return nil, err
	}

	return results, nil
}
