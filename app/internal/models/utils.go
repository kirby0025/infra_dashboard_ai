package models

import (
	"fmt"
	"sort"
	"time"
)

// OSUtils provides utility functions for operating system operations
type OSUtils struct{}

// NewOSUtils creates a new OSUtils instance
func NewOSUtils() *OSUtils {
	return &OSUtils{}
}

// GroupOSByFamily groups operating systems by their family/distribution name
func (u *OSUtils) GroupOSByFamily(oss []OS) map[string][]OS {
	grouped := make(map[string][]OS)

	for _, os := range oss {
		grouped[os.Name] = append(grouped[os.Name], os)
	}

	// Sort versions within each group
	for family := range grouped {
		sort.Slice(grouped[family], func(i, j int) bool {
			return grouped[family][i].Version < grouped[family][j].Version
		})
	}

	return grouped
}

// GetLatestVersionByFamily returns the latest version for each OS family
func (u *OSUtils) GetLatestVersionByFamily(oss []OS) map[string]OS {
	grouped := u.GroupOSByFamily(oss)
	latest := make(map[string]OS)

	for family, versions := range grouped {
		if len(versions) > 0 {
			// Get the last element (highest version after sorting)
			latest[family] = versions[len(versions)-1]
		}
	}

	return latest
}

// FilterByEndOfSupport filters operating systems by their support status
func (u *OSUtils) FilterByEndOfSupport(oss []OS, status SupportStatus) []OS {
	now := time.Now()
	var filtered []OS

	for _, os := range oss {
		switch status {
		case SupportStatusActive:
			if os.EndOfSupport.After(now) {
				filtered = append(filtered, os)
			}
		case SupportStatusEndOfLife:
			if os.EndOfSupport.Before(now) {
				filtered = append(filtered, os)
			}
		case SupportStatusEndingSoon:
			// Consider "ending soon" as within 6 months
			sixMonthsFromNow := now.AddDate(0, 6, 0)
			if os.EndOfSupport.After(now) && os.EndOfSupport.Before(sixMonthsFromNow) {
				filtered = append(filtered, os)
			}
		}
	}

	return filtered
}

// GetSupportStatusString returns a human-readable support status
func (u *OSUtils) GetSupportStatusString(os OS) string {
	now := time.Now()

	if os.EndOfSupport.Before(now) {
		return "End of Life"
	}

	sixMonthsFromNow := now.AddDate(0, 6, 0)
	if os.EndOfSupport.Before(sixMonthsFromNow) {
		return "Ending Soon"
	}

	return "Supported"
}

// GetDaysUntilEndOfSupport returns the number of days until end of support
func (u *OSUtils) GetDaysUntilEndOfSupport(os OS) int {
	now := time.Now()
	diff := os.EndOfSupport.Sub(now)
	return int(diff.Hours() / 24)
}

// SupportStatus represents the support status of an operating system
type SupportStatus int

const (
	SupportStatusActive SupportStatus = iota
	SupportStatusEndOfLife
	SupportStatusEndingSoon
)

// ServerUtils provides utility functions for server operations
type ServerUtils struct{}

// NewServerUtils creates a new ServerUtils instance
func NewServerUtils() *ServerUtils {
	return &ServerUtils{}
}

// GroupServersByOS groups servers by their operating system
func (u *ServerUtils) GroupServersByOS(servers []Server) map[string][]Server {
	grouped := make(map[string][]Server)

	for _, server := range servers {
		if server.OS != nil {
			key := fmt.Sprintf("%s %s", server.OS.Name, server.OS.Version)
			grouped[key] = append(grouped[key], server)
		}
	}

	return grouped
}

// GetServersWithEndOfLifeOS returns servers running end-of-life operating systems
func (u *ServerUtils) GetServersWithEndOfLifeOS(servers []Server) []Server {
	now := time.Now()
	var eolServers []Server

	for _, server := range servers {
		if server.OS != nil && server.OS.EndOfSupport.Before(now) {
			eolServers = append(eolServers, server)
		}
	}

	return eolServers
}

// GetServersWithEndingSoonOS returns servers with OS support ending soon
func (u *ServerUtils) GetServersWithEndingSoonOS(servers []Server) []Server {
	now := time.Now()
	sixMonthsFromNow := now.AddDate(0, 6, 0)
	var endingSoonServers []Server

	for _, server := range servers {
		if server.OS != nil &&
			server.OS.EndOfSupport.After(now) &&
			server.OS.EndOfSupport.Before(sixMonthsFromNow) {
			endingSoonServers = append(endingSoonServers, server)
		}
	}

	return endingSoonServers
}

// GetOSDistribution returns a count of servers by operating system
func (u *ServerUtils) GetOSDistribution(servers []Server) map[string]int {
	distribution := make(map[string]int)

	for _, server := range servers {
		if server.OS != nil {
			key := fmt.Sprintf("%s %s", server.OS.Name, server.OS.Version)
			distribution[key]++
		}
	}

	return distribution
}

// GetOSFamilyDistribution returns a count of servers by OS family
func (u *ServerUtils) GetOSFamilyDistribution(servers []Server) map[string]int {
	distribution := make(map[string]int)

	for _, server := range servers {
		if server.OS != nil {
			distribution[server.OS.Name]++
		}
	}

	return distribution
}

// FindServersByOSID returns servers using a specific operating system ID
func (u *ServerUtils) FindServersByOSID(servers []Server, osID int) []Server {
	var matches []Server

	for _, server := range servers {
		if server.OSID == osID {
			matches = append(matches, server)
		}
	}

	return matches
}

// ValidateServerName checks if a server name follows naming conventions
func (u *ServerUtils) ValidateServerName(name string) error {
	if len(name) == 0 {
		return fmt.Errorf("server name cannot be empty")
	}

	if len(name) > 253 {
		return fmt.Errorf("server name cannot exceed 253 characters")
	}

	// Check for valid characters (alphanumeric, hyphens, dots)
	for _, char := range name {
		if !((char >= 'a' && char <= 'z') ||
			(char >= 'A' && char <= 'Z') ||
			(char >= '0' && char <= '9') ||
			char == '-' || char == '.') {
			return fmt.Errorf("server name contains invalid character: %c", char)
		}
	}

	// Cannot start or end with hyphen or dot
	if name[0] == '-' || name[0] == '.' ||
		name[len(name)-1] == '-' || name[len(name)-1] == '.' {
		return fmt.Errorf("server name cannot start or end with hyphen or dot")
	}

	return nil
}

// ComplianceReport represents a compliance analysis report
type ComplianceReport struct {
	TotalServers         int            `json:"total_servers"`
	SupportedServers     int            `json:"supported_servers"`
	EndOfLifeServers     int            `json:"end_of_life_servers"`
	EndingSoonServers    int            `json:"ending_soon_servers"`
	OSDistribution       map[string]int `json:"os_distribution"`
	OSFamilyDistribution map[string]int `json:"os_family_distribution"`
	EndOfLifeList        []Server       `json:"end_of_life_list"`
	EndingSoonList       []Server       `json:"ending_soon_list"`
	GeneratedAt          time.Time      `json:"generated_at"`
}

// ComplianceUtils provides utility functions for compliance reporting
type ComplianceUtils struct {
	serverUtils *ServerUtils
	osUtils     *OSUtils
}

// NewComplianceUtils creates a new ComplianceUtils instance
func NewComplianceUtils() *ComplianceUtils {
	return &ComplianceUtils{
		serverUtils: NewServerUtils(),
		osUtils:     NewOSUtils(),
	}
}

// GenerateComplianceReport creates a comprehensive compliance report
func (u *ComplianceUtils) GenerateComplianceReport(servers []Server) ComplianceReport {
	endOfLifeServers := u.serverUtils.GetServersWithEndOfLifeOS(servers)
	endingSoonServers := u.serverUtils.GetServersWithEndingSoonOS(servers)

	report := ComplianceReport{
		TotalServers:         len(servers),
		SupportedServers:     len(servers) - len(endOfLifeServers) - len(endingSoonServers),
		EndOfLifeServers:     len(endOfLifeServers),
		EndingSoonServers:    len(endingSoonServers),
		OSDistribution:       u.serverUtils.GetOSDistribution(servers),
		OSFamilyDistribution: u.serverUtils.GetOSFamilyDistribution(servers),
		EndOfLifeList:        endOfLifeServers,
		EndingSoonList:       endingSoonServers,
		GeneratedAt:          time.Now(),
	}

	return report
}

// GetComplianceScore calculates a compliance score (0-100)
func (u *ComplianceUtils) GetComplianceScore(servers []Server) float64 {
	if len(servers) == 0 {
		return 100.0
	}

	endOfLifeServers := u.serverUtils.GetServersWithEndOfLifeOS(servers)
	endingSoonServers := u.serverUtils.GetServersWithEndingSoonOS(servers)

	// End of life servers heavily impact score, ending soon servers have moderate impact
	penalty := float64(len(endOfLifeServers))*2 + float64(len(endingSoonServers))*0.5
	score := (float64(len(servers)) - penalty) / float64(len(servers)) * 100

	if score < 0 {
		return 0
	}

	return score
}

// GetRecommendations provides upgrade recommendations
func (u *ComplianceUtils) GetRecommendations(servers []Server, allOS []OS) []string {
	var recommendations []string

	endOfLifeServers := u.serverUtils.GetServersWithEndOfLifeOS(servers)
	endingSoonServers := u.serverUtils.GetServersWithEndingSoonOS(servers)

	if len(endOfLifeServers) > 0 {
		recommendations = append(recommendations,
			fmt.Sprintf("CRITICAL: %d servers are running end-of-life operating systems and need immediate updates", len(endOfLifeServers)))
	}

	if len(endingSoonServers) > 0 {
		recommendations = append(recommendations,
			fmt.Sprintf("WARNING: %d servers are running operating systems that will reach end-of-life within 6 months", len(endingSoonServers)))
	}

	// Find OS families with newer versions available
	latestVersions := u.osUtils.GetLatestVersionByFamily(allOS)
	osDistribution := u.serverUtils.GetOSDistribution(servers)

	for osVersion, count := range osDistribution {
		// This is a simplified check - in practice, you'd want more sophisticated version comparison
		for family, latest := range latestVersions {
			latestKey := fmt.Sprintf("%s %s", latest.Name, latest.Version)
			if osVersion != latestKey && latest.Name == family {
				recommendations = append(recommendations,
					fmt.Sprintf("SUGGESTION: Consider upgrading %d servers from %s to %s", count, osVersion, latestKey))
			}
		}
	}

	return recommendations
}
