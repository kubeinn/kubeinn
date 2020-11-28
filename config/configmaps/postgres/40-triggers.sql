-- Function to create a role for each new pilgrim
CREATE OR REPLACE FUNCTION create_pilgrim_role_function() RETURNS TRIGGER 
AS $$ 
BEGIN 
    EXECUTE FORMAT('CREATE ROLE "%s" CREATEROLE;', NEW.id);
    EXECUTE FORMAT('GRANT pilgrims TO "%s";',NEW.id);
    EXECUTE FORMAT('GRANT "%s" TO postgrest;',NEW.id);
    RETURN NEW;
END;
$$
LANGUAGE PLPGSQL;

-- Function to create a role for each new innkeeper
CREATE OR REPLACE FUNCTION create_innkeeper_role_function() RETURNS TRIGGER 
AS $$ 
BEGIN 
    EXECUTE FORMAT('CREATE ROLE "%s" CREATEROLE;', NEW.id);
    EXECUTE FORMAT('GRANT innkeepers TO "%s";',NEW.id);
    EXECUTE FORMAT('GRANT "%s" TO postgrest;',NEW.id);
    RETURN NEW;
END;
$$
LANGUAGE PLPGSQL;

-- Trigger to call create_innkeeper_role_function
DROP TRIGGER IF EXISTS after_pilgrim_created ON api.pilgrims;
CREATE TRIGGER after_pilgrim_created AFTER
INSERT ON api.pilgrims FOR EACH ROW EXECUTE FUNCTION create_pilgrim_role_function();

-- Trigger to call create_innkeeper_role_function
DROP TRIGGER IF EXISTS after_innkeeper_created ON api.innkeepers;
CREATE TRIGGER after_innkeeper_created AFTER
INSERT ON api.innkeepers FOR EACH ROW EXECUTE FUNCTION create_innkeeper_role_function();