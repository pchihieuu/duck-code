package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"backend/pkg/logger"
	"backend/pkg/response"
)

// Recover turns panics into a clean 500 JSON response instead of a crash.
func Recover() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				logger.L.Errorw("recovered from panic",
					"panic", r,
					"path", c.Request.URL.Path,
				)
				response.Err(c, http.StatusInternalServerError, "INTERNAL", "something went wrong")
				c.Abort()
			}
		}()
		c.Next()
	}
}
