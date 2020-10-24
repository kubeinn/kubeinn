-- Enable row security policies
ALTER TABLE api.villages ENABLE ROW LEVEL SECURITY;
ALTER TABLE api.pilgrims ENABLE ROW LEVEL SECURITY;
ALTER TABLE api.innkeepers ENABLE ROW LEVEL SECURITY;
ALTER TABLE api.projects ENABLE ROW LEVEL SECURITY;
ALTER TABLE api.tickets ENABLE ROW LEVEL SECURITY;
ALTER TABLE api.usage ENABLE ROW LEVEL SECURITY;
-- Set row security policies
CREATE POLICY project_pilgrim_policy ON api.projects USING (pilgrimID = current_user);
CREATE POLICY usage_pilgrim_policy ON api.usage USING (pilgrimID = current_user);
CREATE POLICY ticket_village_policy ON api.tickets USING (villageID = current_user);
CREATE POLICY project_village_policy ON api.projects USING (villageID = current_user);
CREATE POLICY pilgrim_village_policy ON api.pilgrims USING (villageID = current_user);
