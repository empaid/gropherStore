CREATE EXTENSION IF NOT EXISTS citext;

CREATE TABLE IF NOT EXISTS users (
    id         BIGSERIAL             PRIMARY KEY,
    email      CITEXT                UNIQUE NOT NULL,
    username   VARCHAR(255)          UNIQUE NOT NULL,  -- ← comma here
    password   BYTEA                 NOT NULL,         -- ← and here
    created_at TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT NOW()
);