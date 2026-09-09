CREATE TABLE IF NOT EXISTS activity_log (
    id           UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id      UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    type         VARCHAR(30) NOT NULL, -- 'lesson_completed' | 'exercise_passed' | 'quiz_passed' ...
    reference_id UUID,                 -- lesson_id / exercise_id / quiz_id tuỳ 'type'
    occurred_at  TIMESTAMPTZ NOT NULL DEFAULT now()
);
CREATE INDEX IF NOT EXISTS idx_activity_log_user_occurred ON activity_log (user_id, occurred_at DESC);