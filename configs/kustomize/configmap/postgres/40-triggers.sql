-- Function to create a role for each new pilgrim
CREATE OR REPLACE FUNCTION create_pilgrim_role_function() RETURNS TRIGGER 
AS $$ 
BEGIN 
    EXECUTE FORMAT('CREATE ROLE "%s";', NEW.id);
    EXECUTE FORMAT('GRANT USAGE ON SCHEMA api TO "%s";',NEW.id);
    EXECUTE FORMAT('GRANT ALL PRIVILEGES ON api.projects TO "%s";', NEW.id);
    RETURN NEW;
END;
$$
LANGUAGE PLPGSQL;
-- Function to create a role for each new village
CREATE OR REPLACE FUNCTION create_village_role_function() RETURNS TRIGGER 
AS $$ 
BEGIN 
    EXECUTE FORMAT('CREATE ROLE "%s";', NEW.id);
    EXECUTE FORMAT('GRANT USAGE ON SCHEMA api TO "%s";',NEW.id);
    EXECUTE FORMAT('GRANT ALL PRIVILEGES ON api.pilgrims TO "%s";', NEW.id);
    EXECUTE FORMAT('GRANT ALL PRIVILEGES ON api.projects TO "%s";', NEW.id);
    EXECUTE FORMAT('GRANT ALL PRIVILEGES ON api.tickets TO "%s";', NEW.id);
    EXECUTE FORMAT('GRANT ALL PRIVILEGES ON api.usage TO "%s";', NEW.id);
    RETURN NEW;
END;
$$
LANGUAGE PLPGSQL;
-- Function to update villageID for each new project
CREATE OR REPLACE FUNCTION update_village_id_function() RETURNS TRIGGER 
AS $$ 
BEGIN 
    NEW.villageID := (SELECT villageID FROM api.pilgrims WHERE id = NEW.pilgrimID);
    RETURN NEW;
END;
$$
LANGUAGE PLPGSQL;
-- Function to update vic for each new village
CREATE OR REPLACE FUNCTION update_village_vic_function() RETURNS TRIGGER 
AS $$ 
BEGIN 
    NEW.vic := (SELECT md5(random()::text));
    RETURN NEW;
END;
$$
LANGUAGE PLPGSQL;
-- Trigger to call create_pilgrim_role_function
DROP TRIGGER IF EXISTS after_village_created ON api.villages;
CREATE TRIGGER after_village_created AFTER
INSERT ON api.villages FOR EACH ROW EXECUTE FUNCTION create_village_role_function();
-- Trigger to call update_village_vic_function
DROP TRIGGER IF EXISTS before_village_created ON api.projects;
CREATE TRIGGER before_village_created BEFORE
INSERT ON api.villages FOR EACH ROW EXECUTE FUNCTION update_village_vic_function();
-- Trigger to call create_village_role_function
DROP TRIGGER IF EXISTS after_pilgrim_created ON api.pilgrims;
CREATE TRIGGER after_pilgrim_created AFTER
INSERT ON api.pilgrims FOR EACH ROW EXECUTE FUNCTION create_pilgrim_role_function();
-- Trigger to call update_village_id_function
DROP TRIGGER IF EXISTS before_project_created ON api.projects;
CREATE TRIGGER before_project_created BEFORE
INSERT ON api.projects FOR EACH ROW EXECUTE FUNCTION update_village_id_function();
