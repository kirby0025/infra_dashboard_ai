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
    ('CentOS','8','2021-12-31'),
    ('RedHat','5','2017-03-31'),
    ('RedHat','6','2020-11-30'),
    ('RedHat','7','2024-06-30'),
    ('RedHat','8','2029-05-31'),
    ('RedHat','9','2032-05-31'),
    ('RedHat','10','2035-05-31'),
    ('FreeBSD','9.0','2013-03-31'),
    ('FreeBSD','9.1','2014-12-31'),
    ('FreeBSD','9.2','2014-12-31'),
    ('FreeBSD','9.3','2016-12-31'),
    ('FreeBSD','10.0','2015-02-28'),
    ('FreeBSD','10.1','2016-12-31'),
    ('FreeBSD','10.2','2016-12-31'),
    ('FreeBSD','10.3','2018-04-30'),
    ('FreeBSD','10.4','2018-10-31'),
    ('FreeBSD','11.0','2017-11-30'),
    ('FreeBSD','11.1','2018-09-30'),
    ('FreeBSD','11.2','2019-10-31'),
    ('FreeBSD','11.3','2020-09-30'),
    ('FreeBSD','11.4','2021-09-30'),
    ('FreeBSD','12.0','2020-02-04'),
    ('FreeBSD','12.1','2021-01-31'),
    ('FreeBSD','12.2','2022-03-31'),
    ('FreeBSD','12.3','2023-03-31'),
    ('FreeBSD','12.4','2023-12-31'),
    ('FreeBSD','13.0','2022-08-31'),
    ('FreeBSD','13.1','2023-07-31'),
    ('FreeBSD','13.2','2024-06-30'),
    ('FreeBSD','13.3','2024-12-31'),
    ('FreeBSD','13.4','2025-06-30'),
    ('FreeBSD','13.5','2026-04-30'),
    ('FreeBSD','14.0','2024-09-30'),
    ('FreeBSD','14.1','2025-03-31'),
    ('FreeBSD','14.2','2025-09-30'),
    ('FreeBSD','14.3','2026-06-30'),
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

-- Create indexes for better query performance on change history
CREATE INDEX IF NOT EXISTS idx_change_history_server_id ON server_change_history(server_id);
CREATE INDEX IF NOT EXISTS idx_change_history_change_type ON server_change_history(change_type);
CREATE INDEX IF NOT EXISTS idx_change_history_changed_at ON server_change_history(changed_at);

-- Create a function to log server creation
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

-- Create a function to log OS changes
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

-- Create a function to log server deletion
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
