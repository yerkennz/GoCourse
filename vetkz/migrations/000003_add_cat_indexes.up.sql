CREATE INDEX IF NOT EXISTS cat_title_idx ON cats USING GIN (to_tsvector('simple', title));

