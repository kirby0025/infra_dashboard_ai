package models

import "time"

// Server represents a server in the infrastructure
type Server struct {
	ID        int       `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	OSID      int       `json:"os_id" db:"os_id"`
	OS        *OS       `json:"os,omitempty" db:"-"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// CreateServerRequest represents the request body for creating a server
type CreateServerRequest struct {
	Name string `json:"name" validate:"required"`
	OSID int    `json:"os_id" validate:"required"`
}

// UpdateServerRequest represents the request body for updating a server
type UpdateServerRequest struct {
	Name string `json:"name,omitempty"`
	OSID int    `json:"os_id,omitempty"`
}
