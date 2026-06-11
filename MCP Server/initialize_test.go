package main

import (
	"encoding/json"
	"testing"
)

func TestParseInput_Initialize(t *testing.T) {
	params := InitializeParams{
		ProtocolVersion: "2024-11-05",
		Capabilities:    Capabilities{},
		ClientInfo: ClientInfo{
			Name:    "test-client",
			Version: "1.0.0",
		},
	}

	rawParams, _ := json.Marshal(params)

	req := MCPRequest{
		JSONRPC: "2.0",
		ID:      1,
		Method:  "initialize",
		Params:  json.RawMessage(rawParams),
	}

	body, _ := json.Marshal(req)

	parsed := parseInput(body)

	if parsed.Method != "initialize" {
		t.Errorf("expected method 'initialize', got %q", parsed.Method)
	}
	if parsed.ID != 1 {
		t.Errorf("expected ID 1, got %d", parsed.ID)
	}
	if parsed.JSONRPC != "2.0" {
		t.Errorf("expected jsonrpc '2.0', got %q", parsed.JSONRPC)
	}

	var initParams InitializeParams
	if err := json.Unmarshal(parsed.Params, &initParams); err != nil {
		t.Fatalf("failed to unmarshal params: %v", err)
	}
	if initParams.ProtocolVersion != "2024-11-05" {
		t.Errorf("expected protocolVersion '2024-11-05', got %q", initParams.ProtocolVersion)
	}
	if initParams.ClientInfo.Name != "test-client" {
		t.Errorf("expected client name 'test-client', got %q", initParams.ClientInfo.Name)
	}
}

func TestParseContentLength(t *testing.T) {
	header := "Content-Length: 42\r\n"
	got := parseContentLength(header)
	if got != 42 {
		t.Errorf("expected 42, got %d", got)
	}
}
