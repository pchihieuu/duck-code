package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CORS allows the frontend to call the API with credentials (cookies)
// included. This CANNOT use "Access-Control-Allow-Origin: *" — browsers
// reject wildcard origin together with credentialed requests, and even if
// they didn't, wildcard + credentials is a severe CSRF hole (any site could
// silently ride the user's cookie). allowedOrigins must be an explicit list
// (see config.AllowedOrigins, env ALLOWED_ORIGINS) — this now takes it as a
// parameter instead of hardcoding "*", which is what made the refresh-token
// cookie flow actually work cross-origin in the first place.
func CORS(allowedOrigins []string) gin.HandlerFunc {
	allowed := make(map[string]bool, len(allowedOrigins))
	for _, o := range allowedOrigins {
		allowed[o] = true
	}

	return func(c *gin.Context) {
		origin := c.GetHeader("Origin")
		if origin != "" && allowed[origin] {
			// Echo back the SPECIFIC matched origin, never "*" — required
			// for Allow-Credentials to be valid, and is what scopes the
			// grant to only origins we've actually whitelisted.
			c.Header("Access-Control-Allow-Origin", origin)
			c.Header("Access-Control-Allow-Credentials", "true")
			c.Header("Vary", "Origin") // response differs per Origin — don't let caches conflate them
		}
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	}
}
