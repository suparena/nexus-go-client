package main

import (
	"context"
	"log"
)

func main() {
	client := nexus.NewClient("https://example.endpoint.dev/api", "your-token")

	event := nexus.Event{
		Type: "example.event",
		Data: map[string]interface{}{
			"user_id": "12345",
			"action":  "test_run",
		},
	}

	if err := client.Send(context.Background(), event); err != nil {
		log.Fatalf("Failed to send event: %v", err)
	}

	log.Println("Event sent successfully.")
}
