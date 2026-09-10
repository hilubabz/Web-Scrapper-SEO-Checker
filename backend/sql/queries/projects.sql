-- name: CreateProject :one
INSERT INTO projects (
    name,
    base_url
)
VALUES (
    $1,
    $2
)
RETURNING *;


-- name: GetProject :one
SELECT *
FROM projects
WHERE id = $1;


-- name: ListProjects :many
SELECT *
FROM projects
ORDER BY created_at DESC;


-- name: DeleteProject :exec
DELETE FROM projects
WHERE id = $1;