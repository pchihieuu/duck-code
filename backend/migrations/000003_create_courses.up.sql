CREATE TABLE IF NOT EXISTS courses (
    id           UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title        VARCHAR(150) NOT NULL,
    slug         VARCHAR(150) NOT NULL UNIQUE,
    description  TEXT,
    language     VARCHAR(30)  NOT NULL, -- 'python' | 'html-css-js' ...
    order_index  INT          NOT NULL DEFAULT 0,
    is_published BOOLEAN      NOT NULL DEFAULT false,
    created_at   TIMESTAMPTZ  NOT NULL DEFAULT now(),
    updated_at   TIMESTAMPTZ  NOT NULL DEFAULT now(),
    deleted_at   TIMESTAMPTZ
);
CREATE INDEX IF NOT EXISTS idx_courses_deleted_at ON courses (deleted_at);