package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CORS allows the frontend (Astro/React) to call the API. In production,
// swap the wildcard for cfg.AllowedOrigins loaded from env.
func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	}
}
