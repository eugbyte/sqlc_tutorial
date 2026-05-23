-- name: GetAuthor :one
SELECT sqlc.embed(author), sqlc.embed(publisher)
FROM author
JOIN publisher ON author.publisher_id = publisher.id
WHERE author.id = $1 LIMIT 1;

-- name: ListAuthors :many
SELECT sqlc.embed(author), sqlc.embed(publisher)
FROM author
JOIN publisher ON author.publisher_id = publisher.id
ORDER BY author.name;

-- name: CreateAuthor :one
INSERT INTO author (
  name, bio
) VALUES (
  $1, $2
)
RETURNING *;

-- name: UpdateAuthor :one
UPDATE author
  set name = $2,
  bio = $3
WHERE id = $1
RETURNING *;

-- name: DeleteAuthor :exec
DELETE FROM author
WHERE id = $1;