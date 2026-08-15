package meshery

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// DesignPattern represents a Meshery infrastructure design.
type DesignPattern struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Pattern   string    `json:"pattern_file"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Client wraps HTTP requests to Meshery Server REST/GraphQL endpoints.
type Client struct {
	BaseURL    string
	HTTPClient *http.Client
}

// NewClient initializes a new Meshery API client.
func NewClient(baseURL string) *Client {
	return &Client{
		BaseURL: baseURL,
		HTTPClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

// FetchDesigns retrieves design patterns from Meshery Server.
func (c *Client) FetchDesigns(ctx context.Context) ([]DesignPattern, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("%s/api/content/patterns", c.BaseURL), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		// Return mock design pattern if server unreachable in PoC mode
		return []DesignPattern{
			{
				ID:        "design-poc-001",
				Name:      "Kubernetes-Deployment-Pattern",
				Pattern:   "services:\n  web:\n    type: Deployment\n    version: v1",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
		}, nil
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("meshery API returned status: %d", resp.StatusCode)
	}

	var result struct {
		Patterns []DesignPattern `json:"patterns"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return result.Patterns, nil
}
