package database

import (
	"database/sql"
	"fmt"
	"time"

	"infra-dashboard/internal/config"
	"infra-dashboard/internal/models"

	_ "github.com/lib/pq"
)

// DB wraps the database connection
type DB struct {
	*sql.DB
}

// New creates a new database connection
func New(cfg *config.DatabaseConfig) (*DB, error) {
	db, err := sql.Open("postgres", cfg.GetDSN())
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	// Set connection pool settings
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(5 * time.Minute)

	return &DB{db}, nil
}

// CreateTablesIfNotExist creates the servers table if it doesn't exist
func (db *DB) CreateTablesIfNotExist() error {
	query := `
	CREATE TABLE IF NOT EXISTS servers (
		id SERIAL PRIMARY KEY,
		name VARCHAR(255) NOT NULL UNIQUE,
		os VARCHAR(100) NOT NULL,
		os_version VARCHAR(100) NOT NULL,
		created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
		updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
	);

	CREATE INDEX IF NOT EXISTS idx_servers_name ON servers(name);
	CREATE INDEX IF NOT EXISTS idx_servers_os ON servers(os);
	`

	_, err := db.Exec(query)
	if err != nil {
		return fmt.Errorf("failed to create tables: %w", err)
	}

	return nil
}

// ServerRepository provides database operations for servers
type ServerRepository struct {
	db *DB
}

// NewServerRepository creates a new server repository
func NewServerRepository(db *DB) *ServerRepository {
	return &ServerRepository{db: db}
}

// GetAll retrieves all servers from the database
func (r *ServerRepository) GetAll() ([]models.Server, error) {
	query := `
		SELECT id, name, os, os_version, created_at, updated_at
		FROM servers
		ORDER BY created_at DESC
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to query servers: %w", err)
	}
	defer rows.Close()

	var servers []models.Server
	for rows.Next() {
		var server models.Server
		err := rows.Scan(
			&server.ID,
			&server.Name,
			&server.OS,
			&server.OSVersion,
			&server.CreatedAt,
			&server.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan server: %w", err)
		}
		servers = append(servers, server)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	return servers, nil
}

// GetByID retrieves a server by its ID
func (r *ServerRepository) GetByID(id int) (*models.Server, error) {
	query := `
		SELECT id, name, os, os_version, created_at, updated_at
		FROM servers
		WHERE id = $1
	`

	var server models.Server
	err := r.db.QueryRow(query, id).Scan(
		&server.ID,
		&server.Name,
		&server.OS,
		&server.OSVersion,
		&server.CreatedAt,
		&server.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("server with id %d not found", id)
		}
		return nil, fmt.Errorf("failed to get server: %w", err)
	}

	return &server, nil
}

// Create creates a new server in the database
func (r *ServerRepository) Create(req *models.CreateServerRequest) (*models.Server, error) {
	query := `
		INSERT INTO servers (name, os, os_version, created_at, updated_at)
		VALUES ($1, $2, $3, NOW(), NOW())
		RETURNING id, name, os, os_version, created_at, updated_at
	`

	var server models.Server
	err := r.db.QueryRow(query, req.Name, req.OS, req.OSVersion).Scan(
		&server.ID,
		&server.Name,
		&server.OS,
		&server.OSVersion,
		&server.CreatedAt,
		&server.UpdatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to create server: %w", err)
	}

	return &server, nil
}

// Update updates an existing server in the database
func (r *ServerRepository) Update(id int, req *models.UpdateServerRequest) (*models.Server, error) {
	// Build dynamic update query
	setParts := []string{}
	args := []interface{}{}
	argCount := 1

	if req.Name != "" {
		setParts = append(setParts, fmt.Sprintf("name = $%d", argCount))
		args = append(args, req.Name)
		argCount++
	}

	if req.OS != "" {
		setParts = append(setParts, fmt.Sprintf("os = $%d", argCount))
		args = append(args, req.OS)
		argCount++
	}

	if req.OSVersion != "" {
		setParts = append(setParts, fmt.Sprintf("os_version = $%d", argCount))
		args = append(args, req.OSVersion)
		argCount++
	}

	if len(setParts) == 0 {
		return r.GetByID(id) // No updates, return existing server
	}

	// Always update the updated_at timestamp
	setParts = append(setParts, fmt.Sprintf("updated_at = $%d", argCount))
	args = append(args, time.Now())
	argCount++

	// Add the ID for the WHERE clause
	args = append(args, id)

	// Build the SET clause
	setClause := ""
	for i, part := range setParts {
		if i > 0 {
			setClause += ", "
		}
		setClause += part
	}

	query := fmt.Sprintf(`
		UPDATE servers
		SET %s
		WHERE id = $%d
		RETURNING id, name, os, os_version, created_at, updated_at
	`, setClause, argCount)

	var server models.Server
	err := r.db.QueryRow(query, args...).Scan(
		&server.ID,
		&server.Name,
		&server.OS,
		&server.OSVersion,
		&server.CreatedAt,
		&server.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("server with id %d not found", id)
		}
		return nil, fmt.Errorf("failed to update server: %w", err)
	}

	return &server, nil
}

// Delete removes a server from the database
func (r *ServerRepository) Delete(id int) error {
	query := `DELETE FROM servers WHERE id = $1`

	result, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete server: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("server with id %d not found", id)
	}

	return nil
}
