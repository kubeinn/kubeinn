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
-- Grant permissions on sequences
GRANT USAGE, SELECT ON ALL SEQUENCES IN SCHEMA api TO PUBLIC; 
-- Grant permissions to schema and publically accessible tables
GRANT USAGE ON SCHEMA api TO PUBLIC; 
GRANT ALL PRIVILEGES ON api.projects TO PUBLIC; 
GRANT ALL PRIVILEGES ON api.pilgrims TO PUBLIC; 