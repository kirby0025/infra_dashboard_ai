# Infra Dashboard API

A REST API for managing server infrastructure information built with Go, Gorilla Mux, and PostgreSQL.

## Features

- CRUD operations for server management
- JSON API responses
- PostgreSQL database storage
- Environment-based configuration
- CORS support
- Request logging middleware
- Health check endpoint

## API Endpoints

### Health Check
- `GET /health` - Returns service health status

### Servers
- `GET /api/v1/servers` - Get all servers
- `GET /api/v1/servers/{id}` - Get server by ID
- `POST /api/v1/servers` - Create a new server
- `PUT /api/v1/servers/{id}` - Update server by ID
- `DELETE /api/v1/servers/{id}` - Delete server by ID

## Server Model

```json
{
  "id": 1,
  "name": "web-server-01",
  "os": "Ubuntu",
  "os_version": "22.04 LTS",
  "created_at": "2024-01-01T12:00:00Z",
  "updated_at": "2024-01-01T12:00:00Z"
}
```

## Setup and Installation

### Prerequisites

- Go 1.23 or higher
- PostgreSQL 17 or higher

### Environment Variables

Create a `.env` file or set the following environment variables:

```bash
# Database Configuration
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=infra_dashboard
DB_SSLMODE=disable

# Server Configuration
SERVER_PORT=8080
```

### Database Setup

1. Create the PostgreSQL database:
```sql
CREATE DATABASE infra_dashboard;
```

2. The application will automatically create the required tables on startup.

### Running the Application

1. Clone the repository and navigate to the app directory:
```bash
cd app
```

2. Install dependencies:
```bash
go mod tidy
```

3. Run the application:
```bash
go run cmd/main.go
```

The server will start on `http://localhost:8080`

### Building for Production

Build the binary:
```bash
go build -o infra-dashboard cmd/main.go
```

Run the binary:
```bash
./infra-dashboard
```

## API Usage Examples

### Create a Server
```bash
curl -X POST http://localhost:8080/api/v1/servers \
  -H "Content-Type: application/json" \
  -d '{
    "name": "web-server-01",
    "os": "Ubuntu",
    "os_version": "22.04 LTS"
  }'
```

### Get All Servers
```bash
curl http://localhost:8080/api/v1/servers
```

### Get Server by ID
```bash
curl http://localhost:8080/api/v1/servers/1
```

### Update a Server
```bash
curl -X PUT http://localhost:8080/api/v1/servers/1 \
  -H "Content-Type: application/json" \
  -d '{
    "os_version": "22.04.1 LTS"
  }'
```

### Delete a Server
```bash
curl -X DELETE http://localhost:8080/api/v1/servers/1
```

### Health Check
```bash
curl http://localhost:8080/health
```

## Project Structure

```
app/
├── cmd/
│   └── main.go              # Application entry point
├── internal/
│   ├── config/
│   │   └── config.go        # Configuration management
│   ├── database/
│   │   └── database.go      # Database connection and repository
│   ├── handlers/
│   │   └── server.go        # HTTP handlers
│   └── models/
│       └── server.go        # Data models
├── go.mod                   # Go module definition
└── README.md               # This file
```

## Development

### Adding New Features

1. Add new models to `internal/models/`
2. Extend the database schema in `internal/database/`
3. Create new handlers in `internal/handlers/`
4. Register routes in `cmd/main.go`

### Database Migrations

Currently, the application uses a simple table creation approach. For production use, consider implementing proper database migrations.

## Docker Support

To run with Docker, create a `Dockerfile`:

```dockerfile
FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o infra-dashboard cmd/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/infra-dashboard .
EXPOSE 8080
CMD ["./infra-dashboard"]
```

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests if applicable
5. Submit a pull request

## License

This project is licensed under the MIT License.