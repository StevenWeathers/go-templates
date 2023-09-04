-- name: GetAuthor :one
SELECT * FROM example.authors WHERE id = $1 LIMIT 1;

-- name: ListAuthors :many
SELECT * FROM example.authors ORDER BY name;

-- name: CreateAuthor :one
INSERT INTO example.authors (name, bio) VALUES ($1, $2) RETURNING *;

-- name: DeleteAuthor :exec
DELETE FROM example.authors WHERE id = $1;