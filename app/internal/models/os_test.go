package models

import (
	"encoding/json"
	"testing"
	"time"
)

func TestOSJSONMarshal(t *testing.T) {
	os := OS{
		ID:           1,
		Name:         "Ubuntu",
		Version:      "22.04",
		EndOfSupport: time.Date(2027, 4, 1, 0, 0, 0, 0, time.UTC),
		CreatedAt:    time.Date(2024, 1, 1, 12, 0, 0, 0, time.UTC),
		UpdatedAt:    time.Date(2024, 1, 1, 12, 0, 0, 0, time.UTC),
	}

	jsonData, err := json.Marshal(os)
	if err != nil {
		t.Fatalf("Failed to marshal OS to JSON: %v", err)
	}

	expected := `{"id":1,"name":"Ubuntu","version":"22.04","end_of_support":"2027-04-01T00:00:00Z","created_at":"2024-01-01T12:00:00Z","updated_at":"2024-01-01T12:00:00Z"}`
	if string(jsonData) != expected {
		t.Errorf("JSON marshal result mismatch.\nExpected: %s\nGot: %s", expected, string(jsonData))
	}
}

func TestOSJSONUnmarshal(t *testing.T) {
	jsonData := `{"id":1,"name":"Ubuntu","version":"22.04","end_of_support":"2027-04-01T00:00:00Z","created_at":"2024-01-01T12:00:00Z","updated_at":"2024-01-01T12:00:00Z"}`

	var os OS
	err := json.Unmarshal([]byte(jsonData), &os)
	if err != nil {
		t.Fatalf("Failed to unmarshal JSON to OS: %v", err)
	}

	if os.ID != 1 {
		t.Errorf("Expected ID 1, got %d", os.ID)
	}
	if os.Name != "Ubuntu" {
		t.Errorf("Expected name 'Ubuntu', got '%s'", os.Name)
	}
	if os.Version != "22.04" {
		t.Errorf("Expected version '22.04', got '%s'", os.Version)
	}
	expectedDate := time.Date(2027, 4, 1, 0, 0, 0, 0, time.UTC)
	if !os.EndOfSupport.Equal(expectedDate) {
		t.Errorf("Expected end of support date '%v', got '%v'", expectedDate, os.EndOfSupport)
	}
}

func TestCreateOSRequestValidation(t *testing.T) {
	tests := []struct {
		name    string
		request CreateOSRequest
		wantErr bool
	}{
		{
			name: "valid request",
			request: CreateOSRequest{
				Name:         "Debian",
				Version:      "12",
				EndOfSupport: "2028-06-30",
			},
			wantErr: false,
		},
		{
			name: "empty name",
			request: CreateOSRequest{
				Name:         "",
				Version:      "12",
				EndOfSupport: "2028-06-30",
			},
			wantErr: true,
		},
		{
			name: "empty version",
			request: CreateOSRequest{
				Name:         "Debian",
				Version:      "",
				EndOfSupport: "2028-06-30",
			},
			wantErr: true,
		},
		{
			name: "empty end of support",
			request: CreateOSRequest{
				Name:         "Debian",
				Version:      "12",
				EndOfSupport: "",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hasEmptyField := tt.request.Name == "" || tt.request.Version == "" || tt.request.EndOfSupport == ""
			if hasEmptyField != tt.wantErr {
				t.Errorf("CreateOSRequest validation mismatch. Expected error: %v, but validation result: %v", tt.wantErr, hasEmptyField)
			}
		})
	}
}

func TestUpdateOSRequestJSON(t *testing.T) {
	request := UpdateOSRequest{
		Name:         "CentOS",
		Version:      "8",
		EndOfSupport: "2024-12-31",
	}

	jsonData, err := json.Marshal(request)
	if err != nil {
		t.Fatalf("Failed to marshal UpdateOSRequest to JSON: %v", err)
	}

	var unmarshaledRequest UpdateOSRequest
	err = json.Unmarshal(jsonData, &unmarshaledRequest)
	if err != nil {
		t.Fatalf("Failed to unmarshal JSON to UpdateOSRequest: %v", err)
	}

	if unmarshaledRequest.Name != request.Name {
		t.Errorf("Expected name '%s', got '%s'", request.Name, unmarshaledRequest.Name)
	}
	if unmarshaledRequest.Version != request.Version {
		t.Errorf("Expected version '%s', got '%s'", request.Version, unmarshaledRequest.Version)
	}
	if unmarshaledRequest.EndOfSupport != request.EndOfSupport {
		t.Errorf("Expected end of support '%s', got '%s'", request.EndOfSupport, unmarshaledRequest.EndOfSupport)
	}
}

func TestUpdateOSRequestPartialUpdate(t *testing.T) {
	// Test partial update with only name
	jsonData := `{"name":"partial-update-os"}`

	var request UpdateOSRequest
	err := json.Unmarshal([]byte(jsonData), &request)
	if err != nil {
		t.Fatalf("Failed to unmarshal partial update JSON: %v", err)
	}

	if request.Name != "partial-update-os" {
		t.Errorf("Expected name 'partial-update-os', got '%s'", request.Name)
	}
	if request.Version != "" {
		t.Errorf("Expected empty version, got '%s'", request.Version)
	}
	if request.EndOfSupport != "" {
		t.Errorf("Expected empty end of support, got '%s'", request.EndOfSupport)
	}
}

func TestOSEndOfSupportDateHandling(t *testing.T) {
	tests := []struct {
		name        string
		dateString  string
		expectedErr bool
	}{
		{
			name:        "valid date with time",
			dateString:  "2027-04-01T00:00:00Z",
			expectedErr: false,
		},
		{
			name:        "invalid date format",
			dateString:  "invalid-date",
			expectedErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jsonData := `{"id":1,"name":"Test","version":"1.0","end_of_support":"` + tt.dateString + `","created_at":"2024-01-01T12:00:00Z","updated_at":"2024-01-01T12:00:00Z"}`

			var os OS
			err := json.Unmarshal([]byte(jsonData), &os)

			if tt.expectedErr && err == nil {
				t.Errorf("Expected error when parsing date '%s', but got none", tt.dateString)
			}
			if !tt.expectedErr && err != nil {
				t.Errorf("Unexpected error when parsing date '%s': %v", tt.dateString, err)
			}
		})
	}
}
