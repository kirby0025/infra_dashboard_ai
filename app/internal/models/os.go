package models

import "time"

// OS represents an operating system with support information
type OS struct {
	ID           int       `json:"id" db:"id"`
	Name         string    `json:"name" db:"name"`
	Version      string    `json:"version" db:"version"`
	EndOfSupport time.Time `json:"end_of_support" db:"end_of_support"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
}

// CreateOSRequest represents the request body for creating an OS
type CreateOSRequest struct {
	Name         string `json:"name" validate:"required"`
	Version      string `json:"version" validate:"required"`
	EndOfSupport string `json:"end_of_support" validate:"required"` // Expected format: YYYY-MM-DD
}

// UpdateOSRequest represents the request body for updating an OS
type UpdateOSRequest struct {
	Name         string `json:"name,omitempty"`
	Version      string `json:"version,omitempty"`
	EndOfSupport string `json:"end_of_support,omitempty"` // Expected format: YYYY-MM-DD
}
