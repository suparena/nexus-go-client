# nexus-go-client

A Go client for sending events to the Nexus event pipeline.

## Features

- Lightweight and minimal dependencies
- Supports sending single or batched events
- Uses context-aware HTTP client

## Installation

```bash
go get github.com/suparena/nexus-go-client
```

## Usage

```go
package main

import (
    "context"
    "log"
    "github.com/suparena/nexus-go-client"
)

func main() {
    client := nexus.NewClient("https://example.endpoint.dev/api", "your-token")

    err := client.Send(context.Background(), nexus.Event{
        Type: "match.started",
        Data: map[string]interface{}{
            "match_id": "abc123",
        },
    })

    if err != nil {
        log.Fatalf("failed to send event: %v", err)
    }
}
```

## Event Structure

```go
type Event struct {
    Type string                 `json:"type"`
    Data map[string]interface{} `json:"data"`
}
```

## License

MIT
