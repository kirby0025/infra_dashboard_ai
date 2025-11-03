# Infrastructure Dashboard API Reference

This document provides comprehensive API reference for the Infrastructure Dashboard, which manages servers and operating systems with proper relational database design.

## Base URL

```
http://localhost:8080
```

## Authentication

Currently, the API does not require authentication. This should be implemented for production use.

## Content Type

All requests and responses use `application/json` content type unless otherwise specified.

## Error Responses

The API uses standard HTTP status codes and returns error responses in the following format:

```json
{
  "error": "Error message describing what went wrong"
}
```

### Common HTTP Status Codes

- `200 OK` - Request successful
- `201 Created` - Resource created successfully
- `204 No Content` - Request successful, no response body
- `400 Bad Request` - Invalid request data
- `404 Not Found` - Resource not found
- `409 Conflict` - Resource conflict (e.g., trying to delete OS in use)
- `500 Internal Server Error` - Server error

---

## Health Check

### GET /health

Check if the API service is running.

**Response:**
```json
{
  "status": "healthy",
  "service": "infra-dashboard"
}
```

---

## Operating Systems

### GET /api/v1/os

List all operating systems.

**Response:**
```json
[
  {
    "id": 1,
    "name": "Ubuntu",
    "version": "22.04",
    "end_of_support": "2027-04-01T00:00:00Z",
    "created_at": "2024-01-01T12:00:00Z",
    "updated_at": "2024-01-01T12:00:00Z"
  },
  {
    "id": 2,
    "name": "Debian",
    "version": "12",
    "end_of_support": "2028-06-30T00:00:00Z",
    "created_at": "2024-01-01T12:00:00Z",
    "updated_at": "2024-01-01T12:00:00Z"
  }
]
```

### GET /api/v1/os/{id}

Get a specific operating system by ID.

**Parameters:**
- `id` (integer, required) - Operating system ID

**Response:**
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

**Error Responses:**
- `404 Not Found` - OS with specified ID not found

### POST /api/v1/os

Create a new operating system.

**Request Body:**
```json
{
  "name": "Ubuntu",
  "version": "24.04",
  "end_of_support": "2029-04-01"
}
```

**Required Fields:**
- `name` (string) - Operating system name
- `version` (string) - Version number
- `end_of_support` (string) - End of support date in YYYY-MM-DD format

**Response:**
```json
{
  "id": 61,
  "name": "Ubuntu",
  "version": "24.04",
  "end_of_support": "2029-04-01T00:00:00Z",
  "created_at": "2024-01-01T12:00:00Z",
  "updated_at": "2024-01-01T12:00:00Z"
}
```

**Error Responses:**
- `400 Bad Request` - Invalid request data or date format
- `409 Conflict` - OS with same name and version already exists

### PUT /api/v1/os/{id}

Update an existing operating system.

**Parameters:**
- `id` (integer, required) - Operating system ID

**Request Body (all fields optional):**
```json
{
  "name": "Ubuntu",
  "version": "24.04.1",
  "end_of_support": "2029-06-01"
}
```

**Response:**
```json
{
  "id": 61,
  "name": "Ubuntu",
  "version": "24.04.1",
  "end_of_support": "2029-06-01T00:00:00Z",
  "created_at": "2024-01-01T12:00:00Z",
  "updated_at": "2024-01-01T15:30:00Z"
}
```

**Error Responses:**
- `400 Bad Request` - Invalid request data or date format
- `404 Not Found` - OS with specified ID not found

### DELETE /api/v1/os/{id}

Delete an operating system.

**Parameters:**
- `id` (integer, required) - Operating system ID

**Response:**
- `204 No Content` - OS deleted successfully

**Error Responses:**
- `404 Not Found` - OS with specified ID not found
- `409 Conflict` - Cannot delete OS because servers are using it

---

## Servers

### GET /api/v1/servers

List all servers with embedded OS information.

**Response:**
```json
[
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
]
```

### GET /api/v1/servers/{id}

Get a specific server by ID with embedded OS information.

**Parameters:**
- `id` (integer, required) - Server ID

**Response:**
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

**Error Responses:**
- `404 Not Found` - Server with specified ID not found

### POST /api/v1/servers

Create a new server.

**Request Body:**
```json
{
  "name": "web-server-02",
  "os_id": 28
}
```

**Required Fields:**
- `name` (string) - Server name (must be unique)
- `os_id` (integer) - Operating system ID (must exist)

**Response:**
```json
{
  "id": 2,
  "name": "web-server-02",
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

**Error Responses:**
- `400 Bad Request` - Invalid request data or OS ID doesn't exist
- `409 Conflict` - Server with same name already exists

### PUT /api/v1/servers/{id}

Update an existing server.

**Parameters:**
- `id` (integer, required) - Server ID

**Request Body (all fields optional):**
```json
{
  "name": "web-server-02-updated",
  "os_id": 27
}
```

**Response:**
```json
{
  "id": 2,
  "name": "web-server-02-updated",
  "os_id": 27,
  "os": {
    "id": 27,
    "name": "Ubuntu",
    "version": "20.04",
    "end_of_support": "2025-04-01T00:00:00Z",
    "created_at": "2024-01-01T12:00:00Z",
    "updated_at": "2024-01-01T12:00:00Z"
  },
  "created_at": "2024-01-01T12:00:00Z",
  "updated_at": "2024-01-01T15:30:00Z"
}
```

**Error Responses:**
- `400 Bad Request` - Invalid request data or OS ID doesn't exist
- `404 Not Found` - Server with specified ID not found
- `409 Conflict` - Server name already exists (if updating name)

### DELETE /api/v1/servers/{id}

Delete a server.

**Parameters:**
- `id` (integer, required) - Server ID

**Response:**
- `204 No Content` - Server deleted successfully

**Error Responses:**
- `404 Not Found` - Server with specified ID not found

### GET /api/v1/servers/compliance

Generate a comprehensive compliance report for all servers.

**Response:**
```json
{
  "total_servers": 5,
  "supported_servers": 3,
  "end_of_life_servers": 1,
  "ending_soon_servers": 1,
  "os_distribution": {
    "Ubuntu 20.04": 2,
    "Ubuntu 22.04": 1,
    "CentOS 7": 1,
    "Debian 11": 1
  },
  "os_family_distribution": {
    "Ubuntu": 3,
    "CentOS": 1,
    "Debian": 1
  },
  "end_of_life_list": [
    {
      "id": 1,
      "name": "legacy-server",
      "os_id": 15,
      "os": {
        "id": 15,
        "name": "CentOS",
        "version": "6",
        "end_of_support": "2020-11-30T00:00:00Z",
        "created_at": "2024-01-01T12:00:00Z",
        "updated_at": "2024-01-01T12:00:00Z"
      },
      "created_at": "2024-01-01T12:00:00Z",
      "updated_at": "2024-01-01T12:00:00Z"
    }
  ],
  "ending_soon_list": [
    {
      "id": 2,
      "name": "staging-server",
      "os_id": 27,
      "os": {
        "id": 27,
        "name": "Ubuntu",
        "version": "20.04",
        "end_of_support": "2025-04-01T00:00:00Z",
        "created_at": "2024-01-01T12:00:00Z",
        "updated_at": "2024-01-01T12:00:00Z"
      },
      "created_at": "2024-01-01T12:00:00Z",
      "updated_at": "2024-01-01T12:00:00Z"
    }
  ],
  "generated_at": "2024-01-01T15:30:00Z",
  "compliance_score": 75.0,
  "score_description": "Good - Minor compliance issues that should be addressed",
  "recommendations": [
    "CRITICAL: 1 servers are running end-of-life operating systems and need immediate updates",
    "WARNING: 1 servers are running operating systems that will reach end-of-life within 6 months"
  ]
}
```

**Compliance Score Ranges:**
- `90-100`: Excellent - Infrastructure is well maintained and compliant
- `75-89`: Good - Minor compliance issues that should be addressed
- `50-74`: Fair - Several compliance issues requiring attention
- `25-49`: Poor - Significant compliance issues need immediate action
- `0-24`: Critical - Infrastructure has serious compliance problems

---

## Data Models

### Operating System

```json
{
  "id": 1,
  "name": "Ubuntu",
  "version": "22.04",
  "end_of_support": "2027-04-01T00:00:00Z",
  "created_at": "2024-01-01T12:00:00Z",
  "updated_at": "2024-01-01T12:00:00Z"
}
```

**Fields:**
- `id` (integer) - Unique identifier
- `name` (string) - Operating system name
- `version` (string) - Version number
- `end_of_support` (string, ISO 8601) - End of support date
- `created_at` (string, ISO 8601) - Creation timestamp
- `updated_at` (string, ISO 8601) - Last update timestamp

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

**Fields:**
- `id` (integer) - Unique identifier
- `name` (string) - Server name (must be unique)
- `os_id` (integer) - Foreign key reference to operating system
- `os` (object, optional) - Embedded operating system object
- `created_at` (string, ISO 8601) - Creation timestamp
- `updated_at` (string, ISO 8601) - Last update timestamp

---

## Example Workflows

### Creating a Server

1. **List available operating systems:**
   ```bash
   curl -X GET http://localhost:8080/api/v1/os
   ```

2. **Create a server with chosen OS:**
   ```bash
   curl -X POST http://localhost:8080/api/v1/servers \
     -H "Content-Type: application/json" \
     -d '{"name": "new-server", "os_id": 28}'
   ```

### Checking Compliance

1. **Generate compliance report:**
   ```bash
   curl -X GET http://localhost:8080/api/v1/servers/compliance
   ```

2. **Review end-of-life servers and plan upgrades based on recommendations**

### Migrating Server OS

1. **Find target OS ID:**
   ```bash
   curl -X GET http://localhost:8080/api/v1/os | jq '.[] | select(.name == "Ubuntu" and .version == "22.04")'
   ```

2. **Update server:**
   ```bash
   curl -X PUT http://localhost:8080/api/v1/servers/1 \
     -H "Content-Type: application/json" \
     -d '{"os_id": 28}'
   ```

### Adding New OS Version

```bash
curl -X POST http://localhost:8080/api/v1/os \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Ubuntu",
    "version": "24.04",
    "end_of_support": "2029-04-01"
  }'
```

---

## Rate Limiting

Currently, no rate limiting is implemented. For production use, consider implementing rate limiting based on your requirements.

## Pagination

The current API does not implement pagination. For large datasets, pagination should be added to the list endpoints.

---

## Change History Endpoints

### GET /api/v1/history

Retrieve all server change history records with optional filtering.

**Query Parameters:**
- `server_id` (optional) - Filter by specific server ID
- `change_type` (optional) - Filter by change type: `created`, `os_changed`, or `deleted`
- `start_date` (optional) - Filter changes from this date (format: YYYY-MM-DD)
- `end_date` (optional) - Filter changes until this date (format: YYYY-MM-DD)
- `limit` (optional, default: 100) - Maximum number of records to return
- `offset` (optional, default: 0) - Number of records to skip for pagination

**Example Request:**
```bash
curl "http://localhost:8080/api/v1/history?change_type=os_changed&limit=10"
```

**Response:**
```json
[
  {
    "id": 15,
    "server_id": 3,
    "server_name": "web-server-01",
    "change_type": "os_changed",
    "old_os_id": 28,
    "new_os_id": 29,
    "old_os_name": "Ubuntu",
    "old_os_version": "20.04",
    "new_os_name": "Ubuntu",
    "new_os_version": "22.04",
    "changed_at": "2024-01-15T10:30:00Z"
  }
]
```

---

### GET /api/v1/servers/{id}/history

Retrieve change history for a specific server.

**Path Parameters:**
- `id` - Server ID (required)

**Query Parameters:**
- `limit` (optional, default: 50) - Maximum number of records to return

**Example Request:**
```bash
curl "http://localhost:8080/api/v1/servers/3/history"
```

**Response:**
```json
[
  {
    "id": 15,
    "server_id": 3,
    "server_name": "web-server-01",
    "change_type": "os_changed",
    "old_os_id": 28,
    "new_os_id": 29,
    "old_os_name": "Ubuntu",
    "old_os_version": "20.04",
    "new_os_name": "Ubuntu",
    "new_os_version": "22.04",
    "changed_at": "2024-01-15T10:30:00Z"
  },
  {
    "id": 1,
    "server_id": 3,
    "server_name": "web-server-01",
    "change_type": "created",
    "old_os_id": null,
    "new_os_id": 28,
    "old_os_name": null,
    "old_os_version": null,
    "new_os_name": "Ubuntu",
    "new_os_version": "20.04",
    "changed_at": "2024-01-01T08:00:00Z"
  }
]
```

---

### GET /api/v1/history/{id}

Retrieve a specific change history record by its ID.

**Path Parameters:**
- `id` - Change history record ID (required)

**Example Request:**
```bash
curl "http://localhost:8080/api/v1/history/15"
```

**Response:**
```json
{
  "id": 15,
  "server_id": 3,
  "server_name": "web-server-01",
  "change_type": "os_changed",
  "old_os_id": 28,
  "new_os_id": 29,
  "old_os_name": "Ubuntu",
  "old_os_version": "20.04",
  "new_os_name": "Ubuntu",
  "new_os_version": "22.04",
  "changed_at": "2024-01-15T10:30:00Z"
}
```

**Error Responses:**
- `400 Bad Request` - Invalid history record ID
- `500 Internal Server Error` - Failed to retrieve record

---

## Change History Types

The system automatically tracks three types of changes:

### 1. Created
Recorded when a new server is added.
- `old_os_id`, `old_os_name`, `old_os_version`: `null`
- `new_os_id`, `new_os_name`, `new_os_version`: Set to initial OS

### 2. OS Changed
Recorded when a server's operating system is updated.
- Both old and new OS fields are populated

### 3. Deleted
Recorded when a server is removed.
- `new_os_id`, `new_os_name`, `new_os_version`: `null`
- `old_os_id`, `old_os_name`, `old_os_version`: Set to final OS state

**Note:** Change history is tracked automatically via database triggers. No manual API calls are needed to create history records.

---

## Future Enhancements

- Authentication and authorization
- Rate limiting
- Pagination for large result sets
- Filtering and searching capabilities
- Webhook notifications for compliance issues
- Bulk operations
- API versioning
- User tracking in change history (who made the change)
- Webhook notifications for real-time change alerts
- OpenAPI/Swagger documentation