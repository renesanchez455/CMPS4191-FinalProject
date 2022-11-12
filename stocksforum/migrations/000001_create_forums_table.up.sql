-- Filename: migrations/000001_create_forums_table.up.sql

CREATE TABLE IF NOT EXISTS forums (
    id bigserial PRIMARY KEY,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    name text NOT NULL,
    message text NOT NULL,
    version integer NOT NULL DEFAULT 1
);