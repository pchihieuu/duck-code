package middleware

import (
	"time"

	"github.com/gin-gonic/gin"

	"backend/pkg/logger"
)

// Logger logs each request with method, path, status and latency via Zap.
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()

		logger.L.Infow("request",
			"method", c.Request.Method,
			"path", c.Request.URL.Path,
			"status", c.Writer.Status(),
			"latency_ms", time.Since(start).Milliseconds(),
			"client_ip", c.ClientIP(),
		)
	}
}
