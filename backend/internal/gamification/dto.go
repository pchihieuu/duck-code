package gamification

// Response is returned by GET /gamification/me.
type Response struct {
	TotalXP       int `json:"total_xp"`
	Level         int `json:"level"`
	CurrentStreak int `json:"current_streak"`
	LongestStreak int `json:"longest_streak"`
}