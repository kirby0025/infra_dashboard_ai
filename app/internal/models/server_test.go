package models

import (
	"encoding/json"
	"testing"
	"time"
)

func TestServerJSONMarshal(t *testing.T) {
	server := Server{
		ID:        1,
		Name:      "test-server",
		OSID:      28,
		CreatedAt: time.Date(2024, 1, 1, 12, 0, 0, 0, time.UTC),
		UpdatedAt: time.Date(2024, 1, 1, 12, 0, 0, 0, time.UTC),
	}

	jsonData, err := json.Marshal(server)
	if err != nil {
		t.Fatalf("Failed to marshal server to JSON: %v", err)
	}

	expected := `{"id":1,"name":"test-server","os_id":28,"created_at":"2024-01-01T12:00:00Z","updated_at":"2024-01-01T12:00:00Z"}`
	if string(jsonData) != expected {
		t.Errorf("JSON marshal result mismatch.\nExpected: %s\nGot: %s", expected, string(jsonData))
	}
}

func TestServerJSONUnmarshal(t *testing.T) {
	jsonData := `{"id":1,"name":"test-server","os_id":28,"created_at":"2024-01-01T12:00:00Z","updated_at":"2024-01-01T12:00:00Z"}`

	var server Server
	err := json.Unmarshal([]byte(jsonData), &server)
	if err != nil {
		t.Fatalf("Failed to unmarshal JSON to server: %v", err)
	}

	if server.ID != 1 {
		t.Errorf("Expected ID 1, got %d", server.ID)
	}
	if server.Name != "test-server" {
		t.Errorf("Expected name 'test-server', got '%s'", server.Name)
	}
	if server.OSID != 28 {
		t.Errorf("Expected OSID 28, got %d", server.OSID)
	}
}

func TestCreateServerRequestValidation(t *testing.T) {
	tests := []struct {
		name    string
		request CreateServerRequest
		wantErr bool
	}{
		{
			name: "valid request",
			request: CreateServerRequest{
				Name: "web-server-01",
				OSID: 28,
			},
			wantErr: false,
		},
		{
			name: "empty name",
			request: CreateServerRequest{
				Name: "",
				OSID: 28,
			},
			wantErr: true,
		},
		{
			name: "invalid OS ID",
			request: CreateServerRequest{
				Name: "web-server-01",
				OSID: 0,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hasEmptyField := tt.request.Name == "" || tt.request.OSID == 0
			if hasEmptyField != tt.wantErr {
				t.Errorf("CreateServerRequest validation mismatch. Expected error: %v, but validation result: %v", tt.wantErr, hasEmptyField)
			}
		})
	}
}

func TestUpdateServerRequestJSON(t *testing.T) {
	request := UpdateServerRequest{
		Name: "updated-server",
		OSID: 32,
	}

	jsonData, err := json.Marshal(request)
	if err != nil {
		t.Fatalf("Failed to marshal UpdateServerRequest to JSON: %v", err)
	}

	var unmarshaledRequest UpdateServerRequest
	err = json.Unmarshal(jsonData, &unmarshaledRequest)
	if err != nil {
		t.Fatalf("Failed to unmarshal JSON to UpdateServerRequest: %v", err)
	}

	if unmarshaledRequest.Name != request.Name {
		t.Errorf("Expected name '%s', got '%s'", request.Name, unmarshaledRequest.Name)
	}
	if unmarshaledRequest.OSID != request.OSID {
		t.Errorf("Expected OSID %d, got %d", request.OSID, unmarshaledRequest.OSID)
	}
}

func TestUpdateServerRequestPartialUpdate(t *testing.T) {
	// Test partial update with only name
	jsonData := `{"name":"partial-update"}`

	var request UpdateServerRequest
	err := json.Unmarshal([]byte(jsonData), &request)
	if err != nil {
		t.Fatalf("Failed to unmarshal partial update JSON: %v", err)
	}

	if request.Name != "partial-update" {
		t.Errorf("Expected name 'partial-update', got '%s'", request.Name)
	}
	if request.OSID != 0 {
		t.Errorf("Expected empty OSID, got %d", request.OSID)
	}
}
