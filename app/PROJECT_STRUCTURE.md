# Infra Dashboard - Project Structure

This document provides an overview of the complete project structure for the Infra Dashboard Go API.

## Project Overview

The Infra Dashboard is a REST API built with Go, Gorilla Mux, and PostgreSQL for managing server infrastructure information. It provides CRUD operations for servers with JSON responses.

## Directory Structure

```
app/
├── cmd/
│   └── main.go                    # Application entry point and HTTP server setup
├── internal/
│   ├── config/
│   │   └── config.go             # Configuration management and environment variables
│   ├── database/
│   │   └── database.go           # Database connection, migrations, and repository
│   ├── handlers/
│   │   └── server.go             # HTTP handlers for server CRUD operations
│   └── models/
│       ├── server.go             # Data models and request/response structures
│       └── server_test.go        # Unit tests for server models
├── .air.toml                     # Air configuration for hot reloading
├── .env.example                  # Example environment configuration
├── docker-compose.yml            # Docker Compose setup for development
├── Dockerfile                    # Container configuration
├── go.mod                        # Go module definition
├── go.sum                        # Go module checksums
├── init.sql                      # Database initialization SQL
├── Makefile                      # Build automation and common tasks
├── PROJECT_STRUCTURE.md          # This file
├── README.md                     # Project documentation and setup guide
└── test_api.sh                   # API testing script with curl examples
```

## Core Components

### 1. Application Entry Point (`cmd/main.go`)
- Sets up the HTTP server using Gorilla Mux
- Configures database connection
- Defines API routes and middleware
- Includes CORS and logging middleware

### 2. Configuration (`internal/config/config.go`)
- Environment-based configuration management
- Database and server configuration structures
- Default values and environment variable parsing

### 3. Database Layer (`internal/database/database.go`)
- PostgreSQL connection management
- Server repository with full CRUD operations
- Automatic table creation and indexing
- Connection pooling configuration

### 4. HTTP Handlers (`internal/handlers/server.go`)
- RESTful API endpoints for server management
- JSON request/response handling
- Error handling and HTTP status codes
- Health check endpoint

### 5. Data Models (`internal/models/server.go`)
- Server struct with JSON and database tags
- Request/response DTOs for create and update operations
- Validation structures

## API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/health` | Health check |
| GET | `/api/v1/servers` | Get all servers |
| GET | `/api/v1/servers/{id}` | Get server by ID |
| POST | `/api/v1/servers` | Create new server |
| PUT | `/api/v1/servers/{id}` | Update server |
| DELETE | `/api/v1/servers/{id}` | Delete server |

## Database Schema

### Servers Table
```sql
CREATE TABLE servers (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE,
    os VARCHAR(100) NOT NULL,
    os_version VARCHAR(100) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);
```

## Development Tools

### Docker Support
- `Dockerfile`: Multi-stage build for production deployment
- `docker-compose.yml`: Complete development environment with PostgreSQL and Adminer

### Build Automation
- `Makefile`: Common development tasks (build, run, test, docker operations)
- Hot reloading support with Air (`.air.toml`)

### Testing
- Unit tests for models (`internal/models/server_test.go`)
- API integration test script (`test_api.sh`)

## Environment Configuration

Required environment variables:
- `DB_HOST`: Database host (default: localhost)
- `DB_PORT`: Database port (default: 5432)
- `DB_USER`: Database user (default: postgres)
- `DB_PASSWORD`: Database password (default: postgres)
- `DB_NAME`: Database name (default: infra_dashboard)
- `DB_SSLMODE`: SSL mode (default: disable)
- `SERVER_PORT`: Server port (default: 8080)

## Features

### Core Functionality
- Complete CRUD operations for server management
- JSON API with proper HTTP status codes
- PostgreSQL integration with connection pooling
- Automatic database schema creation

### Production Ready
- Docker containerization
- CORS middleware
- Request logging
- Health check endpoint
- Environment-based configuration
- Database connection management

### Developer Experience
- Hot reloading with Air
- Comprehensive documentation
- API testing scripts
- Docker Compose development environment
- Make targets for common tasks
- Unit tests

## Getting Started

1. **Prerequisites**: Go 1.23+, PostgreSQL 17+
2. **Setup**: `make dev-setup`
3. **Run locally**: `make run`
4. **Run with Docker**: `make docker-compose-up`
5. **Test API**: `./test_api.sh`

## Architecture Decisions

### Project Layout
- Follows Go project layout standards
- `internal/` package prevents external imports
- Clear separation of concerns (handlers, models, database)

### Database Design
- Simple relational model with proper indexing
- Automatic timestamps for audit trail
- Unique constraints for data integrity

### API Design
- RESTful endpoints following HTTP conventions
- JSON-first approach
- Proper error handling and status codes
- Versioned API (`/api/v1/`)

### Development Workflow
- Hot reloading for rapid development
- Docker for consistent environments
- Make targets for common operations
- Comprehensive testing approach