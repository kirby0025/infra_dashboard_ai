# Migration Guide: Adding Change History Feature

## Overview

This guide helps you add the Server Change History feature to an existing Infrastructure Dashboard deployment.

## Prerequisites

- PostgreSQL database access with DDL permissions
- Backup of existing database
- Application downtime tolerance (optional - see Zero Downtime section)

## Migration Steps

### Step 1: Backup Your Database

**CRITICAL: Always backup before schema changes**

```bash
# Create a backup
pg_dump -h localhost -U postgres -d infra_dashboard > backup_$(date +%Y%m%d_%H%M%S).sql

# Verify backup
ls -lh backup_*.sql
```

### Step 2: Apply Database Schema Changes

Connect to your PostgreSQL database and execute the following SQL:

```sql
-- Create the server_change_history table
CREATE TABLE IF NOT EXISTS server_change_history (
    id SERIAL PRIMARY KEY,
    server_id INTEGER,
    server_name VARCHAR(255) NOT NULL,
    change_type VARCHAR(50) NOT NULL, -- 'created', 'os_changed', 'deleted'
    old_os_id INTEGER,
    new_os_id INTEGER,
    old_os_name VARCHAR(100),
    old_os_version VARCHAR(100),
    new_os_name VARCHAR(100),
    new_os_version VARCHAR(100),
    changed_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    FOREIGN KEY (server_id) REFERENCES servers(id) ON DELETE SET NULL,
    FOREIGN KEY (old_os_id) REFERENCES operating_systems(id) ON DELETE SET NULL,
    FOREIGN KEY (new_os_id) REFERENCES operating_systems(id) ON DELETE SET NULL
);

-- Create indexes for better query performance
CREATE INDEX IF NOT EXISTS idx_change_history_server_id ON server_change_history(server_id);
CREATE INDEX IF NOT EXISTS idx_change_history_change_type ON server_change_history(change_type);
CREATE INDEX IF NOT EXISTS idx_change_history_changed_at ON server_change_history(changed_at);

-- Create function to log server creation
CREATE OR REPLACE FUNCTION log_server_creation()
RETURNS TRIGGER AS $$
BEGIN
    INSERT INTO server_change_history (
        server_id,
        server_name,
        change_type,
        new_os_id,
        new_os_name,
        new_os_version
    )
    SELECT
        NEW.id,
        NEW.name,
        'created',
        NEW.os_id,
        os.name,
        os.version
    FROM operating_systems os
    WHERE os.id = NEW.os_id;

    RETURN NEW;
END;
$$ LANGUAGE 'plpgsql';

-- Create function to log OS changes
CREATE OR REPLACE FUNCTION log_server_os_change()
RETURNS TRIGGER AS $$
BEGIN
    -- Only log if os_id actually changed
    IF OLD.os_id != NEW.os_id THEN
        INSERT INTO server_change_history (
            server_id,
            server_name,
            change_type,
            old_os_id,
            new_os_id,
            old_os_name,
            old_os_version,
            new_os_name,
            new_os_version
        )
        SELECT
            NEW.id,
            NEW.name,
            'os_changed',
            OLD.os_id,
            NEW.os_id,
            old_os.name,
            old_os.version,
            new_os.name,
            new_os.version
        FROM operating_systems old_os, operating_systems new_os
        WHERE old_os.id = OLD.os_id AND new_os.id = NEW.os_id;
    END IF;

    RETURN NEW;
END;
$$ LANGUAGE 'plpgsql';

-- Create function to log server deletion
CREATE OR REPLACE FUNCTION log_server_deletion()
RETURNS TRIGGER AS $$
BEGIN
    INSERT INTO server_change_history (
        server_id,
        server_name,
        change_type,
        old_os_id,
        old_os_name,
        old_os_version
    )
    SELECT
        OLD.id,
        OLD.name,
        'deleted',
        OLD.os_id,
        os.name,
        os.version
    FROM operating_systems os
    WHERE os.id = OLD.os_id;

    RETURN OLD;
END;
$$ LANGUAGE 'plpgsql';

-- Create triggers for logging server changes
CREATE TRIGGER log_server_creation_trigger
    AFTER INSERT ON servers
    FOR EACH ROW
    EXECUTE FUNCTION log_server_creation();

CREATE TRIGGER log_server_os_change_trigger
    AFTER UPDATE ON servers
    FOR EACH ROW
    EXECUTE FUNCTION log_server_os_change();

CREATE TRIGGER log_server_deletion_trigger
    BEFORE DELETE ON servers
    FOR EACH ROW
    EXECUTE FUNCTION log_server_deletion();
```

### Step 3: Verify Database Changes

```sql
-- Verify table exists
SELECT table_name FROM information_schema.tables 
WHERE table_schema = 'public' AND table_name = 'server_change_history';

-- Verify indexes
SELECT indexname FROM pg_indexes 
WHERE tablename = 'server_change_history';

-- Verify triggers
SELECT tgname, tgtype, tgenabled 
FROM pg_trigger 
WHERE tgname LIKE 'log_server%';

-- Verify functions
SELECT routine_name 
FROM information_schema.routines 
WHERE routine_schema = 'public' AND routine_name LIKE 'log_server%';
```

Expected output:
- 1 table: `server_change_history`
- 3 indexes: `idx_change_history_server_id`, `idx_change_history_change_type`, `idx_change_history_changed_at`
- 3 triggers: `log_server_creation_trigger`, `log_server_os_change_trigger`, `log_server_deletion_trigger`
- 3 functions: `log_server_creation`, `log_server_os_change`, `log_server_deletion`

### Step 4: (Optional) Backfill Historical Data

If you want to create historical records for existing servers:

```sql
-- Create 'created' records for all existing servers
INSERT INTO server_change_history (
    server_id,
    server_name,
    change_type,
    new_os_id,
    new_os_name,
    new_os_version,
    changed_at
)
SELECT
    s.id,
    s.name,
    'created',
    s.os_id,
    os.name,
    os.version,
    s.created_at
FROM servers s
JOIN operating_systems os ON s.os_id = os.id;
```

**Note**: This creates synthetic history records. The `changed_at` will be set to the server's `created_at` timestamp.

### Step 5: Deploy New Application Code

```bash
# Pull the latest code
git pull origin main

# Build the application
cd app
go build -o infra-dashboard ./cmd/main.go

# Stop the old version
sudo systemctl stop infra-dashboard
# or
docker compose down

# Deploy the new version
sudo systemctl start infra-dashboard
# or
docker compose up -d

# Verify the service is running
curl http://localhost:8080/health
```

### Step 6: Verify Change History is Working

```bash
# Test the new endpoints
curl http://localhost:8080/api/v1/history

# Create a test server
curl -X POST http://localhost:8080/api/v1/servers \
  -H "Content-Type: application/json" \
  -d '{"name":"migration-test-server","os_id":28}'

# Verify the creation was logged
curl "http://localhost:8080/api/v1/history?change_type=created" | tail -20

# Or run the automated test script
chmod +x test_change_history.sh
./test_change_history.sh
```

## Zero Downtime Migration (Advanced)

For production systems that require zero downtime:

### Phase 1: Add Schema (Application Still Running)

1. Add the table, functions, and triggers while the application is running
2. The triggers will start logging changes immediately
3. Existing functionality is not affected

### Phase 2: Deploy New Code (Rolling Restart)

1. Deploy new application version with change history endpoints
2. Use rolling restart to avoid downtime
3. New endpoints become available gradually

This works because:
- Database changes are additive (no breaking changes)
- Old code doesn't need the new table
- New code benefits from the new features

## Rollback Procedure

If you need to rollback the migration:

```sql
-- Remove triggers first
DROP TRIGGER IF EXISTS log_server_creation_trigger ON servers;
DROP TRIGGER IF EXISTS log_server_os_change_trigger ON servers;
DROP TRIGGER IF EXISTS log_server_deletion_trigger ON servers;

-- Remove functions
DROP FUNCTION IF EXISTS log_server_creation();
DROP FUNCTION IF EXISTS log_server_os_change();
DROP FUNCTION IF EXISTS log_server_deletion();

-- (Optional) Drop the table if you want to remove all history
DROP TABLE IF EXISTS server_change_history;

-- Restore from backup if needed
psql -h localhost -U postgres -d infra_dashboard < backup_YYYYMMDD_HHMMSS.sql
```

**Note**: If you drop the table, all change history will be lost. Consider exporting first:

```bash
# Export change history before dropping
pg_dump -h localhost -U postgres -d infra_dashboard -t server_change_history > change_history_backup.sql
```

## Post-Migration Tasks

### 1. Monitor Performance

```sql
-- Check table size
SELECT pg_size_pretty(pg_total_relation_size('server_change_history')) as size;

-- Check index usage
SELECT schemaname, tablename, indexname, idx_scan
FROM pg_stat_user_indexes
WHERE tablename = 'server_change_history';

-- Check trigger overhead
SELECT schemaname, tablename, n_tup_ins, n_tup_upd, n_tup_del
FROM pg_stat_user_tables
WHERE tablename IN ('servers', 'server_change_history');
```

### 2. Set Up Retention Policy (Optional)

For large deployments, consider implementing a retention policy:

```sql
-- Create a function to clean old history (e.g., older than 2 years)
CREATE OR REPLACE FUNCTION cleanup_old_change_history()
RETURNS INTEGER AS $$
DECLARE
    deleted_count INTEGER;
BEGIN
    DELETE FROM server_change_history
    WHERE changed_at < NOW() - INTERVAL '2 years';
    
    GET DIAGNOSTICS deleted_count = ROW_COUNT;
    RETURN deleted_count;
END;
$$ LANGUAGE plpgsql;

-- Run manually or set up a cron job
SELECT cleanup_old_change_history();
```

### 3. Configure Monitoring

Add monitoring for:
- Change history table growth rate
- Query performance on history endpoints
- Trigger execution time (if performance-sensitive)

### 4. Update Documentation

- Inform team about new endpoints
- Update API documentation
- Add to incident response procedures
- Include in compliance reports

## Troubleshooting

### Triggers Not Firing

Check if triggers are enabled:

```sql
SELECT tgname, tgenabled 
FROM pg_trigger 
WHERE tgname LIKE 'log_server%';
```

If disabled, enable them:

```sql
ALTER TABLE servers ENABLE TRIGGER log_server_creation_trigger;
ALTER TABLE servers ENABLE TRIGGER log_server_os_change_trigger;
ALTER TABLE servers ENABLE TRIGGER log_server_deletion_trigger;
```

### Performance Issues

If you experience performance degradation:

1. Check if indexes are being used:
```sql
EXPLAIN ANALYZE SELECT * FROM server_change_history WHERE server_id = 1;
```

2. Rebuild indexes if needed:
```sql
REINDEX TABLE server_change_history;
```

3. Update table statistics:
```sql
ANALYZE server_change_history;
```

### Foreign Key Violations

If you see foreign key errors, verify referential integrity:

```sql
-- Check for orphaned server_ids
SELECT DISTINCT server_id 
FROM server_change_history 
WHERE server_id IS NOT NULL 
  AND server_id NOT IN (SELECT id FROM servers);

-- Check for orphaned os_ids
SELECT DISTINCT old_os_id 
FROM server_change_history 
WHERE old_os_id IS NOT NULL 
  AND old_os_id NOT IN (SELECT id FROM operating_systems);
```

## Testing Checklist

After migration, verify:

- [ ] Database table created successfully
- [ ] All indexes are present
- [ ] All triggers are active
- [ ] All functions are created
- [ ] Health check endpoint responds
- [ ] Can query /api/v1/history
- [ ] Create server logs to history
- [ ] Update server OS logs to history
- [ ] Delete server logs to history
- [ ] History preserved after server deletion
- [ ] Filters work (server_id, change_type, date range)
- [ ] Pagination works
- [ ] No performance degradation on existing endpoints

## Support

If you encounter issues:

1. Check application logs: `docker compose logs -f` or `journalctl -u infra-dashboard`
2. Check PostgreSQL logs: `tail -f /var/log/postgresql/postgresql-*.log`
3. Review the CHANGE_HISTORY.md documentation
4. Run the test script: `./test_change_history.sh`

## Migration Checklist

- [ ] Database backed up
- [ ] Schema changes applied
- [ ] Indexes created
- [ ] Triggers active
- [ ] Functions created
- [ ] Verification queries run successfully
- [ ] (Optional) Historical data backfilled
- [ ] New application code deployed
- [ ] Service restarted successfully
- [ ] Health check passing
- [ ] Change history endpoints accessible
- [ ] Test scenario completed
- [ ] Team notified
- [ ] Documentation updated

## Estimated Duration

- Small deployment (< 1000 servers): 15-30 minutes
- Medium deployment (1000-10000 servers): 30-60 minutes
- Large deployment (> 10000 servers): 1-2 hours

Most time is spent on:
- Database backup
- Optional historical data backfill
- Testing and verification

The actual schema changes take less than 1 minute.