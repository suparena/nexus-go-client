package nexus

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// Event represents a single event to be sent to the Nexus pipeline.
type Event struct {
	Type      string                 `json:"type"`
	Timestamp string                 `json:"timestamp,omitempty"`
	Data      map[string]interface{} `json:"data"`
}

// Client handles communication with the Nexus event endpoint.
type Client struct {
	url    string
	token  string
	client *http.Client
}

// NewClient returns a new Nexus client instance.
func NewClient(url, token string) *Client {
	return &Client{
		url:    url,
		token:  token,
		client: http.DefaultClient,
	}
}

// Send sends a single event to the Nexus event endpoint.
func (c *Client) Send(ctx context.Context, event Event) error {
	events := []Event{event}
	return c.SendBatch(ctx, events)
}

// SendBatch sends multiple events in a single request.
func (c *Client) SendBatch(ctx context.Context, events []Event) error {
	payload, err := json.Marshal(events)
	if err != nil {
		return fmt.Errorf("failed to marshal events: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", c.url, bytes.NewBuffer(payload))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.token)

	resp, err := c.client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("received non-2xx status: %s", resp.Status)
	}

	return nil
}
