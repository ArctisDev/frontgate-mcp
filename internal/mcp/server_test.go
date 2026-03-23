package mcp

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"log"
	"strings"
	"testing"
)

func TestServerHandlesBackToBackRawJSONMessages(t *testing.T) {
	t.Parallel()

	input := strings.Join([]string{
		`{"jsonrpc":"2.0","id":0,"method":"initialize","params":{"protocolVersion":"2025-11-25","capabilities":{},"clientInfo":{"name":"opencode","version":"1.2.27"}}}`,
		`{"jsonrpc":"2.0","method":"notifications/initialized"}`,
		`{"jsonrpc":"2.0","id":1,"method":"tools/list"}`,
		"",
	}, "\n")

	var output bytes.Buffer
	server := NewServer(log.New(io.Discard, "", 0))

	err := server.Run(context.Background(), strings.NewReader(input), &output)
	if err != nil {
		t.Fatalf("Run() returned error: %v", err)
	}

	lines := strings.Split(strings.TrimSpace(output.String()), "\n")
	if len(lines) != 2 {
		t.Fatalf("expected 2 responses, got %d: %q", len(lines), output.String())
	}

	var initResp Response
	if err := json.Unmarshal([]byte(lines[0]), &initResp); err != nil {
		t.Fatalf("failed to parse initialize response: %v", err)
	}
	if initResp.ID != float64(0) {
		t.Fatalf("expected initialize response id 0, got %#v", initResp.ID)
	}

	var toolsResp struct {
		JSONRPC string `json:"jsonrpc"`
		ID      int    `json:"id"`
		Result  struct {
			Tools []map[string]any `json:"tools"`
		} `json:"result"`
	}
	if err := json.Unmarshal([]byte(lines[1]), &toolsResp); err != nil {
		t.Fatalf("failed to parse tools/list response: %v", err)
	}
	if toolsResp.ID != 1 {
		t.Fatalf("expected tools/list response id 1, got %d", toolsResp.ID)
	}
	if len(toolsResp.Result.Tools) == 0 {
		t.Fatalf("expected tools/list to return tools")
	}
}
