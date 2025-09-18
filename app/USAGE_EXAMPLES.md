# Usage Examples: Server-OS Relationship

This document provides practical examples of how to use the updated infrastructure dashboard with the new OS-Server relationship model.

## Quick Start Guide

### 1. Starting the Application

```bash
# Start PostgreSQL (if using Docker)
docker compose up -d postgres

# Start the API server
go run cmd/main.go
```

The server will start on `http://localhost:8080` by default.

### 2. Basic Workflow

#### Step 1: Explore Available Operating Systems
```bash
# List all available operating systems
curl -X GET http://localhost:8080/api/v1/os | jq

# Get specific OS by ID
curl -X GET http://localhost:8080/api/v1/os/28 | jq
```

#### Step 2: Create Servers Using OS IDs
```bash
# Create a web server with Ubuntu 22.04 (ID: 28)
curl -X POST http://localhost:8080/api/v1/servers \
  -H "Content-Type: application/json" \
  -d '{
    "name": "web-server-prod-01",
    "os_id": 28
  }' | jq

# Create a database server with CentOS 7 (ID: 32)
curl -X POST http://localhost:8080/api/v1/servers \
  -H "Content-Type: application/json" \
  -d '{
    "name": "db-server-prod-01",
    "os_id": 32
  }' | jq
```

#### Step 3: View Servers with Full OS Information
```bash
# List all servers (includes embedded OS details)
curl -X GET http://localhost:8080/api/v1/servers | jq

# Get specific server
curl -X GET http://localhost:8080/api/v1/servers/1 | jq
```

## Common Use Cases

### Finding OS End-of-Life Information

```bash
# Find all Ubuntu versions and their support dates
curl -X GET http://localhost:8080/api/v1/os | jq '.[] | select(.name == "Ubuntu") | {version, end_of_support}'

# Find operating systems ending support soon (example date filter in application logic)
curl -X GET http://localhost:8080/api/v1/os | jq '.[] | select(.end_of_support < "2025-01-01") | {name, version, end_of_support}'
```

### Server Management Workflows

#### Migrating Server to New OS Version
```bash
# Current server on Ubuntu 20.04 (ID: 27), migrate to 22.04 (ID: 28)
curl -X PUT http://localhost:8080/api/v1/servers/1 \
  -H "Content-Type: application/json" \
  -d '{
    "os_id": 28
  }' | jq
```

#### Bulk Server Creation
```bash
# Create multiple servers for a cluster
for i in {1..3}; do
  curl -X POST http://localhost:8080/api/v1/servers \
    -H "Content-Type: application/json" \
    -d "{
      \"name\": \"k8s-node-0${i}\",
      \"os_id\": 28
    }" | jq '.id'
  sleep 1
done
```

### OS Management

#### Adding New Operating System
```bash
# Add a new OS version when it's released
curl -X POST http://localhost:8080/api/v1/os \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Ubuntu",
    "version": "24.10",
    "end_of_support": "2030-07-01"
  }' | jq
```

#### Updating OS Support Dates
```bash
# Update end-of-support date if it changes
curl -X PUT http://localhost:8080/api/v1/os/28 \
  -H "Content-Type: application/json" \
  -d '{
    "end_of_support": "2027-06-01"
  }' | jq
```

## Advanced Examples

### Finding Servers by OS Characteristics

Since servers now include full OS objects, you can filter and analyze:

```bash
# Get all servers and filter by OS family (client-side filtering)
curl -X GET http://localhost:8080/api/v1/servers | jq '.[] | select(.os.name == "Ubuntu")'

# Find servers running end-of-life operating systems
curl -X GET http://localhost:8080/api/v1/servers | jq '.[] | select(.os.end_of_support < now)'
```

### Compliance Reporting

```bash
# Generate a compliance report
curl -X GET http://localhost:8080/api/v1/servers | jq -r '
  "Server Compliance Report",
  "========================",
  "",
  (.[] | 
    "Server: " + .name + 
    " | OS: " + .os.name + " " + .os.version + 
    " | Support until: " + .os.end_of_support
  )'
```

### Error Handling Examples

#### Invalid OS ID
```bash
# This will fail with a 400 error
curl -X POST http://localhost:8080/api/v1/servers \
  -H "Content-Type: application/json" \
  -d '{
    "name": "invalid-server",
    "os_id": 999
  }'
# Response: {"error": "operating system with id 999 does not exist"}
```

#### Attempting to Delete OS in Use
```bash
# This will fail with a 409 error if servers are using this OS
curl -X DELETE http://localhost:8080/api/v1/os/28
# Response: {"error": "cannot delete operating system: 2 servers are using it"}
```

## JavaScript/Frontend Examples

### Fetching OS Options for a Form

```javascript
// Fetch available OS options for server creation form
async function loadOSOptions() {
  const response = await fetch('/api/v1/os');
  const operatingSystems = await response.json();
  
  // Group by OS name for better UX
  const grouped = operatingSystems.reduce((acc, os) => {
    if (!acc[os.name]) acc[os.name] = [];
    acc[os.name].push(os);
    return acc;
  }, {});
  
  return grouped;
}

// Usage in a form
const osGroups = await loadOSOptions();
// Creates: { "Ubuntu": [...], "Debian": [...], "CentOS": [...] }
```

### Creating a Server with Validation

```javascript
async function createServer(name, osId) {
  try {
    const response = await fetch('/api/v1/servers', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ name, os_id: osId }),
    });
    
    if (!response.ok) {
      const error = await response.json();
      throw new Error(error.message || 'Failed to create server');
    }
    
    const server = await response.json();
    console.log('Server created:', server);
    return server;
  } catch (error) {
    console.error('Error creating server:', error);
    throw error;
  }
}
```

### End-of-Life Dashboard

```javascript
async function getEndOfLifeReport() {
  const [servers, operatingSystems] = await Promise.all([
    fetch('/api/v1/servers').then(r => r.json()),
    fetch('/api/v1/os').then(r => r.json())
  ]);
  
  const now = new Date();
  const oneYear = new Date();
  oneYear.setFullYear(now.getFullYear() + 1);
  
  return {
    endOfLife: servers.filter(s => new Date(s.os.end_of_support) < now),
    endingSoon: servers.filter(s => {
      const eol = new Date(s.os.end_of_support);
      return eol > now && eol < oneYear;
    }),
    supported: servers.filter(s => new Date(s.os.end_of_support) >= oneYear)
  };
}
```

## Python Examples

### Bulk Data Analysis

```python
import requests
import json
from datetime import datetime, timedelta

API_BASE = 'http://localhost:8080/api/v1'

def analyze_infrastructure():
    # Fetch all data
    servers = requests.get(f'{API_BASE}/servers').json()
    
    # Analysis
    os_distribution = {}
    eol_servers = []
    
    for server in servers:
        os_name = server['os']['name']
        os_distribution[os_name] = os_distribution.get(os_name, 0) + 1
        
        eol_date = datetime.fromisoformat(server['os']['end_of_support'].replace('Z', '+00:00'))
        if eol_date < datetime.now(eol_date.tzinfo):
            eol_servers.append(server)
    
    return {
        'total_servers': len(servers),
        'os_distribution': os_distribution,
        'end_of_life_servers': eol_servers,
        'eol_count': len(eol_servers)
    }

# Usage
report = analyze_infrastructure()
print(json.dumps(report, indent=2))
```

### Automated OS Updates

```python
def plan_os_migrations():
    servers = requests.get(f'{API_BASE}/servers').json()
    os_list = requests.get(f'{API_BASE}/os').json()
    
    # Find latest versions for each OS family
    latest_versions = {}
    for os in os_list:
        if os['name'] not in latest_versions:
            latest_versions[os['name']] = os
        elif os['version'] > latest_versions[os['name']]['version']:
            latest_versions[os['name']] = os
    
    # Find servers that can be upgraded
    upgrade_plan = []
    for server in servers:
        current_os = server['os']
        latest = latest_versions.get(current_os['name'])
        
        if latest and latest['version'] > current_os['version']:
            upgrade_plan.append({
                'server_id': server['id'],
                'server_name': server['name'],
                'current_version': current_os['version'],
                'target_version': latest['version'],
                'target_os_id': latest['id']
            })
    
    return upgrade_plan
```

## Docker Compose Example

Complete setup with database:

```yaml
version: "3.8"

services:
  postgres:
    image: postgres:17-alpine
    environment:
      POSTGRES_USER: postgres
      POSTGRES_PASSWORD: postgres
      POSTGRES_DB: infra_dashboard
    ports:
      - "5432:5432"
    volumes:
      - postgres_data:/var/lib/postgresql/data
      - ./init.sql:/docker-entrypoint-initdb.d/init.sql

  api:
    build: .
    environment:
      DB_HOST: postgres
      DB_PORT: 5432
      DB_USER: postgres
      DB_PASSWORD: postgres
      DB_NAME: infra_dashboard
      DB_SSLMODE: disable
      SERVER_PORT: 8080
    ports:
      - "8080:8080"
    depends_on:
      - postgres

volumes:
  postgres_data:
```

## Best Practices

### 1. OS ID Management
- Always fetch current OS list before creating servers
- Cache OS list in frontend applications with appropriate TTL
- Validate OS IDs on the client side for better UX

### 2. Error Handling
- Always check for OS existence errors (400)
- Handle constraint violations gracefully (409)
- Provide meaningful error messages to users

### 3. Data Consistency
- Use transactions when creating multiple related records
- Implement proper retry logic for network requests
- Validate data integrity after bulk operations

### 4. Performance
- Use the embedded OS objects in server responses to avoid extra API calls
- Implement pagination for large server lists
- Consider caching frequently accessed OS data

This new architecture provides a robust foundation for infrastructure management with proper data normalization and comprehensive OS lifecycle tracking.