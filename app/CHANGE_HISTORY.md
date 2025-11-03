# Server Change History Documentation

## Overview

The Server Change History functionality provides automatic tracking of all changes made to servers in the infrastructure dashboard. Every creation, OS change, and deletion is automatically logged with full details.

## Database Schema

### server_change_history Table

The `server_change_history` table stores all change records with the following structure:

| Column | Type | Description |
|--------|------|-------------|
| id | SERIAL | Primary key |
| server_id | INTEGER | Foreign key to servers table (nullable after deletion) |
| server_name | VARCHAR(255) | Name of the server at time of change |
| change_type | VARCHAR(50) | Type of change: 'created', 'os_changed', 'deleted' |
| old_os_id | INTEGER | OS ID before change (null for creation) |
| new_os_id | INTEGER | OS ID after change (null for deletion) |
| old_os_name | VARCHAR(100) | OS name before change (null for creation) |
| old_os_version | VARCHAR(100) | OS version before change (null for creation) |
| new_os_name | VARCHAR(100) | OS name after change (null for deletion) |
| new_os_version | VARCHAR(100) | OS version after change (null for deletion) |
| changed_at | TIMESTAMP | When the change occurred |

### Automatic Tracking

Changes are tracked automatically using PostgreSQL triggers:

1. **Server Creation** - `log_server_creation_trigger`
   - Fired after INSERT on servers table
   - Records the new server and its initial OS

2. **OS Change** - `log_server_os_change_trigger`
   - Fired after UPDATE on servers table
   - Only logs when os_id actually changes
   - Records both old and new OS details

3. **Server Deletion** - `log_server_deletion_trigger`
   - Fired before DELETE on servers table
   - Records the server's final state before deletion

## API Endpoints

### Get All Change History

```
GET /api/v1/history
```

Retrieves all change history records with optional filtering.

**Query Parameters:**
- `server_id` (optional) - Filter by specific server ID
- `change_type` (optional) - Filter by change type: `created`, `os_changed`, or `deleted`
- `start_date` (optional) - Filter changes from this date (format: YYYY-MM-DD)
- `end_date` (optional) - Filter changes until this date (format: YYYY-MM-DD)
- `limit` (optional) - Maximum number of records to return (default: 100)
- `offset` (optional) - Number of records to skip for pagination (default: 0)

**Example Request:**
```bash
curl "http://localhost:8080/api/v1/history?change_type=os_changed&limit=10"
```

**Example Response:**
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

### Get Change History for Specific Server

```
GET /api/v1/servers/{id}/history
```

Retrieves change history for a specific server.

**Path Parameters:**
- `id` - Server ID

**Query Parameters:**
- `limit` (optional) - Maximum number of records to return (default: 50)

**Example Request:**
```bash
curl "http://localhost:8080/api/v1/servers/3/history"
```

**Example Response:**
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

### Get Specific Change History Record

```
GET /api/v1/history/{id}
```

Retrieves a single change history record by its ID.

**Path Parameters:**
- `id` - Change history record ID

**Example Request:**
```bash
curl "http://localhost:8080/api/v1/history/15"
```

**Example Response:**
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

## Change Types

### 1. Created

Records when a new server is added to the system.

**Characteristics:**
- `change_type`: `"created"`
- `old_os_id`, `old_os_name`, `old_os_version`: `null`
- `new_os_id`, `new_os_name`, `new_os_version`: Set to initial OS

**Example:**
```json
{
  "id": 1,
  "server_id": 5,
  "server_name": "db-server-01",
  "change_type": "created",
  "old_os_id": null,
  "new_os_id": 42,
  "old_os_name": null,
  "old_os_version": null,
  "new_os_name": "Debian",
  "new_os_version": "12",
  "changed_at": "2024-01-10T14:22:00Z"
}
```

### 2. OS Changed

Records when a server's operating system is updated.

**Characteristics:**
- `change_type`: `"os_changed"`
- Both old and new OS fields are populated
- `server_id` remains the same

**Example:**
```json
{
  "id": 8,
  "server_id": 2,
  "server_name": "app-server-03",
  "change_type": "os_changed",
  "old_os_id": 35,
  "new_os_id": 39,
  "old_os_name": "CentOS",
  "old_os_version": "7",
  "new_os_name": "RedHat",
  "new_os_version": "9",
  "changed_at": "2024-01-18T16:45:00Z"
}
```

### 3. Deleted

Records when a server is removed from the system.

**Characteristics:**
- `change_type`: `"deleted"`
- `new_os_id`, `new_os_name`, `new_os_version`: `null`
- `old_os_id`, `old_os_name`, `old_os_version`: Set to final OS state
- `server_id` may become `null` if the server record is hard deleted

**Example:**
```json
{
  "id": 12,
  "server_id": 7,
  "server_name": "legacy-server-01",
  "change_type": "deleted",
  "old_os_id": 15,
  "new_os_id": null,
  "old_os_name": "Ubuntu",
  "old_os_version": "18.04",
  "new_os_name": null,
  "new_os_version": null,
  "changed_at": "2024-01-20T09:15:00Z"
}
```

## Usage Examples

### Audit Trail for Compliance

Get all changes in the last 30 days:
```bash
START_DATE=$(date -d '30 days ago' +%Y-%m-%d)
END_DATE=$(date +%Y-%m-%d)
curl "http://localhost:8080/api/v1/history?start_date=$START_DATE&end_date=$END_DATE"
```

### Track OS Migrations

Get all OS change events:
```bash
curl "http://localhost:8080/api/v1/history?change_type=os_changed"
```

### Server Lifecycle

Get complete history for a specific server:
```bash
curl "http://localhost:8080/api/v1/servers/5/history?limit=100"
```

### Monitor Deletions

Get all deleted servers:
```bash
curl "http://localhost:8080/api/v1/history?change_type=deleted"
```

## Implementation Details

### Database Triggers

The system uses three PostgreSQL trigger functions:

1. **log_server_creation()** - Captures new server details
2. **log_server_os_change()** - Detects and logs OS changes only
3. **log_server_deletion()** - Preserves server state before deletion

These triggers ensure that no change goes unrecorded, even if the application crashes or the database is accessed directly.

### Foreign Key Behavior

The `server_id` foreign key uses `ON DELETE SET NULL` to preserve history even after a server is deleted. This ensures:
- History records are never lost
- Deleted servers can still be audited
- The change log remains complete

### Indexes

The following indexes optimize query performance:
- `idx_change_history_server_id` - Fast lookups by server
- `idx_change_history_change_type` - Fast filtering by change type
- `idx_change_history_changed_at` - Efficient date range queries

## Data Retention

By default, all change history is retained indefinitely. For large deployments, consider implementing a retention policy:

```sql
-- Example: Delete history older than 2 years
DELETE FROM server_change_history 
WHERE changed_at < NOW() - INTERVAL '2 years';
```

## Security Considerations

- Change history is **read-only** via the API
- No endpoints exist to modify or delete history records
- Direct database access is required to alter history
- Consider implementing additional authentication for history endpoints in production

## Testing

Run the change history tests:
```bash
go test ./internal/models/... -v -run TestServerChangeHistory
go test ./internal/models/... -v -run TestChangeHistoryFilter
```

## Future Enhancements

Potential improvements for the change history system:
- Add user tracking (who made the change)
- Include IP address or session information
- Add name change tracking (currently only tracks OS changes)
- Export functionality (CSV, JSON)
- Webhooks for real-time change notifications
- Restore functionality to revert changes