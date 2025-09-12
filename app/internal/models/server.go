package models

import "time"

// Server represents a server in the infrastructure
type Server struct {
	ID        int       `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	OS        string    `json:"os" db:"os"`
	OSVersion string    `json:"os_version" db:"os_version"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// CreateServerRequest represents the request body for creating a server
type CreateServerRequest struct {
	Name      string `json:"name" validate:"required"`
	OS        string `json:"os" validate:"required"`
	OSVersion string `json:"os_version" validate:"required"`
}

// UpdateServerRequest represents the request body for updating a server
type UpdateServerRequest struct {
	Name      string `json:"name,omitempty"`
	OS        string `json:"os,omitempty"`
	OSVersion string `json:"os_version,omitempty"`
}
