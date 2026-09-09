// package middleware

// import (
// 	"net/http"
// 	"strings"

// 	"github.com/gin-gonic/gin"

// 	"backend/internal/config"
// 	"backend/pkg/jwtutil"
// 	"backend/pkg/response"
// )

// const (
// 	ctxUserID = "user_id"
// 	ctxRole   = "role"
// )

// // RequireAuth verifies the Bearer access token and injects user_id/role into context.
// func RequireAuth(cfg *config.Config) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		header := c.GetHeader("Authorization")
// 		if header == "" || !strings.HasPrefix(header, "Bearer ") {
// 			response.Err(c, http.StatusUnauthorized, "UNAUTHORIZED", "missing bearer token")
// 			c.Abort()
// 			return
// 		}

// 		tokenStr := strings.TrimPrefix(header, "Bearer ")
// 		claims, err := jwtutil.ParseToken(tokenStr, cfg.JWTAccessSecret)
// 		if err != nil {
// 			response.Err(c, http.StatusUnauthorized, "UNAUTHORIZED", "invalid or expired token")
// 			c.Abort()
// 			return
// 		}

// 		c.Set(ctxUserID, claims.UserID)
// 		c.Set(ctxRole, claims.Role)
// 		c.Next()
// 	}
// }

// // RequireRole restricts an already-authenticated route to specific roles (e.g. "admin").
// func RequireRole(roles ...string) gin.HandlerFunc {
// 	allowed := make(map[string]bool, len(roles))
// 	for _, r := range roles {
// 		allowed[r] = true
// 	}
// 	return func(c *gin.Context) {
// 		role, _ := c.Get(ctxRole)
// 		roleStr, _ := role.(string)
// 		if !allowed[roleStr] {
// 			response.Err(c, http.StatusForbidden, "FORBIDDEN", "insufficient permissions")
// 			c.Abort()
// 			return
// 		}
// 		c.Next()
// 	}
// }


package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

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

// GetUserID lấy userID đã được RequireAuth set vào context. Mọi handler ở
// mọi module (user, profile, gamification, progress, activity, ...) dùng
// đúng 1 hàm này thay vì tự c.Get("user_id") + type-assert rải rác — tránh
// lệch key hoặc lệch kiểu (claims.UserID đổi từ string sang uuid.UUID chẳng
// hạn) chỉ phải sửa 1 chỗ.
//
// panic nếu gọi ở route không có RequireAuth phía trước — đây là lỗi lập
// trình (thiếu middleware), nên fail loud thay vì âm thầm trả userID rỗng.
func GetUserID(c *gin.Context) uuid.UUID {
	v := c.MustGet(ctxUserID)
	switch id := v.(type) {
	case uuid.UUID:
		return id
	case string:
		parsed, err := uuid.Parse(id)
		if err != nil {
			panic("middleware: user_id in context is not a valid uuid: " + err.Error())
		}
		return parsed
	default:
		panic("middleware: unexpected type for user_id in context")
	}
}

// GetRole lấy role đã set bởi RequireAuth (rỗng nếu route không auth).
func GetRole(c *gin.Context) string {
	role, _ := c.Get(ctxRole)
	roleStr, _ := role.(string)
	return roleStr
}