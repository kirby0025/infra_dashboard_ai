-- Initialize the infra_dashboard database
-- This file is used by Docker Compose to set up the database schema

-- Create the servers table
CREATE TABLE IF NOT EXISTS servers (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE,
    os VARCHAR(100) NOT NULL,
    os_version VARCHAR(100) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- Create indexes for better query performance
CREATE INDEX IF NOT EXISTS idx_servers_name ON servers(name);
CREATE INDEX IF NOT EXISTS idx_servers_os ON servers(os);

-- Insert some sample data
INSERT INTO servers (name, os, os_version) VALUES
    ('web-server-01', 'Ubuntu', '22.04 LTS'),
    ('web-server-02', 'Ubuntu', '20.04 LTS'),
    ('db-server-01', 'CentOS', '8.5'),
    ('app-server-01', 'Red Hat Enterprise Linux', '9.0'),
    ('cache-server-01', 'Alpine Linux', '3.17')
ON CONFLICT (name) DO NOTHING;

-- Create a function to automatically update the updated_at timestamp
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ language 'plpgsql';

-- Create a trigger to automatically update the updated_at column
CREATE TRIGGER update_servers_updated_at
    BEFORE UPDATE ON servers
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();
