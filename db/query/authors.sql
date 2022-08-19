-- name: GetAuthorById :one
SELECT * FROM authors
WHERE id = $1 LIMIT 1;

-- name: ListAuthors :many
SELECT * FROM authors
ORDER BY name
LIMIT $1
OFFSET $2;

-- name: CreateAuthor :one
INSERT INTO authors (
  name
) VALUES (
  $1
)
RETURNING *;

-- name: DeleteAuthor :exec
DELETE FROM authors
WHERE id = $1;

-- name: UpdateAuthor :one
UPDATE authors
set name = $2
WHERE id = $1
RETURNING *;