package profile

import "github.com/google/uuid"

// UpdateRequest is the payload for PATCH /profile/me. All fields optional —
// only non-nil/non-empty ones are applied (partial update).
type UpdateRequest struct {
	Bio               *string `json:"bio" binding:"omitempty,max=500"`
	GithubURL         *string `json:"github_url" binding:"omitempty,url"`
	WebsiteURL        *string `json:"website_url" binding:"omitempty,url"`
	Timezone          *string `json:"timezone" binding:"omitempty,max=64"`
	PreferredLanguage *string `json:"preferred_language" binding:"omitempty,max=32"`
	ShowOnLeaderboard *bool   `json:"show_on_leaderboard"`
}

// Response is the shape returned by the profile endpoints.
type Response struct {
	UserID            uuid.UUID `json:"user_id"`
	Bio               *string   `json:"bio,omitempty"`
	GithubURL         *string   `json:"github_url,omitempty"`
	WebsiteURL        *string   `json:"website_url,omitempty"`
	Timezone          string    `json:"timezone"`
	PreferredLanguage *string   `json:"preferred_language,omitempty"`
	ShowOnLeaderboard bool      `json:"show_on_leaderboard"`
}

func ToResponse(p *Profile) Response {
	return Response{
		UserID:            p.UserID,
		Bio:               p.Bio,
		GithubURL:         p.GithubURL,
		WebsiteURL:        p.WebsiteURL,
		Timezone:          p.Timezone,
		PreferredLanguage: p.PreferredLanguage,
		ShowOnLeaderboard: p.ShowOnLeaderboard,
	}
}
