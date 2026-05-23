CREATE TABLE books
(
    id           BIGSERIAL PRIMARY KEY,
    author_id    BIGINT NOT NULL REFERENCES authors (id) ON DELETE CASCADE,
    title        text   NOT NULL,
    summary      text,
    published_at date
);
