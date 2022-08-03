-- name: GetPostById :one
SELECT * FROM posts
WHERE id = $1 LIMIT 1;

-- name: ListPosts :many
SELECT * FROM posts
ORDER BY title
LIMIT $1
OFFSET $2;

-- name: CreatePost :one
INSERT INTO posts (
  title,
  description,
  author_id
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: DeletePost :exec
DELETE FROM posts
WHERE id = $1;

-- name: UpdatePost :one
UPDATE posts
set title = $2,
description = $3
WHERE id = $1
RETURNING *;

-- name: GetAllPostsForAuthor :many
SELECT * FROM posts
WHERE author_id = $1
LIMIT $2
OFFSET $3;