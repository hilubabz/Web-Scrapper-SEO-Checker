-- name: CreatePage :one
INSERT INTO pages (
    audit_id,
    url,
    depth,
    status_code,
    title,
    title_length,
    meta_description,
    meta_description_length,
    h1_count,
    h2_count,
    image_count,
    images_without_alt,
    link_count,
    internal_link_count,
    external_link_count,
    canonical_url,
    has_canonical,
    robots_meta,
    has_robots_meta,
    score
)
VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10,
    $11, $12, $13, $14, $15, $16, $17, $18, $19, $20
)
RETURNING *;


-- name: GetAuditPages :many
SELECT *
FROM pages
WHERE audit_id = $1
ORDER BY id;


-- name: GetPage :one
SELECT *
FROM pages
WHERE id = $1;