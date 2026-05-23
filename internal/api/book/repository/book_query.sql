-- name: GetBook :one
SELECT * FROM books
WHERE id = $1 LIMIT 1;

-- name: ListBooks :many
SELECT * FROM books
ORDER BY title;

-- name: ListBooksByAuthor :many
SELECT * FROM books
WHERE author_id = $1
ORDER BY title;

-- name: CreateBook :one
INSERT INTO books (
	author_id, title, summary, published_at
) VALUES (
	$1, $2, $3, $4
)
RETURNING *;

-- name: UpdateBook :one
UPDATE books
SET author_id = $2,
		title = $3,
		summary = $4,
		published_at = $5
WHERE id = $1
RETURNING *;

-- name: DeleteBook :exec
DELETE FROM books
WHERE id = $1;
