-- name: GetAuthor :one
SELECT sqlc.embed(authors), sqlc.embed(publishers)
FROM authors
JOIN publishers ON authors.publisher_id = publishers.id
WHERE authors.id = $1 LIMIT 1;

-- name: ListAuthors :many
SELECT sqlc.embed(authors), sqlc.embed(publishers)
FROM authors
JOIN publishers ON authors.publisher_id = publishers.id
ORDER BY authors.name;

-- name: CreateAuthor :one
INSERT INTO authors (
  name, bio
) VALUES (
  $1, $2
)
RETURNING *;

-- name: UpdateAuthor :one
UPDATE authors
  set name = $2,
  bio = $3
WHERE id = $1
RETURNING *;

-- name: DeleteAuthor :exec
DELETE FROM authors
WHERE id = $1;