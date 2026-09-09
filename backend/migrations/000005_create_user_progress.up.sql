CREATE TABLE IF NOT EXISTS user_progress (
    id           UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id      UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    lesson_id    UUID NOT NULL REFERENCES lessons(id) ON DELETE CASCADE,
    status       VARCHAR(20)  NOT NULL DEFAULT 'not_started', -- not_started | in_progress | completed
    completed_at TIMESTAMPTZ,
    created_at   TIMESTAMPTZ  NOT NULL DEFAULT now(),
    updated_at   TIMESTAMPTZ  NOT NULL DEFAULT now(),
    UNIQUE (user_id, lesson_id)
);
CREATE INDEX IF NOT EXISTS idx_user_progress_user_id ON user_progress (user_id);