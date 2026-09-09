ALTER TABLE users
    DROP COLUMN IF EXISTS total_xp,
    DROP COLUMN IF EXISTS current_streak,
    DROP COLUMN IF EXISTS longest_streak,
    DROP COLUMN IF EXISTS last_activity_date;