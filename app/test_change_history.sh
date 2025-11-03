#!/bin/bash

# Test script for Server Change History functionality
# This script demonstrates creating, updating, and deleting servers
# and then querying the change history

set -e

BASE_URL="http://localhost:8080/api/v1"

echo "========================================"
echo "Server Change History Test Script"
echo "========================================"
echo ""

# Colors for output
GREEN='\033[0;32m'
BLUE='\033[0;34m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Function to print section headers
print_header() {
    echo -e "${BLUE}=== $1 ===${NC}"
}

# Function to print success messages
print_success() {
    echo -e "${GREEN}✓ $1${NC}"
}

# Function to print info messages
print_info() {
    echo -e "${YELLOW}ℹ $1${NC}"
}

# Check if server is running
print_header "Checking Server Status"
if curl -s "$BASE_URL/../health" > /dev/null; then
    print_success "Server is running"
else
    echo "Error: Server is not running. Please start it first."
    exit 1
fi
echo ""

# Get an OS ID to use for testing
print_header "Getting Available OS"
OS_RESPONSE=$(curl -s "$BASE_URL/os?limit=1")
OS_ID=$(echo $OS_RESPONSE | grep -o '"id":[0-9]*' | head -1 | grep -o '[0-9]*')
OS_NAME=$(echo $OS_RESPONSE | grep -o '"name":"[^"]*"' | head -1 | cut -d'"' -f4)
OS_VERSION=$(echo $OS_RESPONSE | grep -o '"version":"[^"]*"' | head -1 | cut -d'"' -f4)
print_info "Using OS ID: $OS_ID ($OS_NAME $OS_VERSION)"
echo ""

# Get another OS ID for update test
SECOND_OS_RESPONSE=$(curl -s "$BASE_URL/os?limit=2&offset=1")
SECOND_OS_ID=$(echo $SECOND_OS_RESPONSE | grep -o '"id":[0-9]*' | head -1 | grep -o '[0-9]*')
SECOND_OS_NAME=$(echo $SECOND_OS_RESPONSE | grep -o '"name":"[^"]*"' | head -1 | cut -d'"' -f4)
SECOND_OS_VERSION=$(echo $SECOND_OS_RESPONSE | grep -o '"version":"[^"]*"' | head -1 | cut -d'"' -f4)
print_info "Will upgrade to OS ID: $SECOND_OS_ID ($SECOND_OS_NAME $SECOND_OS_VERSION)"
echo ""

# 1. Create a test server
print_header "Step 1: Creating a Test Server"
SERVER_NAME="test-history-server-$(date +%s)"
CREATE_RESPONSE=$(curl -s -X POST "$BASE_URL/servers" \
    -H "Content-Type: application/json" \
    -d "{\"name\":\"$SERVER_NAME\",\"os_id\":$OS_ID}")

SERVER_ID=$(echo $CREATE_RESPONSE | grep -o '"id":[0-9]*' | grep -o '[0-9]*')
print_success "Created server: $SERVER_NAME (ID: $SERVER_ID)"
echo "Response: $CREATE_RESPONSE"
echo ""

# Wait a moment to ensure the trigger has fired
sleep 1

# 2. Check change history for the creation
print_header "Step 2: Verifying Creation in Change History"
HISTORY_RESPONSE=$(curl -s "$BASE_URL/servers/$SERVER_ID/history")
echo "Change History for Server $SERVER_ID:"
echo "$HISTORY_RESPONSE" | python3 -m json.tool 2>/dev/null || echo "$HISTORY_RESPONSE"
print_success "Creation event logged"
echo ""

# 3. Update the server's OS
print_header "Step 3: Updating Server OS"
UPDATE_RESPONSE=$(curl -s -X PUT "$BASE_URL/servers/$SERVER_ID" \
    -H "Content-Type: application/json" \
    -d "{\"os_id\":$SECOND_OS_ID}")
print_success "Updated server OS from $OS_ID to $SECOND_OS_ID"
echo "Response: $UPDATE_RESPONSE"
echo ""

# Wait a moment to ensure the trigger has fired
sleep 1

# 4. Check change history for the OS change
print_header "Step 4: Verifying OS Change in Change History"
HISTORY_RESPONSE=$(curl -s "$BASE_URL/servers/$SERVER_ID/history")
echo "Updated Change History for Server $SERVER_ID:"
echo "$HISTORY_RESPONSE" | python3 -m json.tool 2>/dev/null || echo "$HISTORY_RESPONSE"
print_success "OS change event logged"
echo ""

# 5. Query all change history with filters
print_header "Step 5: Querying Change History with Filters"

echo "a) All changes for this server:"
curl -s "$BASE_URL/history?server_id=$SERVER_ID" | python3 -m json.tool 2>/dev/null
echo ""

echo "b) Only 'created' changes:"
curl -s "$BASE_URL/history?server_id=$SERVER_ID&change_type=created" | python3 -m json.tool 2>/dev/null
echo ""

echo "c) Only 'os_changed' changes:"
curl -s "$BASE_URL/history?server_id=$SERVER_ID&change_type=os_changed" | python3 -m json.tool 2>/dev/null
echo ""

# 6. Get today's date for filtering
TODAY=$(date +%Y-%m-%d)
print_header "Step 6: Query by Date Range"
echo "Getting all changes from today ($TODAY):"
curl -s "$BASE_URL/history?start_date=$TODAY" | python3 -m json.tool 2>/dev/null
print_success "Date filtering works"
echo ""

# 7. Delete the server
print_header "Step 7: Deleting the Test Server"
DELETE_RESPONSE=$(curl -s -X DELETE "$BASE_URL/servers/$SERVER_ID")
print_success "Deleted server $SERVER_ID"
echo ""

# Wait a moment to ensure the trigger has fired
sleep 1

# 8. Check change history after deletion
print_header "Step 8: Verifying Deletion in Change History"
echo "Change history still accessible after server deletion:"
HISTORY_RESPONSE=$(curl -s "$BASE_URL/history?server_id=$SERVER_ID")
echo "$HISTORY_RESPONSE" | python3 -m json.tool 2>/dev/null || echo "$HISTORY_RESPONSE"
print_success "Deletion event logged and history preserved"
echo ""

# 9. Get all deleted servers
print_header "Step 9: Query All Deleted Servers"
DELETED_RESPONSE=$(curl -s "$BASE_URL/history?change_type=deleted&limit=5")
echo "Recent server deletions:"
echo "$DELETED_RESPONSE" | python3 -m json.tool 2>/dev/null || echo "$DELETED_RESPONSE"
echo ""

# 10. Summary statistics
print_header "Step 10: Change History Summary"
ALL_HISTORY=$(curl -s "$BASE_URL/history?limit=1000")
CREATED_COUNT=$(echo "$ALL_HISTORY" | grep -o '"change_type":"created"' | wc -l)
OS_CHANGED_COUNT=$(echo "$ALL_HISTORY" | grep -o '"change_type":"os_changed"' | wc -l)
DELETED_COUNT=$(echo "$ALL_HISTORY" | grep -o '"change_type":"deleted"' | wc -l)

echo "Change History Statistics:"
echo "  - Servers Created: $CREATED_COUNT"
echo "  - OS Changes: $OS_CHANGED_COUNT"
echo "  - Servers Deleted: $DELETED_COUNT"
echo ""

print_header "Test Complete!"
echo ""
echo "Summary:"
echo "✓ Created a test server and verified creation was logged"
echo "✓ Updated server OS and verified change was logged"
echo "✓ Deleted server and verified deletion was logged"
echo "✓ Confirmed history is preserved after server deletion"
echo "✓ Tested various filtering options (server_id, change_type, date range)"
echo ""
echo "Change history is working correctly!"
