-- Function to create a role for each new pilgrim
CREATE OR REPLACE FUNCTION create_pilgrim_role_function() RETURNS TRIGGER 
AS $$ 
BEGIN 
    EXECUTE FORMAT('CREATE ROLE "%s";', NEW.id);
    EXECUTE FORMAT('GRANT pilgrims TO "%s";',NEW.id);
    RETURN NEW;
END;
$$
LANGUAGE PLPGSQL;
-- Function to create a role for each new village
CREATE OR REPLACE FUNCTION create_village_role_function() RETURNS TRIGGER 
AS $$ 
BEGIN 
    EXECUTE FORMAT('CREATE ROLE "%s" CREATEROLE;', NEW.id);
    EXECUTE FORMAT('GRANT villages TO "%s";',NEW.id);
    RETURN NEW;
END;
$$
LANGUAGE PLPGSQL;

-- Function to update villageid for each new project
CREATE OR REPLACE FUNCTION update_village_id_function() RETURNS TRIGGER 
AS $$ 
BEGIN 
    NEW.villageid := (SELECT villageid FROM api.pilgrims WHERE id = NEW.pilgrimid);
    RETURN NEW;
END;
$$
LANGUAGE PLPGSQL;
-- Function to update vic for each new village
CREATE OR REPLACE FUNCTION update_pilgrim_function() RETURNS TRIGGER 
AS $$ 
BEGIN 
    NEW.villageid := current_user;
    NEW.passwd := (SELECT md5(random()::text));
    NEW.regcode := (SELECT md5(random()::text));
    RETURN NEW;
END;
$$
LANGUAGE PLPGSQL;
-- Trigger to call create_village_role_function
DROP TRIGGER IF EXISTS after_village_created ON api.villages;
CREATE TRIGGER after_village_created AFTER
INSERT ON api.villages FOR EACH ROW EXECUTE FUNCTION create_village_role_function();
-- Trigger to call update_village_vic_function
DROP TRIGGER IF EXISTS before_pilgrim_created ON api.pilgrims;
CREATE TRIGGER before_pilgrim_created BEFORE
INSERT ON api.pilgrims FOR EACH ROW EXECUTE FUNCTION update_pilgrim_function();
-- Trigger to call create_village_role_function
DROP TRIGGER IF EXISTS after_pilgrim_created ON api.pilgrims;
CREATE TRIGGER after_pilgrim_created AFTER
INSERT ON api.pilgrims FOR EACH ROW EXECUTE FUNCTION create_pilgrim_role_function();
-- Trigger to call update_village_id_function
DROP TRIGGER IF EXISTS before_project_created ON api.projects;
CREATE TRIGGER before_project_created BEFORE 
INSERT OR UPDATE ON api.projects FOR EACH ROW EXECUTE FUNCTION update_village_id_function();
