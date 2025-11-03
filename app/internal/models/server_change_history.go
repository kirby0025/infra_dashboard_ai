package models

import "time"

// ServerChangeHistory represents a change made to a server
type ServerChangeHistory struct {
	ID           int       `json:"id" db:"id"`
	ServerID     *int      `json:"server_id" db:"server_id"`
	ServerName   string    `json:"server_name" db:"server_name"`
	ChangeType   string    `json:"change_type" db:"change_type"` // 'created', 'os_changed', 'deleted'
	OldOSID      *int      `json:"old_os_id,omitempty" db:"old_os_id"`
	NewOSID      *int      `json:"new_os_id,omitempty" db:"new_os_id"`
	OldOSName    *string   `json:"old_os_name,omitempty" db:"old_os_name"`
	OldOSVersion *string   `json:"old_os_version,omitempty" db:"old_os_version"`
	NewOSName    *string   `json:"new_os_name,omitempty" db:"new_os_name"`
	NewOSVersion *string   `json:"new_os_version,omitempty" db:"new_os_version"`
	ChangedAt    time.Time `json:"changed_at" db:"changed_at"`
}

// ChangeHistoryFilter represents filters for querying change history
type ChangeHistoryFilter struct {
	ServerID   *int
	ChangeType *string
	StartDate  *time.Time
	EndDate    *time.Time
	Limit      int
	Offset     int
}
