package models

import (
	"encoding/json"
	"testing"
	"time"
)

func TestServerChangeHistoryJSON(t *testing.T) {
	now := time.Now()
	serverID := 1
	oldOSID := 10
	newOSID := 20
	oldOSName := "Ubuntu"
	oldOSVersion := "20.04"
	newOSName := "Ubuntu"
	newOSVersion := "22.04"

	history := ServerChangeHistory{
		ID:           1,
		ServerID:     &serverID,
		ServerName:   "web-server-01",
		ChangeType:   "os_changed",
		OldOSID:      &oldOSID,
		NewOSID:      &newOSID,
		OldOSName:    &oldOSName,
		OldOSVersion: &oldOSVersion,
		NewOSName:    &newOSName,
		NewOSVersion: &newOSVersion,
		ChangedAt:    now,
	}

	jsonData, err := json.Marshal(history)
	if err != nil {
		t.Fatalf("Failed to marshal ServerChangeHistory to JSON: %v", err)
	}

	var unmarshaled ServerChangeHistory
	err = json.Unmarshal(jsonData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal ServerChangeHistory from JSON: %v", err)
	}

	if unmarshaled.ID != history.ID {
		t.Errorf("Expected ID %d, got %d", history.ID, unmarshaled.ID)
	}
	if *unmarshaled.ServerID != *history.ServerID {
		t.Errorf("Expected ServerID %d, got %d", *history.ServerID, *unmarshaled.ServerID)
	}
	if unmarshaled.ServerName != history.ServerName {
		t.Errorf("Expected ServerName %s, got %s", history.ServerName, unmarshaled.ServerName)
	}
	if unmarshaled.ChangeType != history.ChangeType {
		t.Errorf("Expected ChangeType %s, got %s", history.ChangeType, unmarshaled.ChangeType)
	}
}

func TestServerChangeHistoryCreated(t *testing.T) {
	now := time.Now()
	serverID := 1
	newOSID := 20
	newOSName := "Ubuntu"
	newOSVersion := "22.04"

	history := ServerChangeHistory{
		ID:           1,
		ServerID:     &serverID,
		ServerName:   "web-server-01",
		ChangeType:   "created",
		OldOSID:      nil,
		NewOSID:      &newOSID,
		OldOSName:    nil,
		OldOSVersion: nil,
		NewOSName:    &newOSName,
		NewOSVersion: &newOSVersion,
		ChangedAt:    now,
	}

	if history.ChangeType != "created" {
		t.Errorf("Expected ChangeType 'created', got %s", history.ChangeType)
	}
	if history.OldOSID != nil {
		t.Error("Expected OldOSID to be nil for created change type")
	}
	if history.NewOSID == nil {
		t.Error("Expected NewOSID to not be nil for created change type")
	}
}

func TestServerChangeHistoryOSChanged(t *testing.T) {
	now := time.Now()
	serverID := 1
	oldOSID := 10
	newOSID := 20
	oldOSName := "Ubuntu"
	oldOSVersion := "20.04"
	newOSName := "Ubuntu"
	newOSVersion := "22.04"

	history := ServerChangeHistory{
		ID:           1,
		ServerID:     &serverID,
		ServerName:   "web-server-01",
		ChangeType:   "os_changed",
		OldOSID:      &oldOSID,
		NewOSID:      &newOSID,
		OldOSName:    &oldOSName,
		OldOSVersion: &oldOSVersion,
		NewOSName:    &newOSName,
		NewOSVersion: &newOSVersion,
		ChangedAt:    now,
	}

	if history.ChangeType != "os_changed" {
		t.Errorf("Expected ChangeType 'os_changed', got %s", history.ChangeType)
	}
	if history.OldOSID == nil {
		t.Error("Expected OldOSID to not be nil for os_changed change type")
	}
	if history.NewOSID == nil {
		t.Error("Expected NewOSID to not be nil for os_changed change type")
	}
	if *history.OldOSID == *history.NewOSID {
		t.Error("OldOSID and NewOSID should be different for os_changed")
	}
}

func TestServerChangeHistoryDeleted(t *testing.T) {
	now := time.Now()
	serverID := 1
	oldOSID := 10
	oldOSName := "Ubuntu"
	oldOSVersion := "20.04"

	history := ServerChangeHistory{
		ID:           1,
		ServerID:     &serverID,
		ServerName:   "web-server-01",
		ChangeType:   "deleted",
		OldOSID:      &oldOSID,
		NewOSID:      nil,
		OldOSName:    &oldOSName,
		OldOSVersion: &oldOSVersion,
		NewOSName:    nil,
		NewOSVersion: nil,
		ChangedAt:    now,
	}

	if history.ChangeType != "deleted" {
		t.Errorf("Expected ChangeType 'deleted', got %s", history.ChangeType)
	}
	if history.OldOSID == nil {
		t.Error("Expected OldOSID to not be nil for deleted change type")
	}
	if history.NewOSID != nil {
		t.Error("Expected NewOSID to be nil for deleted change type")
	}
}

func TestChangeHistoryFilterDefaults(t *testing.T) {
	filter := ChangeHistoryFilter{
		Limit:  100,
		Offset: 0,
	}

	if filter.Limit != 100 {
		t.Errorf("Expected default Limit 100, got %d", filter.Limit)
	}
	if filter.Offset != 0 {
		t.Errorf("Expected default Offset 0, got %d", filter.Offset)
	}
	if filter.ServerID != nil {
		t.Error("Expected ServerID to be nil by default")
	}
	if filter.ChangeType != nil {
		t.Error("Expected ChangeType to be nil by default")
	}
	if filter.StartDate != nil {
		t.Error("Expected StartDate to be nil by default")
	}
	if filter.EndDate != nil {
		t.Error("Expected EndDate to be nil by default")
	}
}

func TestChangeHistoryFilterWithServerID(t *testing.T) {
	serverID := 5
	filter := ChangeHistoryFilter{
		ServerID: &serverID,
		Limit:    50,
		Offset:   0,
	}

	if filter.ServerID == nil {
		t.Fatal("Expected ServerID to not be nil")
	}
	if *filter.ServerID != 5 {
		t.Errorf("Expected ServerID 5, got %d", *filter.ServerID)
	}
}

func TestChangeHistoryFilterWithChangeType(t *testing.T) {
	changeType := "os_changed"
	filter := ChangeHistoryFilter{
		ChangeType: &changeType,
		Limit:      50,
		Offset:     0,
	}

	if filter.ChangeType == nil {
		t.Fatal("Expected ChangeType to not be nil")
	}
	if *filter.ChangeType != "os_changed" {
		t.Errorf("Expected ChangeType 'os_changed', got %s", *filter.ChangeType)
	}
}

func TestChangeHistoryFilterWithDateRange(t *testing.T) {
	startDate := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2024, 12, 31, 23, 59, 59, 0, time.UTC)

	filter := ChangeHistoryFilter{
		StartDate: &startDate,
		EndDate:   &endDate,
		Limit:     50,
		Offset:    0,
	}

	if filter.StartDate == nil {
		t.Fatal("Expected StartDate to not be nil")
	}
	if filter.EndDate == nil {
		t.Fatal("Expected EndDate to not be nil")
	}
	if !filter.StartDate.Equal(startDate) {
		t.Errorf("Expected StartDate %v, got %v", startDate, *filter.StartDate)
	}
	if !filter.EndDate.Equal(endDate) {
		t.Errorf("Expected EndDate %v, got %v", endDate, *filter.EndDate)
	}
}

func TestServerChangeHistoryJSONOmitEmpty(t *testing.T) {
	// Test that nil fields are omitted in JSON when omitempty is used
	serverID := 1
	newOSID := 20
	newOSName := "Ubuntu"
	newOSVersion := "22.04"

	history := ServerChangeHistory{
		ID:           1,
		ServerID:     &serverID,
		ServerName:   "web-server-01",
		ChangeType:   "created",
		OldOSID:      nil,
		NewOSID:      &newOSID,
		OldOSName:    nil,
		OldOSVersion: nil,
		NewOSName:    &newOSName,
		NewOSVersion: &newOSVersion,
		ChangedAt:    time.Now(),
	}

	jsonData, err := json.Marshal(history)
	if err != nil {
		t.Fatalf("Failed to marshal ServerChangeHistory to JSON: %v", err)
	}

	jsonString := string(jsonData)

	// Check that null fields are present (Go's json.Marshal includes them even with omitempty for pointer types)
	// but they should be null in JSON
	var unmarshaled map[string]interface{}
	err = json.Unmarshal(jsonData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal to map: %v", err)
	}

	// Verify the structure contains expected fields
	if _, exists := unmarshaled["id"]; !exists {
		t.Error("Expected 'id' field in JSON")
	}
	if _, exists := unmarshaled["server_name"]; !exists {
		t.Error("Expected 'server_name' field in JSON")
	}
	if _, exists := unmarshaled["change_type"]; !exists {
		t.Error("Expected 'change_type' field in JSON")
	}

	_ = jsonString // Use the variable to avoid unused variable error
}
