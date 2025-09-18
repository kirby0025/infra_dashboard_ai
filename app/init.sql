-- Initialize the infra_dashboard database
-- This file is used by Docker Compose to set up the database schema

-- Create a function to automatically update the updated_at timestamp
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ language 'plpgsql';

-- Create the operating_systems table
CREATE TABLE IF NOT EXISTS operating_systems (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    version VARCHAR(100) NOT NULL,
    end_of_support DATE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    UNIQUE(name, version)
);

-- Create indexes for better query performance
CREATE INDEX IF NOT EXISTS idx_os_name ON operating_systems(name);
CREATE INDEX IF NOT EXISTS idx_os_end_of_support ON operating_systems(end_of_support);

-- Insert OS data (discarding the first value from each tuple)
INSERT INTO operating_systems (name, version, end_of_support) VALUES
    ('Debian','4','2010-02-28'),
    ('Debian','5','2012-02-06'),
    ('Debian','6','2016-02-06'),
    ('Debian','7','2018-04-01'),
    ('Debian','8','2020-04-01'),
    ('Debian','9','2022-06-30'),
    ('Debian','10','2024-06-30'),
    ('Debian','11','2026-06-30'),
    ('Debian','12','2028-06-30'),
    ('Debian','13','2030-06-30'),
    ('Ubuntu','10.04','2015-04-30'),
    ('Ubuntu','10.10','2012-04-10'),
    ('Ubuntu','11.04','2012-10-28'),
    ('Ubuntu','11.10','2013-05-09'),
    ('Ubuntu','12.04','2017-04-01'),
    ('Ubuntu','12.10','2014-05-16'),
    ('Ubuntu','13.04','2014-01-27'),
    ('Ubuntu','13.10','2014-07-17'),
    ('Ubuntu','14.04','2019-04-01'),
    ('Ubuntu','14.10','2015-07-23'),
    ('Ubuntu','15.04','2016-02-04'),
    ('Ubuntu','15.10','2016-07-28'),
    ('Ubuntu','16.04','2021-04-01'),
    ('Ubuntu','16.10','2017-07-01'),
    ('Ubuntu','17.04','2018-01-31'),
    ('Ubuntu','17.10','2018-07-31'),
    ('Ubuntu','18.04','2023-04-01'),
    ('Ubuntu','20.04','2025-04-01'),
    ('Ubuntu','22.04','2027-04-01'),
    ('Ubuntu','24.04','2029-04-01'),
    ('Ubuntu','26.04','2031-04-01'),
    ('CentOS','4','2012-02-29'),
    ('CentOS','5','2017-03-31'),
    ('CentOS','6','2020-11-30'),
    ('CentOS','7','2024-06-30'),
    ('RedHat','5','2017-03-31'),
    ('RedHat','6','2020-11-30'),
    ('RedHat','7','2024-06-30'),
    ('FreeBSD','9.0','2013-03-31'),
    ('FreeBSD','9.1','2014-12-31'),
    ('FreeBSD','9.2','2014-12-31'),
    ('FreeBSD','9.3','2016-12-31'),
    ('FreeBSD','10.0','2015-02-28'),
    ('FreeBSD','10.1','2016-12-31'),
    ('FreeBSD','10.2','2016-12-31'),
    ('FreeBSD','10.3','2018-04-30'),
    ('OpenBSD','5.0','2012-11-01'),
    ('OpenBSD','5.1','2013-05-01'),
    ('OpenBSD','5.2','2013-11-01'),
    ('OpenBSD','5.3','2014-05-01'),
    ('OpenBSD','5.4','2014-11-01'),
    ('OpenBSD','5.5','2015-05-01'),
    ('OpenBSD','5.6','2015-10-18'),
    ('OpenBSD','5.7','2016-03-29'),
    ('OpenBSD','5.8','2016-09-01'),
    ('OpenBSD','5.9','2017-04-11'),
    ('OpenBSD','6.0','2017-11-09'),
    ('OpenBSD','6.1','2018-05-01'),
    ('OpenBSD','6.2','2018-11-01'),
    ('OpenBSD','6.3','2019-05-01'),
    ('OpenBSD','6.4','2019-11-01'),
    ('OpenBSD','6.5','2020-05-19'),
    ('OpenBSD','6.6','2020-10-18'),
    ('OpenBSD','6.7','2021-05-01'),
    ('OpenBSD','6.8','2021-10-14'),
    ('OpenBSD','6.9','2022-04-21'),
    ('OpenBSD','7.0','2022-10-20'),
    ('OpenBSD','7.1','2023-04-10'),
    ('OpenBSD','7.2','2023-10-16'),
    ('OpenBSD','7.3','2024-04-05'),
    ('OpenBSD','7.4','2024-10-08'),
    ('OpenBSD','7.5','2025-04-28'),
    ('OpenBSD','7.6','2025-11-01')
ON CONFLICT (name, version) DO NOTHING;

-- Create a trigger to automatically update the updated_at column for operating_systems
CREATE TRIGGER update_operating_systems_updated_at
    BEFORE UPDATE ON operating_systems
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();

-- Create the servers table
CREATE TABLE IF NOT EXISTS servers (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE,
    os_id INTEGER NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    FOREIGN KEY (os_id) REFERENCES operating_systems(id)
);

-- Create indexes for better query performance
CREATE INDEX IF NOT EXISTS idx_servers_name ON servers(name);
CREATE INDEX IF NOT EXISTS idx_servers_os_id ON servers(os_id);

-- Insert some sample data (using OS IDs from operating_systems table)
INSERT INTO servers (name, os_id) VALUES
    ('web-server-01', (SELECT id FROM operating_systems WHERE name = 'Ubuntu' AND version = '22.04' LIMIT 1)),
    ('web-server-02', (SELECT id FROM operating_systems WHERE name = 'Ubuntu' AND version = '20.04' LIMIT 1)),
    ('db-server-01', (SELECT id FROM operating_systems WHERE name = 'CentOS' AND version = '7' LIMIT 1)),
    ('app-server-01', (SELECT id FROM operating_systems WHERE name = 'RedHat' AND version = '7' LIMIT 1)),
    ('cache-server-01', (SELECT id FROM operating_systems WHERE name = 'Debian' AND version = '12' LIMIT 1))
ON CONFLICT (name) DO NOTHING;

-- Create a trigger to automatically update the updated_at column
CREATE TRIGGER update_servers_updated_at
    BEFORE UPDATE ON servers
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();
