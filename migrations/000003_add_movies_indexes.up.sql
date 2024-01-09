CREATE INDEX IF NOT EXISTS movies_title_idx ON movies USING GIN (title gin_trgm_ops);
CREATE INDEX IF NOT EXISTS movies_genres_idx ON movies USING GIN (genres);
