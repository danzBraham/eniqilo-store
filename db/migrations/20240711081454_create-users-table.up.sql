BEGIN;

CREATE TYPE user_role AS ENUM ('Staff', 'Customer');

CREATE TABLE IF NOT EXISTS users (
  id VARCHAR(26) PRIMARY KEY NOT NULL,
  phone_number VARCHAR(16) NOT NULL,
  name VARCHAR(50) NOT NULL,
  password VARCHAR(60) NOT NULL,
  role user_role NOT NULL,
  is_deleted BOOLEAN NOT NULL DEFAULT false,
  created_at TIMESTAMP DEFAULT NOW(),
  updated_at TIMESTAMP DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_users_phone_number ON users(phone_number);
CREATE INDEX IF NOT EXISTS idx_users_role ON users(role);
CREATE INDEX IF NOT EXISTS idx_users_name ON users(lower(name));

COMMIT;