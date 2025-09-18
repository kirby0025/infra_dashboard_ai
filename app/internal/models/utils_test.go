package models

import (
	"testing"
	"time"
)

func TestOSUtils_GroupOSByFamily(t *testing.T) {
	utils := NewOSUtils()

	oss := []OS{
		{ID: 1, Name: "Ubuntu", Version: "20.04", EndOfSupport: time.Date(2025, 4, 1, 0, 0, 0, 0, time.UTC)},
		{ID: 2, Name: "Ubuntu", Version: "22.04", EndOfSupport: time.Date(2027, 4, 1, 0, 0, 0, 0, time.UTC)},
		{ID: 3, Name: "Debian", Version: "11", EndOfSupport: time.Date(2026, 6, 30, 0, 0, 0, 0, time.UTC)},
		{ID: 4, Name: "Debian", Version: "12", EndOfSupport: time.Date(2028, 6, 30, 0, 0, 0, 0, time.UTC)},
	}

	grouped := utils.GroupOSByFamily(oss)

	if len(grouped) != 2 {
		t.Errorf("Expected 2 OS families, got %d", len(grouped))
	}

	if len(grouped["Ubuntu"]) != 2 {
		t.Errorf("Expected 2 Ubuntu versions, got %d", len(grouped["Ubuntu"]))
	}

	if len(grouped["Debian"]) != 2 {
		t.Errorf("Expected 2 Debian versions, got %d", len(grouped["Debian"]))
	}

	// Check sorting (20.04 should come before 22.04)
	if grouped["Ubuntu"][0].Version != "20.04" {
		t.Errorf("Expected Ubuntu 20.04 first, got %s", grouped["Ubuntu"][0].Version)
	}
}

func TestOSUtils_GetLatestVersionByFamily(t *testing.T) {
	utils := NewOSUtils()

	oss := []OS{
		{ID: 1, Name: "Ubuntu", Version: "20.04", EndOfSupport: time.Date(2025, 4, 1, 0, 0, 0, 0, time.UTC)},
		{ID: 2, Name: "Ubuntu", Version: "22.04", EndOfSupport: time.Date(2027, 4, 1, 0, 0, 0, 0, time.UTC)},
		{ID: 3, Name: "Debian", Version: "11", EndOfSupport: time.Date(2026, 6, 30, 0, 0, 0, 0, time.UTC)},
		{ID: 4, Name: "Debian", Version: "12", EndOfSupport: time.Date(2028, 6, 30, 0, 0, 0, 0, time.UTC)},
	}

	latest := utils.GetLatestVersionByFamily(oss)

	if len(latest) != 2 {
		t.Errorf("Expected 2 OS families, got %d", len(latest))
	}

	if latest["Ubuntu"].Version != "22.04" {
		t.Errorf("Expected latest Ubuntu to be 22.04, got %s", latest["Ubuntu"].Version)
	}

	if latest["Debian"].Version != "12" {
		t.Errorf("Expected latest Debian to be 12, got %s", latest["Debian"].Version)
	}
}

func TestOSUtils_FilterByEndOfSupport(t *testing.T) {
	utils := NewOSUtils()
	now := time.Now()

	oss := []OS{
		{ID: 1, Name: "Ubuntu", Version: "18.04", EndOfSupport: now.AddDate(-1, 0, 0)}, // End of life
		{ID: 2, Name: "Ubuntu", Version: "20.04", EndOfSupport: now.AddDate(0, 3, 0)},  // Ending soon
		{ID: 3, Name: "Ubuntu", Version: "22.04", EndOfSupport: now.AddDate(2, 0, 0)},  // Active
	}

	// Test End of Life filter
	eolOS := utils.FilterByEndOfSupport(oss, SupportStatusEndOfLife)
	if len(eolOS) != 1 {
		t.Errorf("Expected 1 end-of-life OS, got %d", len(eolOS))
	}
	if eolOS[0].Version != "18.04" {
		t.Errorf("Expected 18.04 to be end-of-life, got %s", eolOS[0].Version)
	}

	// Test Ending Soon filter
	endingSoon := utils.FilterByEndOfSupport(oss, SupportStatusEndingSoon)
	if len(endingSoon) != 1 {
		t.Errorf("Expected 1 ending-soon OS, got %d", len(endingSoon))
	}
	if endingSoon[0].Version != "20.04" {
		t.Errorf("Expected 20.04 to be ending soon, got %s", endingSoon[0].Version)
	}

	// Test Active filter
	active := utils.FilterByEndOfSupport(oss, SupportStatusActive)
	if len(active) != 2 { // Both 20.04 and 22.04 are still active (20.04 is ending soon but still active)
		t.Errorf("Expected 2 active OS, got %d", len(active))
	}
}

func TestOSUtils_GetSupportStatusString(t *testing.T) {
	utils := NewOSUtils()
	now := time.Now()

	tests := []struct {
		name         string
		endOfSupport time.Time
		expected     string
	}{
		{
			name:         "End of Life",
			endOfSupport: now.AddDate(-1, 0, 0),
			expected:     "End of Life",
		},
		{
			name:         "Ending Soon",
			endOfSupport: now.AddDate(0, 3, 0),
			expected:     "Ending Soon",
		},
		{
			name:         "Supported",
			endOfSupport: now.AddDate(2, 0, 0),
			expected:     "Supported",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os := OS{EndOfSupport: tt.endOfSupport}
			status := utils.GetSupportStatusString(os)
			if status != tt.expected {
				t.Errorf("Expected %s, got %s", tt.expected, status)
			}
		})
	}
}

func TestOSUtils_GetDaysUntilEndOfSupport(t *testing.T) {
	utils := NewOSUtils()
	now := time.Now()

	os := OS{EndOfSupport: now.AddDate(0, 0, 30)} // 30 days from now
	days := utils.GetDaysUntilEndOfSupport(os)

	// Allow for some variation due to timing
	if days < 29 || days > 31 {
		t.Errorf("Expected approximately 30 days, got %d", days)
	}
}

func TestServerUtils_GroupServersByOS(t *testing.T) {
	utils := NewServerUtils()

	servers := []Server{
		{
			ID:   1,
			Name: "server1",
			OSID: 1,
			OS:   &OS{Name: "Ubuntu", Version: "20.04"},
		},
		{
			ID:   2,
			Name: "server2",
			OSID: 1,
			OS:   &OS{Name: "Ubuntu", Version: "20.04"},
		},
		{
			ID:   3,
			Name: "server3",
			OSID: 2,
			OS:   &OS{Name: "Ubuntu", Version: "22.04"},
		},
	}

	grouped := utils.GroupServersByOS(servers)

	if len(grouped) != 2 {
		t.Errorf("Expected 2 OS groups, got %d", len(grouped))
	}

	if len(grouped["Ubuntu 20.04"]) != 2 {
		t.Errorf("Expected 2 servers on Ubuntu 20.04, got %d", len(grouped["Ubuntu 20.04"]))
	}

	if len(grouped["Ubuntu 22.04"]) != 1 {
		t.Errorf("Expected 1 server on Ubuntu 22.04, got %d", len(grouped["Ubuntu 22.04"]))
	}
}

func TestServerUtils_GetServersWithEndOfLifeOS(t *testing.T) {
	utils := NewServerUtils()
	now := time.Now()

	servers := []Server{
		{
			ID:   1,
			Name: "server1",
			OSID: 1,
			OS:   &OS{Name: "Ubuntu", Version: "18.04", EndOfSupport: now.AddDate(-1, 0, 0)},
		},
		{
			ID:   2,
			Name: "server2",
			OSID: 2,
			OS:   &OS{Name: "Ubuntu", Version: "20.04", EndOfSupport: now.AddDate(1, 0, 0)},
		},
	}

	eolServers := utils.GetServersWithEndOfLifeOS(servers)

	if len(eolServers) != 1 {
		t.Errorf("Expected 1 server with end-of-life OS, got %d", len(eolServers))
	}

	if eolServers[0].Name != "server1" {
		t.Errorf("Expected server1 to have end-of-life OS, got %s", eolServers[0].Name)
	}
}

func TestServerUtils_GetOSDistribution(t *testing.T) {
	utils := NewServerUtils()

	servers := []Server{
		{ID: 1, Name: "server1", OSID: 1, OS: &OS{Name: "Ubuntu", Version: "20.04"}},
		{ID: 2, Name: "server2", OSID: 1, OS: &OS{Name: "Ubuntu", Version: "20.04"}},
		{ID: 3, Name: "server3", OSID: 2, OS: &OS{Name: "Ubuntu", Version: "22.04"}},
		{ID: 4, Name: "server4", OSID: 3, OS: &OS{Name: "Debian", Version: "11"}},
	}

	distribution := utils.GetOSDistribution(servers)

	expected := map[string]int{
		"Ubuntu 20.04": 2,
		"Ubuntu 22.04": 1,
		"Debian 11":    1,
	}

	if len(distribution) != len(expected) {
		t.Errorf("Expected %d OS types, got %d", len(expected), len(distribution))
	}

	for os, expectedCount := range expected {
		if count, exists := distribution[os]; !exists || count != expectedCount {
			t.Errorf("Expected %d servers for %s, got %d", expectedCount, os, count)
		}
	}
}

func TestServerUtils_GetOSFamilyDistribution(t *testing.T) {
	utils := NewServerUtils()

	servers := []Server{
		{ID: 1, Name: "server1", OSID: 1, OS: &OS{Name: "Ubuntu", Version: "20.04"}},
		{ID: 2, Name: "server2", OSID: 2, OS: &OS{Name: "Ubuntu", Version: "22.04"}},
		{ID: 3, Name: "server3", OSID: 3, OS: &OS{Name: "Debian", Version: "11"}},
		{ID: 4, Name: "server4", OSID: 4, OS: &OS{Name: "CentOS", Version: "7"}},
	}

	distribution := utils.GetOSFamilyDistribution(servers)

	expected := map[string]int{
		"Ubuntu": 2,
		"Debian": 1,
		"CentOS": 1,
	}

	if len(distribution) != len(expected) {
		t.Errorf("Expected %d OS families, got %d", len(expected), len(distribution))
	}

	for family, expectedCount := range expected {
		if count, exists := distribution[family]; !exists || count != expectedCount {
			t.Errorf("Expected %d servers for %s family, got %d", expectedCount, family, count)
		}
	}
}

func TestServerUtils_FindServersByOSID(t *testing.T) {
	utils := NewServerUtils()

	servers := []Server{
		{ID: 1, Name: "server1", OSID: 1},
		{ID: 2, Name: "server2", OSID: 1},
		{ID: 3, Name: "server3", OSID: 2},
	}

	matches := utils.FindServersByOSID(servers, 1)

	if len(matches) != 2 {
		t.Errorf("Expected 2 servers with OSID 1, got %d", len(matches))
	}

	for _, server := range matches {
		if server.OSID != 1 {
			t.Errorf("Expected OSID 1, got %d", server.OSID)
		}
	}
}

func TestServerUtils_ValidateServerName(t *testing.T) {
	utils := NewServerUtils()

	tests := []struct {
		name      string
		input     string
		expectErr bool
	}{
		{
			name:      "Valid name",
			input:     "web-server-01",
			expectErr: false,
		},
		{
			name:      "Valid name with dots",
			input:     "web.server.01",
			expectErr: false,
		},
		{
			name:      "Empty name",
			input:     "",
			expectErr: true,
		},
		{
			name:      "Name starting with hyphen",
			input:     "-invalid",
			expectErr: true,
		},
		{
			name:      "Name ending with hyphen",
			input:     "invalid-",
			expectErr: true,
		},
		{
			name:      "Name with invalid characters",
			input:     "server@01",
			expectErr: true,
		},
		{
			name:      "Name too long",
			input:     string(make([]byte, 254)),
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := utils.ValidateServerName(tt.input)
			if tt.expectErr && err == nil {
				t.Errorf("Expected error for input %s, but got none", tt.input)
			}
			if !tt.expectErr && err != nil {
				t.Errorf("Unexpected error for input %s: %v", tt.input, err)
			}
		})
	}
}

func TestComplianceUtils_GenerateComplianceReport(t *testing.T) {
	utils := NewComplianceUtils()
	now := time.Now()

	servers := []Server{
		{
			ID:   1,
			Name: "server1",
			OSID: 1,
			OS:   &OS{Name: "Ubuntu", Version: "18.04", EndOfSupport: now.AddDate(-1, 0, 0)}, // EOL
		},
		{
			ID:   2,
			Name: "server2",
			OSID: 2,
			OS:   &OS{Name: "Ubuntu", Version: "20.04", EndOfSupport: now.AddDate(0, 3, 0)}, // Ending soon
		},
		{
			ID:   3,
			Name: "server3",
			OSID: 3,
			OS:   &OS{Name: "Ubuntu", Version: "22.04", EndOfSupport: now.AddDate(2, 0, 0)}, // Supported
		},
	}

	report := utils.GenerateComplianceReport(servers)

	if report.TotalServers != 3 {
		t.Errorf("Expected 3 total servers, got %d", report.TotalServers)
	}

	if report.EndOfLifeServers != 1 {
		t.Errorf("Expected 1 end-of-life server, got %d", report.EndOfLifeServers)
	}

	if report.EndingSoonServers != 1 {
		t.Errorf("Expected 1 ending-soon server, got %d", report.EndingSoonServers)
	}

	if report.SupportedServers != 1 {
		t.Errorf("Expected 1 supported server, got %d", report.SupportedServers)
	}

	if len(report.EndOfLifeList) != 1 {
		t.Errorf("Expected 1 server in end-of-life list, got %d", len(report.EndOfLifeList))
	}

	if len(report.EndingSoonList) != 1 {
		t.Errorf("Expected 1 server in ending-soon list, got %d", len(report.EndingSoonList))
	}
}

func TestComplianceUtils_GetComplianceScore(t *testing.T) {
	utils := NewComplianceUtils()
	now := time.Now()

	tests := []struct {
		name     string
		servers  []Server
		expected float64
	}{
		{
			name:     "No servers",
			servers:  []Server{},
			expected: 100.0,
		},
		{
			name: "All supported servers",
			servers: []Server{
				{ID: 1, OSID: 1, OS: &OS{EndOfSupport: now.AddDate(2, 0, 0)}},
				{ID: 2, OSID: 2, OS: &OS{EndOfSupport: now.AddDate(2, 0, 0)}},
			},
			expected: 100.0,
		},
		{
			name: "Mixed compliance",
			servers: []Server{
				{ID: 1, OSID: 1, OS: &OS{EndOfSupport: now.AddDate(-1, 0, 0)}}, // EOL (penalty: 2)
				{ID: 2, OSID: 2, OS: &OS{EndOfSupport: now.AddDate(0, 3, 0)}},  // Ending soon (penalty: 0.5)
				{ID: 3, OSID: 3, OS: &OS{EndOfSupport: now.AddDate(2, 0, 0)}},  // Supported (penalty: 0)
				{ID: 4, OSID: 4, OS: &OS{EndOfSupport: now.AddDate(2, 0, 0)}},  // Supported (penalty: 0)
			},
			expected: 37.5, // (4 - 2 - 0.5) / 4 * 100 = 37.5
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			score := utils.GetComplianceScore(tt.servers)
			if score != tt.expected {
				t.Errorf("Expected compliance score %.1f, got %.1f", tt.expected, score)
			}
		})
	}
}

func TestComplianceUtils_GetRecommendations(t *testing.T) {
	utils := NewComplianceUtils()
	now := time.Now()

	servers := []Server{
		{
			ID:   1,
			Name: "server1",
			OSID: 1,
			OS:   &OS{Name: "Ubuntu", Version: "18.04", EndOfSupport: now.AddDate(-1, 0, 0)},
		},
		{
			ID:   2,
			Name: "server2",
			OSID: 2,
			OS:   &OS{Name: "Ubuntu", Version: "20.04", EndOfSupport: now.AddDate(0, 3, 0)},
		},
	}

	allOS := []OS{
		{ID: 1, Name: "Ubuntu", Version: "18.04", EndOfSupport: now.AddDate(-1, 0, 0)},
		{ID: 2, Name: "Ubuntu", Version: "20.04", EndOfSupport: now.AddDate(0, 3, 0)},
		{ID: 3, Name: "Ubuntu", Version: "22.04", EndOfSupport: now.AddDate(2, 0, 0)},
	}

	recommendations := utils.GetRecommendations(servers, allOS)

	if len(recommendations) < 2 {
		t.Errorf("Expected at least 2 recommendations, got %d", len(recommendations))
	}

	// Check that critical and warning recommendations are present
	hasCritical := false
	hasWarning := false
	for _, rec := range recommendations {
		if len(rec) > 8 && rec[:8] == "CRITICAL" {
			hasCritical = true
		}
		if len(rec) > 7 && rec[:7] == "WARNING" {
			hasWarning = true
		}
	}

	if !hasCritical {
		t.Error("Expected a CRITICAL recommendation for end-of-life servers")
	}

	if !hasWarning {
		t.Error("Expected a WARNING recommendation for ending-soon servers")
	}
}
