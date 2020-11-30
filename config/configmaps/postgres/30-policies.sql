-- Enable row security policies
ALTER TABLE api.innkeepers ENABLE ROW LEVEL SECURITY;
ALTER TABLE api.pilgrims ENABLE ROW LEVEL SECURITY;
ALTER TABLE api.projects ENABLE ROW LEVEL SECURITY;
ALTER TABLE api.tickets ENABLE ROW LEVEL SECURITY;
ALTER TABLE api.usage ENABLE ROW LEVEL SECURITY;
-- Set row security policies for innkeepers
CREATE POLICY innkeeper_innkeeper_policy ON api.innkeepers TO innkeepers USING (true) WITH CHECK (true);
CREATE POLICY innkeeper_pilgrim_policy ON api.pilgrims TO innkeepers USING (true) WITH CHECK (true);
CREATE POLICY innkeeper_projects_policy ON api.projects TO innkeepers USING (true) WITH CHECK (true);
CREATE POLICY innkeeper_tickets_policy ON api.tickets TO innkeepers USING (true) WITH CHECK (true);
CREATE POLICY innkeeper_usage_policy ON api.usage TO innkeepers USING (true) WITH CHECK (true);
-- Set row security policies for pilgrims
CREATE POLICY project_pilgrim_policy ON api.projects FOR ALL USING (pilgrimid = current_user);
CREATE POLICY usage_pilgrim_policy ON api.usage FOR ALL USING (pilgrimid = current_user);
CREATE POLICY ticket_pilgrim_policy ON api.tickets FOR ALL USING (pilgrimid = current_user);
-- Grant permission for schema and sequences to group roles
GRANT USAGE ON SCHEMA api TO pilgrims;
GRANT USAGE ON ALL SEQUENCES IN SCHEMA api TO pilgrims;
GRANT USAGE ON SCHEMA api TO innkeepers;
GRANT USAGE ON ALL SEQUENCES IN SCHEMA api TO innkeepers;
-- Create stored procedures
CREATE OR REPLACE FUNCTION api.update_innkeeper(id TEXT, username TEXT, email TEXT, passwd TEXT) RETURNS VOID AS $$ BEGIN IF (passwd = '') IS NOT FALSE THEN EXECUTE FORMAT(
        'UPDATE api.innkeepers SET username=''%s'', email=''%s''s WHERE id=''%s'';',
        username,
        email,
        id
    );
    ELSE EXECUTE FORMAT(
        'UPDATE api.innkeepers SET username=''%s'', email=''%s'', passwd=''%s'' WHERE id=''%s'';',
        username,
        email,
        passwd,
        id
    );
END IF;
END;
$$ LANGUAGE PLPGSQL;
CREATE OR REPLACE FUNCTION api.create_innkeeper(username TEXT, email TEXT, passwd TEXT) RETURNS VOID AS $$ BEGIN EXECUTE FORMAT(
        'INSERT INTO api.innkeepers (username,email,passwd) VALUES (''%s'',''%s'',''%s'');',
        username,
        email,
        passwd
    );
END;
$$ LANGUAGE PLPGSQL;
CREATE OR REPLACE FUNCTION api.update_pilgrim (
        id TEXT,
        organization TEXT,
        description TEXT,
        username TEXT,
        email TEXT,
        passwd TEXT,
        status TEXT
    ) RETURNS VOID AS $$ BEGIN IF (passwd = '') IS NOT FALSE THEN EXECUTE FORMAT(
        'UPDATE api.pilgrims SET organization=''%s'', description=''%s'', username=''%s'', email=''%s'', status=''%s'' WHERE id=''%s'';',
        organization,
        description,
        username,
        email,
        status,
        id
    );
    ELSE
    EXECUTE FORMAT(
        'UPDATE api.pilgrims SET organization=''%s'', description=''%s'', username=''%s'', email=''%s'', passwd=''%s'', status=''%s'' WHERE id=''%s'';',
        organization,
        description,
        username,
        email,
        passwd,
        status,
        id
    );
END IF;
END;
$$ LANGUAGE PLPGSQL;
CREATE OR REPLACE FUNCTION api.create_pilgrim(
        organization TEXT,
        description TEXT,
        username TEXT,
        email TEXT,
        passwd TEXT,
        status TEXT
    ) RETURNS VOID AS $$ BEGIN EXECUTE FORMAT(
        'INSERT INTO api.pilgrims (organization,description,username,email,passwd,status) VALUES (''%s'',''%s'',''%s'',''%s'',''%s'',''%s'');',
        organization,
        description,
        username,
        email,
        passwd,
        status
    );
END;
$$ LANGUAGE PLPGSQL;
-- Grant permissions to role groups innkeepers
GRANT SELECT (id, username, email) ON api.innkeepers TO innkeepers;
GRANT SELECT (
        id,
        organization,
        description,
        username,
        email,
        status
    ) ON api.pilgrims TO innkeepers;
GRANT INSERT,
    UPDATE,
    DELETE ON api.innkeepers TO innkeepers;
GRANT INSERT,
    UPDATE,
    DELETE ON api.pilgrims TO innkeepers;
GRANT ALL PRIVILEGES ON api.projects TO innkeepers;
GRANT ALL PRIVILEGES ON api.tickets TO innkeepers;
GRANT ALL PRIVILEGES ON api.usage TO innkeepers;
-- Grant permissions to role groups pilgrims
GRANT ALL PRIVILEGES ON api.tickets TO pilgrims;
GRANT ALL PRIVILEGES ON api.usage TO pilgrims;
GRANT ALL PRIVILEGES ON api.projects TO pilgrims;
GRANT EXECUTE ON FUNCTION update_pilgrim TO pilgrims;
GRANT EXECUTE ON FUNCTION create_pilgrim TO pilgrims;