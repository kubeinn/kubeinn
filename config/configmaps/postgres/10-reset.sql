-- Reset schema and tables
DROP SCHEMA IF EXISTS api CASCADE;
-- Drop all sequences
DROP SEQUENCE IF EXISTS pilgrim_sequence;
DROP SEQUENCE IF EXISTS village_sequence;
DROP SEQUENCE IF EXISTS innkeeper_sequence;
DROP SEQUENCE IF EXISTS reeve_sequence;
-- Drop all group roles
DROP ROLE IF EXISTS pilgrims;
DROP ROLE IF EXISTS villages;
DROP ROLE IF EXISTS innkeepers;
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
-- Drop all village roles
DO $do$ 
DECLARE temprow record;
BEGIN FOR temprow IN
SELECT rolname
FROM pg_roles
WHERE rolname SIMILAR TO 'village-[0-9]*' LOOP 
EXECUTE FORMAT('DROP ROLE "%s";', temprow.rolname);
END LOOP;
END $do$;
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
-- Drop all reeve roles
DO $do$ 
DECLARE temprow record;
BEGIN FOR temprow IN
SELECT rolname
FROM pg_roles
WHERE rolname SIMILAR TO 'reeve-[0-9]*' LOOP 
EXECUTE FORMAT('DROP ROLE "%s";', temprow.rolname);
END LOOP;
END $do$;