-- Test Sample Data for Infrastructure Dashboard
-- This file contains 300 sample servers with realistic OS distribution
-- Distribution: 70% Debian, 10% Ubuntu, 10% RedHat, 10% OpenBSD
-- EOL constraint: Maximum 10% (30 servers) running end-of-life operating systems

-- Clear existing sample data (optional)
-- DELETE FROM servers WHERE name LIKE '%-sample-%';

-- Insert 300 sample servers
-- The OS IDs are calculated based on the order in init.sql starting from ID 1

-- DEBIAN SERVERS (210 servers - 70%)
-- Mix of supported and EOL Debian versions
-- 15 EOL Debian servers, 195 supported Debian servers

-- EOL Debian servers (15 servers) - Using Debian 7, 8, 9
INSERT INTO servers (name, os_id) VALUES
('web-sample-001', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '7' LIMIT 1)),
('web-sample-002', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '8' LIMIT 1)),
('web-sample-003', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '9' LIMIT 1)),
('db-sample-001', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '7' LIMIT 1)),
('db-sample-002', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '8' LIMIT 1)),
('app-sample-001', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '9' LIMIT 1)),
('app-sample-002', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '7' LIMIT 1)),
('cache-sample-001', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '8' LIMIT 1)),
('mail-sample-001', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '9' LIMIT 1)),
('file-sample-001', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '7' LIMIT 1)),
('dns-sample-001', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '8' LIMIT 1)),
('proxy-sample-001', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '9' LIMIT 1)),
('backup-sample-001', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '7' LIMIT 1)),
('monitor-sample-001', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '8' LIMIT 1)),
('log-sample-001', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '9' LIMIT 1));

-- Supported Debian servers (195 servers) - Using Debian 10, 11, 12, 13
INSERT INTO servers (name, os_id) VALUES
-- Web servers (50 servers)
('web-sample-004', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '12' LIMIT 1)),
('web-sample-005', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '11' LIMIT 1)),
('web-sample-006', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '12' LIMIT 1)),
('web-sample-007', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '11' LIMIT 1)),
('web-sample-008', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '12' LIMIT 1)),
('web-sample-009', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '13' LIMIT 1)),
('web-sample-010', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '12' LIMIT 1)),
('web-sample-011', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '11' LIMIT 1)),
('web-sample-012', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '12' LIMIT 1)),
('web-sample-013', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '13' LIMIT 1)),
('web-sample-014', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '12' LIMIT 1)),
('web-sample-015', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '11' LIMIT 1)),
('web-sample-016', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '12' LIMIT 1)),
('web-sample-017', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '13' LIMIT 1)),
('web-sample-018', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '12' LIMIT 1)),
('web-sample-019', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '11' LIMIT 1)),
('web-sample-020', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '12' LIMIT 1)),
('web-sample-021', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '13' LIMIT 1)),
('web-sample-022', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '12' LIMIT 1)),
('web-sample-023', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '11' LIMIT 1)),
('web-sample-024', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '12' LIMIT 1)),
('web-sample-025', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '13' LIMIT 1)),
('web-sample-026', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '12' LIMIT 1)),
('web-sample-027', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '11' LIMIT 1)),
('web-sample-028', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '12' LIMIT 1)),
('web-sample-029', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '13' LIMIT 1)),
('web-sample-030', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '12' LIMIT 1)),
('web-sample-031', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '11' LIMIT 1)),
('web-sample-032', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '12' LIMIT 1)),
('web-sample-033', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '13' LIMIT 1)),
('web-sample-034', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '12' LIMIT 1)),
('web-sample-035', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '11' LIMIT 1)),
('web-sample-036', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '12' LIMIT 1)),
('web-sample-037', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '13' LIMIT 1)),
('web-sample-038', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '12' LIMIT 1)),
('web-sample-039', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '11' LIMIT 1)),
('web-sample-040', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '12' LIMIT 1)),
('web-sample-041', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '13' LIMIT 1)),
('web-sample-042', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '12' LIMIT 1)),
('web-sample-043', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '11' LIMIT 1)),
('web-sample-044', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '12' LIMIT 1)),
('web-sample-045', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '13' LIMIT 1)),
('web-sample-046', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '12' LIMIT 1)),
('web-sample-047', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '11' LIMIT 1)),
('web-sample-048', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '12' LIMIT 1)),
('web-sample-049', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '13' LIMIT 1)),
('web-sample-050', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '12' LIMIT 1)),
('web-sample-051', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '11' LIMIT 1)),
('web-sample-052', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '12' LIMIT 1)),
('web-sample-053', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '13' LIMIT 1)),

-- Database servers (50 servers)
('db-sample-003', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '12' LIMIT 1)),
('db-sample-004', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '11' LIMIT 1)),
('db-sample-005', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '12' LIMIT 1)),
('db-sample-006', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '13' LIMIT 1)),
('db-sample-007', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '12' LIMIT 1)),
('db-sample-008', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '11' LIMIT 1)),
('db-sample-009', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '12' LIMIT 1)),
('db-sample-010', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '13' LIMIT 1)),
('db-sample-011', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '12' LIMIT 1)),
('db-sample-012', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '11' LIMIT 1)),
('db-sample-013', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '12' LIMIT 1)),
('db-sample-014', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '13' LIMIT 1)),
('db-sample-015', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '12' LIMIT 1)),
('db-sample-016', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '11' LIMIT 1)),
('db-sample-017', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '12' LIMIT 1)),
('db-sample-018', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '13' LIMIT 1)),
('db-sample-019', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '12' LIMIT 1)),
('db-sample-020', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '11' LIMIT 1)),
('db-sample-021', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '12' LIMIT 1)),
('db-sample-022', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '13' LIMIT 1)),
('db-sample-023', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '12' LIMIT 1)),
('db-sample-024', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '11' LIMIT 1)),
('db-sample-025', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '12' LIMIT 1)),
('db-sample-026', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '13' LIMIT 1)),
('db-sample-027', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '12' LIMIT 1)),
('db-sample-028', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '11' LIMIT 1)),
('db-sample-029', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '12' LIMIT 1)),
('db-sample-030', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '13' LIMIT 1)),
('db-sample-031', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '12' LIMIT 1)),
('db-sample-032', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '11' LIMIT 1)),
('db-sample-033', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '12' LIMIT 1)),
('db-sample-034', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '13' LIMIT 1)),
('db-sample-035', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '12' LIMIT 1)),
('db-sample-036', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '11' LIMIT 1)),
('db-sample-037', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '12' LIMIT 1)),
('db-sample-038', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '13' LIMIT 1)),
('db-sample-039', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '12' LIMIT 1)),
('db-sample-040', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '11' LIMIT 1)),
('db-sample-041', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '12' LIMIT 1)),
('db-sample-042', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '13' LIMIT 1)),
('db-sample-043', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '12' LIMIT 1)),
('db-sample-044', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '11' LIMIT 1)),
('db-sample-045', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '12' LIMIT 1)),
('db-sample-046', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '13' LIMIT 1)),
('db-sample-047', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '12' LIMIT 1)),
('db-sample-048', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '11' LIMIT 1)),
('db-sample-049', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '12' LIMIT 1)),
('db-sample-050', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '13' LIMIT 1)),
('db-sample-051', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '12' LIMIT 1)),
('db-sample-052', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '11' LIMIT 1)),

-- Application servers (50 servers)
('app-sample-003', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '12' LIMIT 1)),
('app-sample-004', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '11' LIMIT 1)),
('app-sample-005', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '12' LIMIT 1)),
('app-sample-006', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '13' LIMIT 1)),
('app-sample-007', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '12' LIMIT 1)),
('app-sample-008', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '11' LIMIT 1)),
('app-sample-009', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '12' LIMIT 1)),
('app-sample-010', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '13' LIMIT 1)),
('app-sample-011', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '12' LIMIT 1)),
('app-sample-012', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '11' LIMIT 1)),
('app-sample-013', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '12' LIMIT 1)),
('app-sample-014', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '13' LIMIT 1)),
('app-sample-015', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '12' LIMIT 1)),
('app-sample-016', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '11' LIMIT 1)),
('app-sample-017', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '12' LIMIT 1)),
('app-sample-018', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '13' LIMIT 1)),
('app-sample-019', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '12' LIMIT 1)),
('app-sample-020', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '11' LIMIT 1)),
('app-sample-021', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '12' LIMIT 1)),
('app-sample-022', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '13' LIMIT 1)),
('app-sample-023', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '12' LIMIT 1)),
('app-sample-024', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '11' LIMIT 1)),
('app-sample-025', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '12' LIMIT 1)),
('app-sample-026', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '13' LIMIT 1)),
('app-sample-027', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '12' LIMIT 1)),
('app-sample-028', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '11' LIMIT 1)),
('app-sample-029', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '12' LIMIT 1)),
('app-sample-030', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '13' LIMIT 1)),
('app-sample-031', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '12' LIMIT 1)),
('app-sample-032', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '11' LIMIT 1)),
('app-sample-033', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '12' LIMIT 1)),
('app-sample-034', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '13' LIMIT 1)),
('app-sample-035', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '12' LIMIT 1)),
('app-sample-036', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '11' LIMIT 1)),
('app-sample-037', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '12' LIMIT 1)),
('app-sample-038', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '13' LIMIT 1)),
('app-sample-039', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '12' LIMIT 1)),
('app-sample-040', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '11' LIMIT 1)),
('app-sample-041', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '12' LIMIT 1)),
('app-sample-042', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '13' LIMIT 1)),
('app-sample-043', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '12' LIMIT 1)),
('app-sample-044', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '11' LIMIT 1)),
('app-sample-045', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '12' LIMIT 1)),
('app-sample-046', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '13' LIMIT 1)),
('app-sample-047', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '12' LIMIT 1)),
('app-sample-048', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '11' LIMIT 1)),
('app-sample-049', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '12' LIMIT 1)),
('app-sample-050', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '13' LIMIT 1)),
('app-sample-051', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '12' LIMIT 1)),
('app-sample-052', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '11' LIMIT 1)),

-- Remaining Debian servers (45 servers) - Various services
('cache-sample-002', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '12' LIMIT 1)),
('cache-sample-003', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '11' LIMIT 1)),
('cache-sample-004', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '13' LIMIT 1)),
('mail-sample-002', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '12' LIMIT 1)),
('mail-sample-003', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '11' LIMIT 1)),
('mail-sample-004', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '13' LIMIT 1)),
('file-sample-002', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '12' LIMIT 1)),
('file-sample-003', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '11' LIMIT 1)),
('file-sample-004', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '13' LIMIT 1)),
('dns-sample-002', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '12' LIMIT 1)),
('dns-sample-003', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '11' LIMIT 1)),
('dns-sample-004', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '13' LIMIT 1)),
('proxy-sample-002', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '12' LIMIT 1)),
('proxy-sample-003', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '11' LIMIT 1)),
('proxy-sample-004', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '13' LIMIT 1)),
('backup-sample-002', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '12' LIMIT 1)),
('backup-sample-003', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '11' LIMIT 1)),
('backup-sample-004', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '13' LIMIT 1)),
('monitor-sample-002', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '12' LIMIT 1)),
('monitor-sample-003', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '11' LIMIT 1)),
('monitor-sample-004', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '13' LIMIT 1)),
('log-sample-002', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '12' LIMIT 1)),
('log-sample-003', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '11' LIMIT 1)),
('log-sample-004', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '13' LIMIT 1)),
('api-sample-001', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '12' LIMIT 1)),
('api-sample-002', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '11' LIMIT 1)),
('api-sample-003', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '13' LIMIT 1)),
('queue-sample-001', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '12' LIMIT 1)),
('queue-sample-002', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '11' LIMIT 1)),
('queue-sample-003', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '13' LIMIT 1)),
('search-sample-001', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '12' LIMIT 1)),
('search-sample-002', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '11' LIMIT 1)),
('search-sample-003', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '13' LIMIT 1)),
('worker-sample-001', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '12' LIMIT 1)),
('worker-sample-002', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '11' LIMIT 1)),
('worker-sample-003', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '13' LIMIT 1)),
('storage-sample-001', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '12' LIMIT 1)),
('storage-sample-002', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '11' LIMIT 1)),
('storage-sample-003', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '13' LIMIT 1)),
('metrics-sample-001', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '12' LIMIT 1)),
('metrics-sample-002', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '11' LIMIT 1)),
('metrics-sample-003', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '13' LIMIT 1)),
('gateway-sample-001', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '12' LIMIT 1)),
('gateway-sample-002', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '11' LIMIT 1)),
('gateway-sample-003', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '13' LIMIT 1)),
('config-sample-001', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '12' LIMIT 1)),
('config-sample-002', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '11' LIMIT 1)),
('config-sample-003', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '13' LIMIT 1));

-- UBUNTU SERVERS (30 servers - 10%)
-- Mix of supported and EOL Ubuntu versions
-- 5 EOL Ubuntu servers, 25 supported Ubuntu servers

-- EOL Ubuntu servers (5 servers) - Using Ubuntu 16.04, 18.04
INSERT INTO servers (name, os_id) VALUES
('ubuntu-web-sample-001', (SELECT id FROM operating_systems WHERE name = 'Ubuntu' AND version = '16.04' LIMIT 1)),
('ubuntu-web-sample-002', (SELECT id FROM operating_systems WHERE name = 'Ubuntu' AND version = '18.04' LIMIT 1)),
('ubuntu-db-sample-001', (SELECT id FROM operating_systems WHERE name = 'Ubuntu' AND version = '16.04' LIMIT 1)),
('ubuntu-app-sample-001', (SELECT id FROM operating_systems WHERE name = 'Ubuntu' AND version = '18.04' LIMIT 1)),
('ubuntu-api-sample-001', (SELECT id FROM operating_systems WHERE name = 'Ubuntu' AND version = '16.04' LIMIT 1));

-- Supported Ubuntu servers (25 servers) - Using Ubuntu 20.04, 22.04, 24.04
INSERT INTO servers (name, os_id) VALUES
('ubuntu-web-sample-003', (SELECT id FROM operating_systems WHERE name = 'Ubuntu' AND version = '22.04' LIMIT 1)),
('ubuntu-web-sample-004', (SELECT id FROM operating_systems WHERE name = 'Ubuntu' AND version = '20.04' LIMIT 1)),
('ubuntu-web-sample-005', (SELECT id FROM operating_systems WHERE name = 'Ubuntu' AND version = '24.04' LIMIT 1)),
('ubuntu-web-sample-006', (SELECT id FROM operating_systems WHERE name = 'Ubuntu' AND version = '22.04' LIMIT 1)),
('ubuntu-web-sample-007', (SELECT id FROM operating_systems WHERE name = 'Ubuntu' AND version = '20.04' LIMIT 1)),
('ubuntu-db-sample-002', (SELECT id FROM operating_systems WHERE name = 'Ubuntu' AND version = '22.04' LIMIT 1)),
('ubuntu-db-sample-003', (SELECT id FROM operating_systems WHERE name = 'Ubuntu' AND version = '20.04' LIMIT 1)),
('ubuntu-db-sample-004', (SELECT id FROM operating_systems WHERE name = 'Ubuntu' AND version = '24.04' LIMIT 1)),
('ubuntu-db-sample-005', (SELECT id FROM operating_systems WHERE name = 'Ubuntu' AND version = '22.04' LIMIT 1)),
('ubuntu-app-sample-002', (SELECT id FROM operating_systems WHERE name = 'Ubuntu' AND version = '22.04' LIMIT 1)),
('ubuntu-app-sample-003', (SELECT id FROM operating_systems WHERE name = 'Ubuntu' AND version = '20.04' LIMIT 1)),
('ubuntu-app-sample-004', (SELECT id FROM operating_systems WHERE name = 'Ubuntu' AND version = '24.04' LIMIT 1)),
('ubuntu-app-sample-005', (SELECT id FROM operating_systems WHERE name = 'Ubuntu' AND version = '22.04' LIMIT 1)),
('ubuntu-api-sample-002', (SELECT id FROM operating_systems WHERE name = 'Ubuntu' AND version = '20.04' LIMIT 1)),
('ubuntu-api-sample-003', (SELECT id FROM operating_systems WHERE name = 'Ubuntu' AND version = '22.04' LIMIT 1)),
('ubuntu-api-sample-004', (SELECT id FROM operating_systems WHERE name = 'Ubuntu' AND version = '24.04' LIMIT 1)),
('ubuntu-cache-sample-001', (SELECT id FROM operating_systems WHERE name = 'Ubuntu' AND version = '22.04' LIMIT 1)),
('ubuntu-cache-sample-002', (SELECT id FROM operating_systems WHERE name = 'Ubuntu' AND version = '20.04' LIMIT 1)),
('ubuntu-mail-sample-001', (SELECT id FROM operating_systems WHERE name = 'Ubuntu' AND version = '22.04' LIMIT 1)),
('ubuntu-file-sample-001', (SELECT id FROM operating_systems WHERE name = 'Ubuntu' AND version = '20.04' LIMIT 1)),
('ubuntu-dns-sample-001', (SELECT id FROM operating_systems WHERE name = 'Ubuntu' AND version = '24.04' LIMIT 1)),
('ubuntu-proxy-sample-001', (SELECT id FROM operating_systems WHERE name = 'Ubuntu' AND version = '22.04' LIMIT 1)),
('ubuntu-backup-sample-001', (SELECT id FROM operating_systems WHERE name = 'Ubuntu' AND version = '20.04' LIMIT 1)),
('ubuntu-monitor-sample-001', (SELECT id FROM operating_systems WHERE name = 'Ubuntu' AND version = '22.04' LIMIT 1)),
('ubuntu-log-sample-001', (SELECT id FROM operating_systems WHERE name = 'Ubuntu' AND version = '24.04' LIMIT 1)),
('ubuntu-queue-sample-001', (SELECT id FROM operating_systems WHERE name = 'Ubuntu' AND version = '22.04' LIMIT 1)),
('ubuntu-search-sample-001', (SELECT id FROM operating_systems WHERE name = 'Ubuntu' AND version = '20.04' LIMIT 1)),
('ubuntu-worker-sample-001', (SELECT id FROM operating_systems WHERE name = 'Ubuntu' AND version = '22.04' LIMIT 1)),
('ubuntu-storage-sample-001', (SELECT id FROM operating_systems WHERE name = 'Ubuntu' AND version = '24.04' LIMIT 1));

-- REDHAT SERVERS (30 servers - 10%)
-- Mix of supported and EOL RedHat versions
-- 5 EOL RedHat servers, 25 supported RedHat servers

-- EOL RedHat servers (5 servers) - Using RedHat 5, 6
INSERT INTO servers (name, os_id) VALUES
('redhat-web-sample-001', (SELECT id FROM operating_systems WHERE name = 'RedHat' AND version = '5' LIMIT 1)),
('redhat-web-sample-002', (SELECT id FROM operating_systems WHERE name = 'RedHat' AND version = '6' LIMIT 1)),
('redhat-db-sample-001', (SELECT id FROM operating_systems WHERE name = 'RedHat' AND version = '5' LIMIT 1)),
('redhat-app-sample-001', (SELECT id FROM operating_systems WHERE name = 'RedHat' AND version = '6' LIMIT 1)),
('redhat-api-sample-001', (SELECT id FROM operating_systems WHERE name = 'RedHat' AND version = '5' LIMIT 1));

-- Supported RedHat servers (25 servers) - Using RedHat 8, 9, 10
INSERT INTO servers (name, os_id) VALUES
('redhat-web-sample-003', (SELECT id FROM operating_systems WHERE name = 'RedHat' AND version = '8' LIMIT 1)),
('redhat-web-sample-004', (SELECT id FROM operating_systems WHERE name = 'RedHat' AND version = '9' LIMIT 1)),
('redhat-web-sample-005', (SELECT id FROM operating_systems WHERE name = 'RedHat' AND version = '10' LIMIT 1)),
('redhat-web-sample-006', (SELECT id FROM operating_systems WHERE name = 'RedHat' AND version = '8' LIMIT 1)),
('redhat-web-sample-007', (SELECT id FROM operating_systems WHERE name = 'RedHat' AND version = '9' LIMIT 1)),
('redhat-db-sample-002', (SELECT id FROM operating_systems WHERE name = 'RedHat' AND version = '8' LIMIT 1)),
('redhat-db-sample-003', (SELECT id FROM operating_systems WHERE name = 'RedHat' AND version = '9' LIMIT 1)),
('redhat-db-sample-004', (SELECT id FROM operating_systems WHERE name = 'RedHat' AND version = '10' LIMIT 1)),
('redhat-db-sample-005', (SELECT id FROM operating_systems WHERE name = 'RedHat' AND version = '8' LIMIT 1)),
('redhat-app-sample-002', (SELECT id FROM operating_systems WHERE name = 'RedHat' AND version = '9' LIMIT 1)),
('redhat-app-sample-003', (SELECT id FROM operating_systems WHERE name = 'RedHat' AND version = '8' LIMIT 1)),
('redhat-app-sample-004', (SELECT id FROM operating_systems WHERE name = 'RedHat' AND version = '10' LIMIT 1)),
('redhat-app-sample-005', (SELECT id FROM operating_systems WHERE name = 'RedHat' AND version = '9' LIMIT 1)),
('redhat-api-sample-002', (SELECT id FROM operating_systems WHERE name = 'RedHat' AND version = '8' LIMIT 1)),
('redhat-api-sample-003', (SELECT id FROM operating_systems WHERE name = 'RedHat' AND version = '9' LIMIT 1)),
('redhat-cache-sample-001', (SELECT id FROM operating_systems WHERE name = 'RedHat' AND version = '10' LIMIT 1)),
('redhat-cache-sample-002', (SELECT id FROM operating_systems WHERE name = 'RedHat' AND version = '8' LIMIT 1)),
('redhat-mail-sample-001', (SELECT id FROM operating_systems WHERE name = 'RedHat' AND version = '9' LIMIT 1)),
('redhat-file-sample-001', (SELECT id FROM operating_systems WHERE name = 'RedHat' AND version = '8' LIMIT 1)),
('redhat-dns-sample-001', (SELECT id FROM operating_systems WHERE name = 'RedHat' AND version = '10' LIMIT 1)),
('redhat-proxy-sample-001', (SELECT id FROM operating_systems WHERE name = 'RedHat' AND version = '9' LIMIT 1)),
('redhat-backup-sample-001', (SELECT id FROM operating_systems WHERE name = 'RedHat' AND version = '8' LIMIT 1)),
('redhat-monitor-sample-001', (SELECT id FROM operating_systems WHERE name = 'RedHat' AND version = '9' LIMIT 1)),
('redhat-log-sample-001', (SELECT id FROM operating_systems WHERE name = 'RedHat' AND version = '10' LIMIT 1)),
('redhat-queue-sample-001', (SELECT id FROM operating_systems WHERE name = 'RedHat' AND version = '8' LIMIT 1));

-- OPENBSD SERVERS (30 servers - 10%)
-- Mix of supported and EOL OpenBSD versions
-- 5 EOL OpenBSD servers, 25 supported OpenBSD servers

-- EOL OpenBSD servers (5 servers) - Using OpenBSD 6.8, 6.9, 7.0
INSERT INTO servers (name, os_id) VALUES
('openbsd-firewall-001', (SELECT id FROM operating_systems WHERE name = 'OpenBSD' AND version = '6.8' LIMIT 1)),
('openbsd-firewall-002', (SELECT id FROM operating_systems WHERE name = 'OpenBSD' AND version = '6.9' LIMIT 1)),
('openbsd-router-001', (SELECT id FROM operating_systems WHERE name = 'OpenBSD' AND version = '7.0' LIMIT 1)),
('openbsd-gateway-001', (SELECT id FROM operating_systems WHERE name = 'OpenBSD' AND version = '6.8' LIMIT 1)),
('openbsd-security-001', (SELECT id FROM operating_systems WHERE name = 'OpenBSD' AND version = '6.9' LIMIT 1));

-- Supported OpenBSD servers (25 servers) - Using OpenBSD 7.4, 7.5, 7.6
INSERT INTO servers (name, os_id) VALUES
('openbsd-firewall-003', (SELECT id FROM operating_systems WHERE name = 'OpenBSD' AND version = '7.5' LIMIT 1)),
('openbsd-firewall-004', (SELECT id FROM operating_systems WHERE name = 'OpenBSD' AND version = '7.6' LIMIT 1)),
('openbsd-firewall-005', (SELECT id FROM operating_systems WHERE name = 'OpenBSD' AND version = '7.4' LIMIT 1)),
('openbsd-firewall-006', (SELECT id FROM operating_systems WHERE name = 'OpenBSD' AND version = '7.5' LIMIT 1)),
('openbsd-firewall-007', (SELECT id FROM operating_systems WHERE name = 'OpenBSD' AND version = '7.6' LIMIT 1)),
('openbsd-router-002', (SELECT id FROM operating_systems WHERE name = 'OpenBSD' AND version = '7.4' LIMIT 1)),
('openbsd-router-003', (SELECT id FROM operating_systems WHERE name = 'OpenBSD' AND version = '7.5' LIMIT 1)),
('openbsd-router-004', (SELECT id FROM operating_systems WHERE name = 'OpenBSD' AND version = '7.6' LIMIT 1)),
('openbsd-router-005', (SELECT id FROM operating_systems WHERE name = 'OpenBSD' AND version = '7.4' LIMIT 1)),
('openbsd-gateway-002', (SELECT id FROM operating_systems WHERE name = 'OpenBSD' AND version = '7.5' LIMIT 1)),
('openbsd-gateway-003', (SELECT id FROM operating_systems WHERE name = 'OpenBSD' AND version = '7.6' LIMIT 1)),
('openbsd-gateway-004', (SELECT id FROM operating_systems WHERE name = 'OpenBSD' AND version = '7.4' LIMIT 1)),
('openbsd-security-002', (SELECT id FROM operating_systems WHERE name = 'OpenBSD' AND version = '7.5' LIMIT 1)),
('openbsd-security-003', (SELECT id FROM operating_systems WHERE name = 'OpenBSD' AND version = '7.6' LIMIT 1)),
('openbsd-security-004', (SELECT id FROM operating_systems WHERE name = 'OpenBSD' AND version = '7.4' LIMIT 1)),
('openbsd-dns-001', (SELECT id FROM operating_systems WHERE name = 'OpenBSD' AND version = '7.5' LIMIT 1)),
('openbsd-dns-002', (SELECT id FROM operating_systems WHERE name = 'OpenBSD' AND version = '7.6' LIMIT 1)),
('openbsd-dns-003', (SELECT id FROM operating_systems WHERE name = 'OpenBSD' AND version = '7.4' LIMIT 1)),
('openbsd-vpn-001', (SELECT id FROM operating_systems WHERE name = 'OpenBSD' AND version = '7.5' LIMIT 1)),
('openbsd-vpn-002', (SELECT id FROM operating_systems WHERE name = 'OpenBSD' AND version = '7.6' LIMIT 1)),
('openbsd-vpn-003', (SELECT id FROM operating_systems WHERE name = 'OpenBSD' AND version = '7.4' LIMIT 1)),
('openbsd-ntp-001', (SELECT id FROM operating_systems WHERE name = 'OpenBSD' AND version = '7.5' LIMIT 1)),
('openbsd-ntp-002', (SELECT id FROM operating_systems WHERE name = 'OpenBSD' AND version = '7.6' LIMIT 1)),
('openbsd-monitor-001', (SELECT id FROM operating_systems WHERE name = 'OpenBSD' AND version = '7.4' LIMIT 1)),
('openbsd-bastion-001', (SELECT id FROM operating_systems WHERE name = 'OpenBSD' AND version = '7.5' LIMIT 1));

-- Summary of sample data:
-- Total servers: 300
-- Distribution:
--   - Debian: 210 servers (70%) - 15 EOL, 195 supported
--   - Ubuntu: 30 servers (10%) - 5 EOL, 25 supported
--   - RedHat: 30 servers (10%) - 5 EOL, 25 supported
--   - OpenBSD: 30 servers (10%) - 5 EOL, 25 supported
-- Total EOL servers: 30 (10% - meets constraint)
-- Total supported servers: 270 (90%)

-- To load this sample data, run:
-- psql -U postgres -d infra_dashboard -f test_sample.sql

-- To verify the distribution, run:
-- SELECT
--   os.name as os_family,
--   COUNT(*) as server_count,
--   ROUND(COUNT(*) * 100.0 / (SELECT COUNT(*) FROM servers), 1) as percentage
-- FROM servers s
-- JOIN operating_systems os ON s.os_id = os.id
-- WHERE s.name LIKE '%-sample-%'
-- GROUP BY os.name
-- ORDER BY server_count DESC;

-- To check EOL status:
-- SELECT
--   CASE WHEN os.end_of_support < CURRENT_DATE THEN 'EOL' ELSE 'Supported' END as status,
--   COUNT(*) as server_count,
--   ROUND(COUNT(*) * 100.0 / (SELECT COUNT(*) FROM servers WHERE name LIKE '%-sample-%'), 1) as percentage
-- FROM servers s
-- JOIN operating_systems os ON s.os_id = os.id
-- WHERE s.name LIKE '%-sample-%'
-- GROUP BY CASE WHEN os.end_of_support < CURRENT_DATE THEN 'EOL' ELSE 'Supported' END
-- ORDER BY server_count DESC
