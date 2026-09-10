-- name: CreateAudit :one
INSERT INTO audits (
    project_id,
    status
)
VALUES (
    $1,
    'running'
)
RETURNING *;


-- name: CompleteAudit :one
UPDATE audits
SET
    completed_at = NOW(),
    status = $2,
    score = $3,
    pages_crawled = $4
WHERE id = $1
RETURNING *;


-- name: GetAudit :one
SELECT *
FROM audits
WHERE id = $1;


-- name: ListProjectAudits :many
SELECT *
FROM audits
WHERE project_id = $1
ORDER BY started_at DESC;