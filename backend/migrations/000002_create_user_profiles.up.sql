CREATE TABLE IF NOT EXISTS user_profiles (
    user_id             UUID PRIMARY KEY REFERENCES users(id) ON DELETE CASCADE,
    bio                 TEXT,
    github_url          TEXT,
    website_url         TEXT,
    timezone            VARCHAR(64) NOT NULL DEFAULT 'UTC',
    preferred_language  VARCHAR(32),
    show_on_leaderboard BOOLEAN     NOT NULL DEFAULT true,
    created_at          TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at          TIMESTAMPTZ NOT NULL DEFAULT now()
);
