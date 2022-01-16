
-- +migrate Up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- +migrate Down
DELETE EXTENSION IF EXISTS "uuid-ossp";