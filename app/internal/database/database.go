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
	CREATE TABLE IF NOT EXISTS operating_systems (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		version VARCHAR(100) NOT NULL,
		end_of_support DATE NOT NULL,
		created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
		updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
		UNIQUE(name, version)
	);

	CREATE TABLE IF NOT EXISTS servers (
		id SERIAL PRIMARY KEY,
		name VARCHAR(255) NOT NULL UNIQUE,
		os_id INTEGER NOT NULL,
		created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
		updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
		FOREIGN KEY (os_id) REFERENCES operating_systems(id)
	);

	CREATE INDEX IF NOT EXISTS idx_servers_name ON servers(name);
	CREATE INDEX IF NOT EXISTS idx_servers_os_id ON servers(os_id);
	CREATE INDEX IF NOT EXISTS idx_os_name ON operating_systems(name);
	CREATE INDEX IF NOT EXISTS idx_os_end_of_support ON operating_systems(end_of_support);
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
		SELECT s.id, s.name, s.os_id, s.created_at, s.updated_at,
		       os.id, os.name, os.version, os.end_of_support, os.created_at, os.updated_at
		FROM servers s
		JOIN operating_systems os ON s.os_id = os.id
		ORDER BY s.created_at DESC
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to query servers: %w", err)
	}
	defer rows.Close()

	var servers []models.Server
	for rows.Next() {
		var server models.Server
		var os models.OS
		err := rows.Scan(
			&server.ID,
			&server.Name,
			&server.OSID,
			&server.CreatedAt,
			&server.UpdatedAt,
			&os.ID,
			&os.Name,
			&os.Version,
			&os.EndOfSupport,
			&os.CreatedAt,
			&os.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan server: %w", err)
		}
		server.OS = &os
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
		SELECT s.id, s.name, s.os_id, s.created_at, s.updated_at,
		       os.id, os.name, os.version, os.end_of_support, os.created_at, os.updated_at
		FROM servers s
		JOIN operating_systems os ON s.os_id = os.id
		WHERE s.id = $1
	`

	var server models.Server
	var os models.OS
	err := r.db.QueryRow(query, id).Scan(
		&server.ID,
		&server.Name,
		&server.OSID,
		&server.CreatedAt,
		&server.UpdatedAt,
		&os.ID,
		&os.Name,
		&os.Version,
		&os.EndOfSupport,
		&os.CreatedAt,
		&os.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("server with id %d not found", id)
		}
		return nil, fmt.Errorf("failed to get server: %w", err)
	}

	server.OS = &os
	return &server, nil
}

// Create creates a new server in the database
func (r *ServerRepository) Create(req *models.CreateServerRequest) (*models.Server, error) {
	// First verify the OS exists
	var osExists bool
	checkQuery := `SELECT EXISTS(SELECT 1 FROM operating_systems WHERE id = $1)`
	err := r.db.QueryRow(checkQuery, req.OSID).Scan(&osExists)
	if err != nil {
		return nil, fmt.Errorf("failed to check OS existence: %w", err)
	}
	if !osExists {
		return nil, fmt.Errorf("operating system with id %d does not exist", req.OSID)
	}

	query := `
		INSERT INTO servers (name, os_id, created_at, updated_at)
		VALUES ($1, $2, NOW(), NOW())
		RETURNING id, name, os_id, created_at, updated_at
	`

	var server models.Server
	err = r.db.QueryRow(query, req.Name, req.OSID).Scan(
		&server.ID,
		&server.Name,
		&server.OSID,
		&server.CreatedAt,
		&server.UpdatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to create server: %w", err)
	}

	// Fetch the full server with OS details
	return r.GetByID(server.ID)
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

	if req.OSID != 0 {
		// First verify the OS exists
		var osExists bool
		checkQuery := `SELECT EXISTS(SELECT 1 FROM operating_systems WHERE id = $1)`
		err := r.db.QueryRow(checkQuery, req.OSID).Scan(&osExists)
		if err != nil {
			return nil, fmt.Errorf("failed to check OS existence: %w", err)
		}
		if !osExists {
			return nil, fmt.Errorf("operating system with id %d does not exist", req.OSID)
		}

		setParts = append(setParts, fmt.Sprintf("os_id = $%d", argCount))
		args = append(args, req.OSID)
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
		RETURNING id, name, os_id, created_at, updated_at
	`, setClause, argCount)

	var server models.Server
	err := r.db.QueryRow(query, args...).Scan(
		&server.ID,
		&server.Name,
		&server.OSID,
		&server.CreatedAt,
		&server.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("server with id %d not found", id)
		}
		return nil, fmt.Errorf("failed to update server: %w", err)
	}

	// Fetch the full server with OS details
	return r.GetByID(server.ID)
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

// OSRepository provides database operations for operating systems
type OSRepository struct {
	db *DB
}

// NewOSRepository creates a new OS repository
func NewOSRepository(db *DB) *OSRepository {
	return &OSRepository{db: db}
}

// GetAll retrieves all operating systems from the database
func (r *OSRepository) GetAll() ([]models.OS, error) {
	query := `
		SELECT id, name, version, end_of_support, created_at, updated_at
		FROM operating_systems
		ORDER BY name, version
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to query operating systems: %w", err)
	}
	defer rows.Close()

	var oss []models.OS
	for rows.Next() {
		var os models.OS
		err := rows.Scan(
			&os.ID,
			&os.Name,
			&os.Version,
			&os.EndOfSupport,
			&os.CreatedAt,
			&os.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan OS: %w", err)
		}
		oss = append(oss, os)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	return oss, nil
}

// GetByID retrieves an operating system by its ID
func (r *OSRepository) GetByID(id int) (*models.OS, error) {
	query := `
		SELECT id, name, version, end_of_support, created_at, updated_at
		FROM operating_systems
		WHERE id = $1
	`

	var os models.OS
	err := r.db.QueryRow(query, id).Scan(
		&os.ID,
		&os.Name,
		&os.Version,
		&os.EndOfSupport,
		&os.CreatedAt,
		&os.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("operating system with id %d not found", id)
		}
		return nil, fmt.Errorf("failed to get operating system: %w", err)
	}

	return &os, nil
}

// Create creates a new operating system in the database
func (r *OSRepository) Create(req *models.CreateOSRequest) (*models.OS, error) {
	// Parse the end of support date
	endOfSupport, err := time.Parse("2006-01-02", req.EndOfSupport)
	if err != nil {
		return nil, fmt.Errorf("invalid end of support date format: %w", err)
	}

	query := `
		INSERT INTO operating_systems (name, version, end_of_support, created_at, updated_at)
		VALUES ($1, $2, $3, NOW(), NOW())
		RETURNING id, name, version, end_of_support, created_at, updated_at
	`

	var os models.OS
	err = r.db.QueryRow(query, req.Name, req.Version, endOfSupport).Scan(
		&os.ID,
		&os.Name,
		&os.Version,
		&os.EndOfSupport,
		&os.CreatedAt,
		&os.UpdatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to create operating system: %w", err)
	}

	return &os, nil
}

// Update updates an existing operating system in the database
func (r *OSRepository) Update(id int, req *models.UpdateOSRequest) (*models.OS, error) {
	// Build dynamic update query
	setParts := []string{}
	args := []interface{}{}
	argCount := 1

	if req.Name != "" {
		setParts = append(setParts, fmt.Sprintf("name = $%d", argCount))
		args = append(args, req.Name)
		argCount++
	}

	if req.Version != "" {
		setParts = append(setParts, fmt.Sprintf("version = $%d", argCount))
		args = append(args, req.Version)
		argCount++
	}

	if req.EndOfSupport != "" {
		endOfSupport, err := time.Parse("2006-01-02", req.EndOfSupport)
		if err != nil {
			return nil, fmt.Errorf("invalid end of support date format: %w", err)
		}
		setParts = append(setParts, fmt.Sprintf("end_of_support = $%d", argCount))
		args = append(args, endOfSupport)
		argCount++
	}

	if len(setParts) == 0 {
		return r.GetByID(id) // No updates, return existing OS
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
		UPDATE operating_systems
		SET %s
		WHERE id = $%d
		RETURNING id, name, version, end_of_support, created_at, updated_at
	`, setClause, argCount)

	var os models.OS
	err := r.db.QueryRow(query, args...).Scan(
		&os.ID,
		&os.Name,
		&os.Version,
		&os.EndOfSupport,
		&os.CreatedAt,
		&os.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("operating system with id %d not found", id)
		}
		return nil, fmt.Errorf("failed to update operating system: %w", err)
	}

	return &os, nil
}

// Delete removes an operating system from the database
func (r *OSRepository) Delete(id int) error {
	// Check if any servers are using this OS
	var count int
	checkQuery := `SELECT COUNT(*) FROM servers WHERE os_id = $1`
	err := r.db.QueryRow(checkQuery, id).Scan(&count)
	if err != nil {
		return fmt.Errorf("failed to check OS usage: %w", err)
	}
	if count > 0 {
		return fmt.Errorf("cannot delete operating system: %d servers are using it", count)
	}

	query := `DELETE FROM operating_systems WHERE id = $1`

	result, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete operating system: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("operating system with id %d not found", id)
	}

	return nil
}
