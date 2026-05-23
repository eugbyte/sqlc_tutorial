-- +goose Up
CREATE TABLE publisher (
  id   BIGSERIAL PRIMARY KEY,
  name text      NOT NULL UNIQUE
);

-- +goose Down
DROP TABLE publisher;