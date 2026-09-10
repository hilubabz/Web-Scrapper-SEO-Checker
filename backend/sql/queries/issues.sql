-- name: CreateSEOIssue :one
INSERT INTO seo_issues (
    page_id,
    rule,
    severity,
    message
)
VALUES (
    $1,
    $2,
    $3,
    $4
)
RETURNING *;


-- name: GetPageIssues :many
SELECT *
FROM seo_issues
WHERE page_id = $1
ORDER BY id;