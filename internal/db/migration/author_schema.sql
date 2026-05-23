CREATE TABLE authors (
  id           BIGSERIAL PRIMARY KEY,
  publisher_id BIGINT    NOT NULL REFERENCES publishers(id),
  name         text      NOT NULL,
  bio          text
);