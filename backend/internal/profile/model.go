package profile

import (
	"time"

	"github.com/google/uuid"
)

// Profile holds extended, optional-to-fill data about a user — separate
// from internal/user (which holds auth-critical + always-needed-for-UI
// fields like display_name/avatar_url). One-to-one with users via UserID.
//
// Rule of thumb for what belongs here vs in internal/user: if a field is
// read on nearly every request that renders the user anywhere (submission
// list, leaderboard row, comment author) it stays in `users`; if it's only
// needed on the user's own profile page or an admin view, it belongs here.
type Profile struct {
	UserID            uuid.UUID `gorm:"type:uuid;primaryKey" json:"user_id"`
	Bio               *string   `json:"bio,omitempty"`
	GithubURL         *string   `json:"github_url,omitempty"`
	WebsiteURL        *string   `json:"website_url,omitempty"`
	Timezone          string    `gorm:"not null;default:UTC" json:"timezone"`
	PreferredLanguage *string   `json:"preferred_language,omitempty"` // ngôn ngữ lập trình đang tập trung học, vd "python"
	ShowOnLeaderboard bool      `gorm:"not null;default:true" json:"show_on_leaderboard"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

func (Profile) TableName() string {
	return "user_profiles"
}
