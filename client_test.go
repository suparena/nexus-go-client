package nexus

import (
	"context"
	"os"
	"testing"
)

func TestSend(t *testing.T) {
	endpoint := os.Getenv("NEXUS_ENDPOINT")
	token := os.Getenv("NEXUS_TOKEN")

	if endpoint == "" || token == "" {
		t.Skip("NEXUS_ENDPOINT or NEXUS_TOKEN not set; skipping integration test")
	}

	client := NewClient(endpoint, token)

	err := client.Send(context.Background(), Event{
		Type: "test.event",
		Data: map[string]interface{}{
			"key": "value",
		},
	})

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}
