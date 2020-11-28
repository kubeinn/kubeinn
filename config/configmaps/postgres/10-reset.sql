-- Reset schema and tables
DROP SCHEMA IF EXISTS api CASCADE;
-- Drop all sequences
DROP SEQUENCE IF EXISTS innkeeper_sequence;
DROP SEQUENCE IF EXISTS pilgrim_sequence;
-- Drop all group roles
DROP ROLE IF EXISTS postgrest;
DROP ROLE IF EXISTS innkeepers;
DROP ROLE IF EXISTS pilgrims;
-- Drop all innkeeper roles
DO $do$ 
DECLARE temprow record;
BEGIN FOR temprow IN
SELECT rolname
FROM pg_roles
WHERE rolname SIMILAR TO 'innkeeper-[0-9]*' LOOP 
EXECUTE FORMAT('DROP ROLE "%s";', temprow.rolname);
END LOOP;
END $do$;
-- Drop all pilgrim roles
DO $do$ 
DECLARE temprow record;
BEGIN FOR temprow IN
SELECT rolname
FROM pg_roles
WHERE rolname SIMILAR TO 'pilgrim-[0-9]*' LOOP 
EXECUTE FORMAT('DROP ROLE "%s";', temprow.rolname);
END LOOP;
END $do$;