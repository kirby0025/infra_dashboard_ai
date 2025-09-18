#!/bin/bash

# API Test Script for Infra Dashboard
# This script demonstrates all the available API endpoints

set -e

API_BASE="http://localhost:8080"
API_V1="$API_BASE/api/v1"

echo "========================================="
echo "Infra Dashboard API Test Script"
echo "========================================="
echo "API Base URL: $API_BASE"
echo ""

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Function to print colored output
print_step() {
    echo -e "${BLUE}$1${NC}"
}

print_success() {
    echo -e "${GREEN}✓ $1${NC}"
}

print_warning() {
    echo -e "${YELLOW}⚠ $1${NC}"
}

print_error() {
    echo -e "${RED}✗ $1${NC}"
}

# Function to make API calls and pretty print JSON
api_call() {
    local method=$1
    local endpoint=$2
    local data=$3
    local description=$4

    print_step "$description"
    echo "Request: $method $endpoint"
    if [ -n "$data" ]; then
        echo "Data: $data"
    fi
    echo ""

    if [ -n "$data" ]; then
        response=$(curl -s -X $method "$endpoint" \
            -H "Content-Type: application/json" \
            -d "$data" \
            -w "\nHTTP_STATUS:%{http_code}")
    else
        response=$(curl -s -X $method "$endpoint" \
            -w "\nHTTP_STATUS:%{http_code}")
    fi

    # Extract HTTP status
    http_status=$(echo "$response" | grep "HTTP_STATUS:" | cut -d: -f2)
    json_response=$(echo "$response" | sed '/HTTP_STATUS:/d')

    if [ "$http_status" -ge 200 ] && [ "$http_status" -lt 300 ]; then
        print_success "HTTP Status: $http_status"
    else
        print_error "HTTP Status: $http_status"
    fi

    # Pretty print JSON if response is not empty
    if [ -n "$json_response" ] && [ "$json_response" != "null" ]; then
        echo "Response:"
        echo "$json_response" | python3 -m json.tool 2>/dev/null || echo "$json_response"
    fi

    echo ""
    echo "----------------------------------------"
    echo ""

    # Return the response for further processing
    echo "$json_response"
}

# Check if the server is running
print_step "1. Health Check"
health_response=$(api_call "GET" "$API_BASE/health" "" "Checking API health...")

# Get all operating systems first
print_step "2. Get All Operating Systems"
all_os=$(api_call "GET" "$API_V1/os" "" "Retrieving all operating systems...")

# Get all servers (initially should be empty or have sample data)
print_step "3. Get All Servers (Initial State)"
initial_servers=$(api_call "GET" "$API_V1/servers" "" "Retrieving all servers...")

# Create first server (using Ubuntu 22.04 - should be OS ID 28)
print_step "4. Create Server #1"
server1_data='{
    "name": "web-server-01",
    "os_id": 28
}'
server1_response=$(api_call "POST" "$API_V1/servers" "$server1_data" "Creating web-server-01...")

# Extract server ID from response
server1_id=$(echo "$server1_response" | python3 -c "import sys, json; print(json.load(sys.stdin)['id'])" 2>/dev/null || echo "1")

# Create second server (using CentOS 7 - should be OS ID 32)
print_step "5. Create Server #2"
server2_data='{
    "name": "db-server-01",
    "os_id": 32
}'
server2_response=$(api_call "POST" "$API_V1/servers" "$server2_data" "Creating db-server-01...")

# Extract server ID from response
server2_id=$(echo "$server2_response" | python3 -c "import sys, json; print(json.load(sys.stdin)['id'])" 2>/dev/null || echo "2")

# Create third server (using RedHat 7 - should be OS ID 36)
print_step "6. Create Server #3"
server3_data='{
    "name": "app-server-01",
    "os_id": 36
}'
server3_response=$(api_call "POST" "$API_V1/servers" "$server3_data" "Creating app-server-01...")

# Extract server ID from response
server3_id=$(echo "$server3_response" | python3 -c "import sys, json; print(json.load(sys.stdin)['id'])" 2>/dev/null || echo "3")

# Get all servers after creation
print_step "7. Get All Servers (After Creation)"
all_servers=$(api_call "GET" "$API_V1/servers" "" "Retrieving all servers after creation...")

# Get server by ID
print_step "8. Get Server by ID"
single_server=$(api_call "GET" "$API_V1/servers/$server1_id" "" "Retrieving server with ID $server1_id...")

# Update server (change to Ubuntu 20.04 - should be OS ID 27)
print_step "9. Update Server"
update_data='{
    "os_id": 27
}'
updated_server=$(api_call "PUT" "$API_V1/servers/$server1_id" "$update_data" "Updating server $server1_id...")

# Partial update (only name)
print_step "10. Partial Update Server"
partial_update_data='{
    "name": "web-server-01-updated"
}'
partially_updated_server=$(api_call "PUT" "$API_V1/servers/$server2_id" "$partial_update_data" "Partially updating server $server2_id...")

# Get updated server
print_step "11. Get Updated Server"
updated_server_check=$(api_call "GET" "$API_V1/servers/$server1_id" "" "Checking updated server $server1_id...")

# Test error cases
print_step "12. Error Cases"

# Try to get non-existent server
print_warning "Testing error case: Get non-existent server"
error_response=$(api_call "GET" "$API_V1/servers/999" "" "Trying to get server with ID 999 (should not exist)...")

# Try to create server with missing data
print_warning "Testing error case: Create server with missing data"
invalid_data='{
    "name": "incomplete-server"
}'
error_create=$(api_call "POST" "$API_V1/servers" "$invalid_data" "Trying to create server with incomplete data...")

# Try to create server with invalid OS ID
print_warning "Testing error case: Create server with invalid OS ID"
invalid_os_data='{
    "name": "invalid-os-server",
    "os_id": 999
}'
invalid_os_response=$(api_call "POST" "$API_V1/servers" "$invalid_os_data" "Trying to create server with invalid OS ID...")

# Try to create server with duplicate name
print_warning "Testing error case: Create server with duplicate name"
duplicate_data='{
    "name": "web-server-01-updated",
    "os_id": 27
}'
duplicate_response=$(api_call "POST" "$API_V1/servers" "$duplicate_data" "Trying to create server with duplicate name...")

# Test OS endpoints
print_step "13. Operating System Endpoints"

# Get specific OS by ID
specific_os=$(api_call "GET" "$API_V1/os/28" "" "Retrieving Ubuntu 22.04 (OS ID 28)...")

# Create new OS
print_step "14. Create New Operating System"
new_os_data='{
    "name": "TestOS",
    "version": "1.0",
    "end_of_support": "2025-12-31"
}'
new_os_response=$(api_call "POST" "$API_V1/os" "$new_os_data" "Creating new operating system...")

# Test compliance reporting
print_step "15. Compliance Report"
compliance_report=$(api_call "GET" "$API_V1/servers/compliance" "" "Generating compliance report...")

# Get final state
print_step "16. Final State - Get All Servers"
final_servers=$(api_call "GET" "$API_V1/servers" "" "Retrieving final state of all servers...")

# Optional: Clean up (uncomment if you want to delete test data)
print_step "17. Cleanup (Optional)"
print_warning "Cleanup is commented out. Uncomment the lines below to delete test servers."
echo "# Delete servers created during testing"
echo "# curl -X DELETE $API_V1/servers/$server1_id"
echo "# curl -X DELETE $API_V1/servers/$server2_id"
echo "# curl -X DELETE $API_V1/servers/$server3_id"

# Uncomment the lines below to actually perform cleanup
# print_step "Deleting test servers..."
# api_call "DELETE" "$API_V1/servers/$server3_id" "" "Deleting server $server3_id..."
# api_call "DELETE" "$API_V1/servers/$server2_id" "" "Deleting server $server2_id..."
# api_call "DELETE" "$API_V1/servers/$server1_id" "" "Deleting server $server1_id..."

echo "========================================="
echo "API Testing Complete!"
echo "========================================="
echo ""
echo "Summary of endpoints tested:"
echo "✓ GET /health"
echo "✓ GET /api/v1/os"
echo "✓ GET /api/v1/os/{id}"
echo "✓ POST /api/v1/os"
echo "✓ GET /api/v1/servers"
echo "✓ GET /api/v1/servers/{id}"
echo "✓ POST /api/v1/servers"
echo "✓ PUT /api/v1/servers/{id}"
echo "✓ GET /api/v1/servers/compliance"
echo "✓ DELETE /api/v1/servers/{id} (commented out)"
echo ""
echo "To run this script:"
echo "1. Make sure the API server is running (go run cmd/main.go)"
echo "2. Make sure PostgreSQL is running and configured"
echo "3. Run: chmod +x test_api.sh && ./test_api.sh"
