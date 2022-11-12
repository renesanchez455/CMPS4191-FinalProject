-- Filename: migrations/000004_create_forums_likes_table.up.sql

CREATE TABLE IF NOT EXISTS forumslikes (
    id bigserial PRIMARY KEY,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    users_id BIGINT REFERENCES users (id),
    forums_id BIGINT REFERENCES forums (id),
    UNIQUE(users_id, forums_id)
);
