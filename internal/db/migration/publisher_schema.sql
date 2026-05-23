CREATE TABLE publishers (
  id   BIGSERIAL PRIMARY KEY,
  name text      NOT NULL UNIQUE
);