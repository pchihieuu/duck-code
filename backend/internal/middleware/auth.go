package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"backend/internal/config"
	"backend/pkg/jwtutil"
	"backend/pkg/response"
)

const (
	ctxUserID = "user_id"
	ctxRole   = "role"
)

// RequireAuth verifies the Bearer access token and injects user_id/role into context.
func RequireAuth(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		if header == "" || !strings.HasPrefix(header, "Bearer ") {
			response.Err(c, http.StatusUnauthorized, "UNAUTHORIZED", "missing bearer token")
			c.Abort()
			return
		}

		tokenStr := strings.TrimPrefix(header, "Bearer ")
		claims, err := jwtutil.ParseToken(tokenStr, cfg.JWTAccessSecret)
		if err != nil {
			response.Err(c, http.StatusUnauthorized, "UNAUTHORIZED", "invalid or expired token")
			c.Abort()
			return
		}

		c.Set(ctxUserID, claims.UserID)
		c.Set(ctxRole, claims.Role)
		c.Next()
	}
}

// RequireRole restricts an already-authenticated route to specific roles (e.g. "admin").
func RequireRole(roles ...string) gin.HandlerFunc {
	allowed := make(map[string]bool, len(roles))
	for _, r := range roles {
		allowed[r] = true
	}
	return func(c *gin.Context) {
		role, _ := c.Get(ctxRole)
		roleStr, _ := role.(string)
		if !allowed[roleStr] {
			response.Err(c, http.StatusForbidden, "FORBIDDEN", "insufficient permissions")
			c.Abort()
			return
		}
		c.Next()
	}
}
