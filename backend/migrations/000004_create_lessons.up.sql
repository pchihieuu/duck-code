CREATE TABLE IF NOT EXISTS lessons (
    id           UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    course_id    UUID NOT NULL REFERENCES courses(id) ON DELETE CASCADE,
    title        VARCHAR(150) NOT NULL,
    slug         VARCHAR(150) NOT NULL,
    content_mdx  TEXT         NOT NULL DEFAULT '', -- admin-authored qua dashboard, xem §0.1
    order_index  INT          NOT NULL DEFAULT 0,
    is_published BOOLEAN      NOT NULL DEFAULT false,
    created_at   TIMESTAMPTZ  NOT NULL DEFAULT now(),
    updated_at   TIMESTAMPTZ  NOT NULL DEFAULT now(),
    deleted_at   TIMESTAMPTZ,
    UNIQUE (course_id, slug)
);
CREATE INDEX IF NOT EXISTS idx_lessons_course_id ON lessons (course_id);
CREATE INDEX IF NOT EXISTS idx_lessons_deleted_at ON lessons (deleted_at);