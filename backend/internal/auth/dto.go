package auth

type RegisterRequest struct {
	Email       string `json:"email" binding:"required,email"`
	Username    string `json:"username" binding:"required,alphanum,min=3,max=30"`
	Password    string `json:"password" binding:"required,min=8"`
	DisplayName string `json:"display_name" binding:"required,min=2,max=80"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// TokenPair is the INTERNAL representation passed from Service to Handler —
// it still carries both tokens as plain strings (Service stays HTTP/cookie
// agnostic, per .agents/skills/backend-architecture/SKILL.md §2 — it must
// not know about Gin or cookies). Deliberately has NO json tags: it must
// never be marshaled directly into a response. The Handler is the only
// thing allowed to touch RefreshToken (to put it in the HttpOnly cookie via
// setRefreshCookie) — every client-facing JSON response goes through
// ToAccessTokenResponse instead, which excludes it.
type TokenPair struct {
	AccessToken  string
	RefreshToken string
	TokenType    string
	ExpiresInSec int
}

// AccessTokenResponse is what actually reaches the client's JSON body.
// The refresh token is NOT here — it only ever travels as the HttpOnly
// `refresh_token` cookie (internal/auth/cookie.go), never in a JSON
// response and never readable from JavaScript.
type AccessTokenResponse struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresInSec int    `json:"expires_in"`
}

func ToAccessTokenResponse(p TokenPair) AccessTokenResponse {
	return AccessTokenResponse{
		AccessToken:  p.AccessToken,
		TokenType:    p.TokenType,
		ExpiresInSec: p.ExpiresInSec,
	}
}
