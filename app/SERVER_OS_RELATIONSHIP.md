# Server-OS Relationship Documentation

This document explains the changes made to integrate the Operating System (OS) model with the Server model, creating a proper relational database structure.

## Overview

The infrastructure dashboard has been updated to use a normalized database design where servers reference operating systems through a foreign key relationship instead of storing OS information as string fields.

## Changes Made

### 1. Database Schema Changes

#### Before (Old Schema)
```sql
CREATE TABLE servers (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE,
    os VARCHAR(100) NOT NULL,           -- String field
    os_version VARCHAR(100) NOT NULL,   -- String field
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);
```

#### After (New Schema)
```sql
CREATE TABLE operating_systems (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    version VARCHAR(100) NOT NULL,
    end_of_support DATE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    UNIQUE(name, version)
);

CREATE TABLE servers (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE,
    os_id INTEGER NOT NULL,             -- Foreign key reference
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    FOREIGN KEY (os_id) REFERENCES operating_systems(id)
);
```

### 2. Model Changes

#### Server Model Updates
```go
// Before
type Server struct {
    ID        int       `json:"id" db:"id"`
    Name      string    `json:"name" db:"name"`
    OS        string    `json:"os" db:"os"`
    OSVersion string    `json:"os_version" db:"os_version"`
    CreatedAt time.Time `json:"created_at" db:"created_at"`
    UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// After
type Server struct {
    ID        int       `json:"id" db:"id"`
    Name      string    `json:"name" db:"name"`
    OSID      int       `json:"os_id" db:"os_id"`
    OS        *OS       `json:"os,omitempty" db:"-"`
    CreatedAt time.Time `json:"created_at" db:"created_at"`
    UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
```

#### Request/Response Changes
```go
// Before
type CreateServerRequest struct {
    Name      string `json:"name" validate:"required"`
    OS        string `json:"os" validate:"required"`
    OSVersion string `json:"os_version" validate:"required"`
}

// After
type CreateServerRequest struct {
    Name string `json:"name" validate:"required"`
    OSID int    `json:"os_id" validate:"required"`
}
```

### 3. Database Operations

The database layer now:
- Performs JOINs to fetch OS details with server information
- Validates OS existence before creating/updating servers
- Prevents deletion of OS records that are referenced by servers
- Returns full OS objects embedded in server responses

### 4. API Endpoints

#### New OS Management Endpoints
- `GET /api/v1/os` - List all operating systems
- `GET /api/v1/os/{id}` - Get specific operating system
- `POST /api/v1/os` - Create new operating system
- `PUT /api/v1/os/{id}` - Update operating system
- `DELETE /api/v1/os/{id}` - Delete operating system (if not in use)

#### Updated Server Endpoints
Server creation and updates now require `os_id` instead of `os` and `os_version` strings.

## Benefits

### 1. Data Normalization
- Eliminates duplicate OS information
- Ensures consistency across servers using the same OS
- Reduces storage requirements

### 2. Data Integrity
- Foreign key constraints prevent invalid OS references
- Centralized OS management ensures accuracy
- End-of-support tracking for compliance monitoring

### 3. Enhanced Functionality
- Easy filtering of servers by OS characteristics
- Centralized OS lifecycle management
- Support for OS end-of-life tracking

### 4. Better API Design
- Cleaner separation of concerns
- Standardized OS data across all servers
- Simplified client-side OS selection

## Migration Guide

### For API Clients

#### Before (Creating a Server)
```json
POST /api/v1/servers
{
    "name": "web-server-01",
    "os": "Ubuntu",
    "os_version": "22.04"
}
```

#### After (Creating a Server)
```json
// First, get available OS options
GET /api/v1/os

// Then create server with OS ID
POST /api/v1/servers
{
    "name": "web-server-01",
    "os_id": 28
}
```

#### Response Structure
```json
// Server response now includes full OS object
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

### For Database Migrations

If migrating existing data:

1. Populate the `operating_systems` table with unique OS/version combinations
2. Add the `os_id` column to the `servers` table
3. Update `servers.os_id` based on existing `os` and `os_version` values
4. Drop the old `os` and `os_version` columns
5. Add the foreign key constraint

## Sample Data

The system includes comprehensive OS data:
- **Debian**: versions 4-12 (2010-2028 support)
- **Ubuntu**: versions 10.04-24.04 (2012-2029 support)
- **CentOS**: versions 4-7 (2012-2024 support)
- **RedHat**: versions 5-7 (2017-2024 support)
- **FreeBSD**: versions 9.0-10.3 (2013-2018 support)
- **OpenBSD**: versions 5.0-7.6 (2012-2025 support)

## Error Handling

### Common Error Scenarios
1. **Invalid OS ID**: Returns 400 with "operating system with id X does not exist"
2. **OS in use**: Returns 409 when trying to delete OS referenced by servers
3. **Missing required fields**: Returns 400 with validation error
4. **Invalid date format**: Returns 400 with format requirements

## Testing

Run the updated test script:
```bash
chmod +x test_api.sh
./test_api.sh
```

The test script demonstrates:
- OS listing and retrieval
- Server creation with OS ID references
- Error handling for invalid OS IDs
- Full CRUD operations on both servers and operating systems

## Performance Considerations

- Indexes added on `servers.os_id` for efficient joins
- OS data is cached through embedded objects in server responses
- Unique constraints on `(name, version)` prevent duplicate OS entries
- Foreign key constraints ensure referential integrity

## Future Enhancements

Potential improvements:
1. OS vulnerability tracking
2. Automatic end-of-life notifications
3. OS update recommendations
4. Compliance reporting based on OS support status
5. OS family grouping (e.g., Red Hat family, Debian family)