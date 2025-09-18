# Infrastructure Dashboard API

A comprehensive REST API for managing server infrastructure and operating systems with proper relational database design, built with Go, Gorilla Mux, and PostgreSQL.

## Features

- **Server Management**: Complete CRUD operations for server inventory
- **Operating System Management**: Centralized OS lifecycle tracking with end-of-support dates
- **Relational Data Model**: Normalized database design with foreign key relationships
- **Compliance Reporting**: Automated compliance analysis and recommendations
- **End-of-Life Tracking**: Monitor OS support status and plan upgrades
- **JSON API**: RESTful API with comprehensive error handling
- **PostgreSQL Storage**: Robust data persistence with referential integrity
- **Environment Configuration**: Flexible configuration management
- **CORS Support**: Cross-origin resource sharing enabled
- **Request Logging**: Apache-style request logging middleware
- **Health Monitoring**: Service health check endpoints
- **Comprehensive Testing**: Full test suite with utilities

## Quick Start

### Prerequisites

- Go 1.23 or higher
- PostgreSQL 17 or higher
- Docker (optional)

### Using Docker Compose (Recommended)

1. Clone and navigate to the project:
```bash
git clone <repository-url>
cd app
```

2. Start the services:
```bash
docker compose up -d
```

3. The API will be available at `http://localhost:8080`
4. Database admin interface at `http://localhost:8081` (Adminer)

### Manual Setup

1. Set environment variables:
```bash
export DB_HOST=localhost
export DB_PORT=5432
export DB_USER=postgres
export DB_PASSWORD=postgres
export DB_NAME=infra_dashboard
export DB_SSLMODE=disable
export SERVER_PORT=8080
```

2. Create PostgreSQL database:
```sql
CREATE DATABASE infra_dashboard;
```

3. Run the application:
```bash
go run cmd/main.go
```

## API Endpoints

### Health Check
- `GET /health` - Service health status

### Operating Systems
- `GET /api/v1/os` - List all operating systems
- `GET /api/v1/os/{id}` - Get operating system by ID
- `POST /api/v1/os` - Create new operating system
- `PUT /api/v1/os/{id}` - Update operating system
- `DELETE /api/v1/os/{id}` - Delete operating system (if not in use)

### Servers
- `GET /api/v1/servers` - Get all servers with OS details
- `GET /api/v1/servers/{id}` - Get server by ID with OS details
- `POST /api/v1/servers` - Create new server
- `PUT /api/v1/servers/{id}` - Update server
- `DELETE /api/v1/servers/{id}` - Delete server
- `GET /api/v1/servers/compliance` - Generate compliance report

## Data Models

### Operating System
```json
{
  "id": 28,
  "name": "Ubuntu",
  "version": "22.04",
  "end_of_support": "2027-04-01T00:00:00Z",
  "created_at": "2024-01-01T12:00:00Z",
  "updated_at": "2024-01-01T12:00:00Z"
}
```

### Server
```json
{
  "id": 1,
  "name": "web-server-01",
  "os_id": 28,
  "os": {
    "id": 28,
    "name": "Ubuntu",
    "version": "22.04",
    "end_of_support": "2027-04-01T00:00:00Z",
    "created_at": "2024-01-01T12:00:00Z",
    "updated_at": "2024-01-01T12:00:00Z"
  },
  "created_at": "2024-01-01T12:00:00Z",
  "updated_at": "2024-01-01T12:00:00Z"
}
```

### Compliance Report
```json
{
  "total_servers": 150,
  "supported_servers": 130,
  "end_of_life_servers": 5,
  "ending_soon_servers": 15,
  "compliance_score": 82.5,
  "score_description": "Good - Minor compliance issues",
  "recommendations": [
    "CRITICAL: 5 servers need immediate OS updates",
    "WARNING: 15 servers have OS support ending within 6 months"
  ],
  "os_distribution": {
    "Ubuntu 22.04": 75,
    "Ubuntu 20.04": 45,
    "CentOS 7": 20,
    "Debian 12": 10
  }
}
```

## Database Schema

The system uses a normalized relational database design:

- **operating_systems**: Central OS catalog with lifecycle information
- **servers**: Server inventory with foreign key references to OS
- **Referential Integrity**: Foreign key constraints ensure data consistency
- **Unique Constraints**: Prevent duplicate OS versions and server names

### Pre-loaded OS Data
The database includes 71 operating systems across major distributions:
- **Ubuntu**: 19 versions (10.04-24.04)
- **Debian**: 8 versions (4-12)
- **CentOS**: 4 versions (4-7)
- **RedHat**: 3 versions (5-7)
- **FreeBSD**: 8 versions (9.0-10.3)
- **OpenBSD**: 29 versions (5.0-7.6)

## API Usage Examples

### Basic Workflow

1. **List available operating systems:**
```bash
curl http://localhost:8080/api/v1/os | jq
```

2. **Create a server using OS ID:**
```bash
curl -X POST http://localhost:8080/api/v1/servers \
  -H "Content-Type: application/json" \
  -d '{
    "name": "web-server-prod-01",
    "os_id": 28
  }'
```

3. **View servers with embedded OS information:**
```bash
curl http://localhost:8080/api/v1/servers | jq
```

### Operating System Management

**Create new OS version:**
```bash
curl -X POST http://localhost:8080/api/v1/os \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Ubuntu",
    "version": "24.10",
    "end_of_support": "2030-07-01"
  }'
```

**Update OS support date:**
```bash
curl -X PUT http://localhost:8080/api/v1/os/28 \
  -H "Content-Type: application/json" \
  -d '{
    "end_of_support": "2027-06-01"
  }'
```

### Server Operations

**Migrate server to new OS:**
```bash
curl -X PUT http://localhost:8080/api/v1/servers/1 \
  -H "Content-Type: application/json" \
  -d '{
    "os_id": 60
  }'
```

**Generate compliance report:**
```bash
curl http://localhost:8080/api/v1/servers/compliance | jq
```

### Bulk Operations

**Create multiple servers:**
```bash
for i in {1..3}; do
  curl -X POST http://localhost:8080/api/v1/servers \
    -H "Content-Type: application/json" \
    -d "{
      \"name\": \"k8s-node-0${i}\",
      \"os_id\": 28
    }"
  sleep 1
done
```

## Testing

### Automated Testing
```bash
# Run all tests
go test ./...

# Run model tests with verbose output
cd internal/models && go test -v

# Run integration tests using the test script
chmod +x test_api.sh
./test_api.sh
```

### Test Coverage
The project includes comprehensive test coverage:
- Model validation and JSON serialization
- Utility functions for compliance analysis
- Database operations and error handling
- API endpoint functionality

## Project Structure

```
app/
├── cmd/
│   └── main.go                    # Application entry point and routing
├── internal/
│   ├── config/
│   │   └── config.go              # Environment-based configuration
│   ├── database/
│   │   └── database.go            # DB connection, repositories, CRUD operations
│   ├── handlers/
│   │   ├── server.go              # Server HTTP handlers
│   │   └── os.go                  # Operating System HTTP handlers
│   └── models/
│       ├── server.go              # Server data model and requests
│       ├── server_test.go         # Server model tests
│       ├── os.go                  # Operating System data model
│       ├── os_test.go             # OS model tests
│       ├── utils.go               # Utility functions for analysis
│       └── utils_test.go          # Utility function tests
├── .air.toml                      # Live reload configuration
├── .gitignore                     # Git ignore rules
├── Dockerfile                     # Container build configuration
├── docker-compose.yml             # Multi-service orchestration
├── go.mod                         # Go module dependencies
├── go.sum                         # Dependency checksums
├── init.sql                       # Database schema and seed data
├── Makefile                       # Build and development commands
├── test_api.sh                    # API integration test script
├── API_REFERENCE.md               # Comprehensive API documentation
├── SERVER_OS_RELATIONSHIP.md     # Technical relationship documentation
├── USAGE_EXAMPLES.md              # Practical usage examples
├── LOGGING_STANDARDS.md           # Logging guidelines
├── PROJECT_STRUCTURE.md           # Architecture documentation
└── README.md                      # This file
```

## Development

### Adding New Features

1. **Models**: Add data structures to `internal/models/`
2. **Database**: Extend repositories in `internal/database/`
3. **Handlers**: Create HTTP handlers in `internal/handlers/`
4. **Routes**: Register new routes in `cmd/main.go`
5. **Tests**: Add corresponding test files

### Live Reload Development

Use Air for automatic reloading during development:
```bash
# Install Air
go install github.com/cosmtrek/air@latest

# Start with live reload
air
```

### Database Operations

**Reset database:**
```bash
docker compose down -v  # Remove volumes
docker compose up -d    # Recreate with fresh data
```

**Access database:**
```bash
docker exec -it infra_dashboard_db psql -U postgres -d infra_dashboard
```

### Building for Production

```bash
# Build optimized binary
make build

# Build Docker image
make docker-build

# Run production binary
./infra-dashboard
```

## Configuration

### Environment Variables

| Variable | Default | Description |
|----------|---------|-------------|
| `DB_HOST` | `localhost` | Database host |
| `DB_PORT` | `5432` | Database port |
| `DB_USER` | `postgres` | Database user |
| `DB_PASSWORD` | `postgres` | Database password |
| `DB_NAME` | `infra_dashboard` | Database name |
| `DB_SSLMODE` | `disable` | SSL mode for database |
| `SERVER_PORT` | `8080` | API server port |

### Docker Compose Services

- **postgres**: PostgreSQL database with initialization
- **api**: Go application server
- **adminer**: Database administration interface

## Performance & Scalability

- **Indexing**: Optimized database indexes for fast queries
- **Connection Pooling**: Configured connection pool settings
- **JOIN Optimization**: Efficient queries with embedded object loading
- **Foreign Key Constraints**: Referential integrity without performance impact

## Security Considerations

- **Input Validation**: Comprehensive request validation
- **SQL Injection Protection**: Parameterized queries
- **Error Handling**: Secure error messages without information leakage
- **CORS Configuration**: Configurable cross-origin policies

## Compliance Features

- **End-of-Life Tracking**: Automated OS lifecycle monitoring
- **Compliance Scoring**: 0-100 scale with detailed explanations
- **Risk Assessment**: Critical, warning, and informational alerts
- **Upgrade Recommendations**: Automated suggestions for OS updates
- **Reporting**: Comprehensive compliance reports with actionable insights

## Monitoring & Logging

- **Apache-style Logging**: Detailed request/response logging
- **Health Checks**: Service availability monitoring
- **Error Tracking**: Structured error logging
- **Performance Metrics**: Request timing and database performance

## Future Enhancements

- Authentication and authorization
- Rate limiting and throttling
- WebSocket notifications for compliance alerts
- Bulk import/export capabilities
- Advanced filtering and search
- Audit trails and change history
- Integration with configuration management tools
- Prometheus metrics and monitoring
- OpenAPI/Swagger documentation

## Contributing

1. Fork the repository
2. Create a feature branch: `git checkout -b feature-name`
3. Make your changes with tests
4. Run the test suite: `go test ./...`
5. Submit a pull request with clear description

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Support

- **Documentation**: Comprehensive guides in `/docs` directory
- **Examples**: Practical usage examples in `USAGE_EXAMPLES.md`
- **API Reference**: Complete API documentation in `API_REFERENCE.md`
- **Issues**: GitHub issue tracker for bug reports and feature requests

---

**Built with ❤️ using Go, PostgreSQL, and modern software engineering practices.**