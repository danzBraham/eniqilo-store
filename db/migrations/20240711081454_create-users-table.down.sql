BEGIN;

DROP INDEX IF EXISTS idx_users_phone_number;
DROP INDEX IF EXISTS idx_users_role;
DROP INDEX IF EXISTS idx_users_name;

DROP TABLE IF EXISTS users;

DROP TYPE IF EXISTS user_role;

COMMIT;