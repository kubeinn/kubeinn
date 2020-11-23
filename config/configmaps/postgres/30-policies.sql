-- Enable row security policies
ALTER TABLE api.villages ENABLE ROW LEVEL SECURITY;
ALTER TABLE api.pilgrims ENABLE ROW LEVEL SECURITY;
ALTER TABLE api.innkeepers ENABLE ROW LEVEL SECURITY;
ALTER TABLE api.projects ENABLE ROW LEVEL SECURITY;
ALTER TABLE api.tickets ENABLE ROW LEVEL SECURITY;
ALTER TABLE api.usage ENABLE ROW LEVEL SECURITY;
-- Set row security policies
CREATE POLICY pilgrim_pilgrim_policy ON api.pilgrims FOR ALL USING (id = current_user);
CREATE POLICY project_pilgrim_policy ON api.projects FOR ALL USING (pilgrimid = current_user);
CREATE POLICY usage_pilgrim_policy ON api.usage FOR ALL USING (pilgrimid = current_user);
CREATE POLICY ticket_village_policy ON api.tickets FOR ALL USING (villageid = current_user);
CREATE POLICY project_village_policy ON api.projects FOR ALL USING (villageid = current_user);
CREATE POLICY pilgrim_village_policy ON api.pilgrims FOR ALL USING (villageid = current_user);
-- Grant permission for schema and sequences to group roles
GRANT USAGE ON SCHEMA api TO pilgrims;
GRANT USAGE ON SCHEMA api TO villages;
GRANT USAGE ON ALL SEQUENCES IN SCHEMA api TO pilgrims;
GRANT USAGE ON ALL SEQUENCES IN SCHEMA api TO villages;
-- Grant permissions to role groups
GRANT ALL PRIVILEGES ON api.tickets TO villages;
GRANT ALL PRIVILEGES ON api.usage TO villages;
GRANT ALL PRIVILEGES ON api.projects TO villages; 
GRANT ALL PRIVILEGES ON api.pilgrims TO villages; 
GRANT ALL PRIVILEGES ON api.projects TO pilgrims; 
GRANT ALL PRIVILEGES ON api.pilgrims TO pilgrims; 
