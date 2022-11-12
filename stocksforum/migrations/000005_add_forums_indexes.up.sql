-- Filename: migrations/000005_add_forums_indexes.up.sql
CREATE INDEX IF NOT EXISTS forums_name_idx ON forums USING GIN(to_tsvector('simple', name));