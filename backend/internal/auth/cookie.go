package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"backend/internal/config"
)

// refreshCookieName and refreshCookiePath are shared constants — used by
// both setRefreshCookie and clearRefreshCookie so they can never drift
// apart (a checklist gotcha: clearing a cookie with a different
// Name/Path/Domain than it was set with silently fails — the browser just
// keeps the original cookie).
const (
	refreshCookieName = "refresh_token"
	// Scoped to /api/v1/auth only — the browser will never attach this
	// cookie to unrelated routes (/api/v1/users/me, /api/v1/courses, ...),
	// which also means it never shows up in access logs for those routes.
	refreshCookiePath = "/api/v1/auth"
)

// setRefreshCookie sets the HttpOnly refresh-token cookie. maxAge is in
// seconds and should match the refresh token's own TTL (cfg.JWTRefreshTTLHours),
// so the cookie never outlives (or expires before) the token it carries.
func setRefreshCookie(c *gin.Context, cfg *config.Config, token string, maxAgeSeconds int) {
	c.SetSameSite(sameSiteFromString(cfg.CookieSameSite))
	c.SetCookie(
		refreshCookieName,
		token,
		maxAgeSeconds,
		refreshCookiePath,
		cfg.CookieDomain,
		cfg.CookieSecure,
		true, // httpOnly — never readable from JS (localStorage/sessionStorage can't touch it either, since it's never exposed to JS at all)
	)
}

// clearRefreshCookie removes the refresh-token cookie. Reuses the exact same
// Name/Path/Domain/Secure/SameSite as setRefreshCookie — only maxAge (-1)
// and value ("") differ — so the browser recognizes it as the same cookie
// and actually deletes it instead of silently ignoring a mismatched clear.
func clearRefreshCookie(c *gin.Context, cfg *config.Config) {
	c.SetSameSite(sameSiteFromString(cfg.CookieSameSite))
	c.SetCookie(
		refreshCookieName,
		"",
		-1, // MaxAge -1 => "Expires in the past", browser deletes it now
		refreshCookiePath,
		cfg.CookieDomain,
		cfg.CookieSecure,
		true,
	)
}

func sameSiteFromString(s string) http.SameSite {
	switch s {
	case "None", "none":
		return http.SameSiteNoneMode
	case "Strict", "strict":
		return http.SameSiteStrictMode
	default:
		return http.SameSiteLaxMode
	}
}
