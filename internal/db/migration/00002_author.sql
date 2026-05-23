-- +goose Up
CREATE TABLE author (
  id           BIGSERIAL PRIMARY KEY,
  publisher_id BIGINT    NOT NULL REFERENCES publisher(id),
  name         text      NOT NULL,
  bio          text
);

-- +goose Down
DROP TABLE author;