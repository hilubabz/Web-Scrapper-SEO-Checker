-- +goose Up

CREATE TABLE projects (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name TEXT NOT NULL,
    base_url TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE audits (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    project_id BIGINT NOT NULL REFERENCES projects(id) ON DELETE CASCADE,

    started_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    completed_at TIMESTAMPTZ,

    status TEXT NOT NULL DEFAULT 'running',
    score INTEGER,
    pages_crawled INTEGER NOT NULL DEFAULT 0,

    CONSTRAINT audits_status_check
        CHECK (status IN ('running', 'completed', 'failed')),

    CONSTRAINT audits_score_check
        CHECK (score IS NULL OR (score >= 0 AND score <= 100))
);

CREATE TABLE pages (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    audit_id BIGINT NOT NULL REFERENCES audits(id) ON DELETE CASCADE,

    url TEXT NOT NULL,
    depth INTEGER NOT NULL DEFAULT 0,

    status_code INTEGER,

    title TEXT,
    title_length INTEGER NOT NULL DEFAULT 0,

    meta_description TEXT,
    meta_description_length INTEGER NOT NULL DEFAULT 0,

    h1_count INTEGER NOT NULL DEFAULT 0,
    h2_count INTEGER NOT NULL DEFAULT 0,

    image_count INTEGER NOT NULL DEFAULT 0,
    images_without_alt INTEGER NOT NULL DEFAULT 0,

    link_count INTEGER NOT NULL DEFAULT 0,
    internal_link_count INTEGER NOT NULL DEFAULT 0,
    external_link_count INTEGER NOT NULL DEFAULT 0,

    canonical_url TEXT,
    has_canonical BOOLEAN NOT NULL DEFAULT FALSE,

    robots_meta TEXT,
    has_robots_meta BOOLEAN NOT NULL DEFAULT FALSE,

    score INTEGER,

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    CONSTRAINT pages_score_check
        CHECK (score IS NULL OR (score >= 0 AND score <= 100)),

    CONSTRAINT pages_unique_url_per_audit
        UNIQUE (audit_id, url)
);

CREATE TABLE seo_issues (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    page_id BIGINT NOT NULL REFERENCES pages(id) ON DELETE CASCADE,

    rule TEXT NOT NULL,
    severity TEXT NOT NULL,
    message TEXT NOT NULL,

    CONSTRAINT seo_issues_severity_check
        CHECK (severity IN ('error', 'warning', 'info'))
);

CREATE INDEX idx_audits_project_id
    ON audits(project_id);

CREATE INDEX idx_pages_audit_id
    ON pages(audit_id);

CREATE INDEX idx_seo_issues_page_id
    ON seo_issues(page_id);

CREATE INDEX idx_pages_url
    ON pages(url);


-- +goose Down

DROP TABLE IF EXISTS seo_issues;
DROP TABLE IF EXISTS pages;
DROP TABLE IF EXISTS audits;
DROP TABLE IF EXISTS projects;