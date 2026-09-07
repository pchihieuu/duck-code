CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE TABLE IF NOT EXISTS users (
    id            UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    email         VARCHAR(255) NOT NULL UNIQUE,
    username      VARCHAR(30)  NOT NULL UNIQUE,
    password_hash TEXT         NOT NULL,
    display_name  VARCHAR(80)  NOT NULL,
    avatar_url    TEXT,
    role          VARCHAR(20)  NOT NULL DEFAULT 'student',
    created_at    TIMESTAMPTZ  NOT NULL DEFAULT now(),
    updated_at    TIMESTAMPTZ  NOT NULL DEFAULT now(),
    deleted_at    TIMESTAMPTZ
);

CREATE INDEX IF NOT EXISTS idx_users_deleted_at ON users (deleted_at);
