package user

import "github.com/google/uuid"

// UpdateProfileRequest is the payload for PATCH /users/me.
type UpdateProfileRequest struct {
	DisplayName string `json:"display_name" binding:"omitempty,min=2,max=80"`
	AvatarURL   string `json:"avatar_url" binding:"omitempty,url"`
}

// PublicResponse is the safe-to-expose shape of a user, used everywhere
// except internal auth flows.
type PublicResponse struct {
	ID          uuid.UUID `json:"id"`
	Email       string    `json:"email"`
	Username    string    `json:"username"`
	DisplayName string    `json:"display_name"`
	AvatarURL   *string   `json:"avatar_url,omitempty"`
	Role        string    `json:"role"`
}

func ToPublicResponse(u *User) PublicResponse {
	return PublicResponse{
		ID:          u.ID,
		Email:       u.Email,
		Username:    u.Username,
		DisplayName: u.DisplayName,
		AvatarURL:   u.AvatarURL,
		Role:        u.Role,
	}
}
