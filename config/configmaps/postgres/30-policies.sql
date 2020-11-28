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
-- Grant permissions to role groups innkeepers
GRANT ALL PRIVILEGES ON api.innkeepers TO innkeepers;
GRANT ALL PRIVILEGES ON api.pilgrims TO innkeepers;
GRANT ALL PRIVILEGES ON api.projects TO innkeepers;
GRANT ALL PRIVILEGES ON api.tickets TO innkeepers;
GRANT ALL PRIVILEGES ON api.usage TO innkeepers;
-- Grant permissions to role groups pilgrims
GRANT ALL PRIVILEGES ON api.tickets TO pilgrims;
GRANT ALL PRIVILEGES ON api.usage TO pilgrims;
GRANT ALL PRIVILEGES ON api.projects TO pilgrims;