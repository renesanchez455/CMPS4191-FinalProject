-- Filename: migrations/000003_create_replies_table.up.sql

CREATE TABLE IF NOT EXISTS replies (
    id bigserial PRIMARY KEY,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    message text NOT NULL,
    version integer NOT NULL DEFAULT 1,
    users_id BIGINT REFERENCES users (id),
    forums_id BIGINT REFERENCES forums (id),
    UNIQUE(users_id, forums_id)
);
