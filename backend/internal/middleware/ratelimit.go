// package middleware

// import (
// 	"net/http"
// 	"sync"
// 	"time"

// 	"github.com/gin-gonic/gin"
// 	"golang.org/x/time/rate"

// 	"backend/pkg/response"
// )

// // visitor tracks the limiter for a single client IP plus its last-seen time,
// // so idle entries can be garbage collected.
// type visitor struct {
// 	limiter  *rate.Limiter
// 	lastSeen time.Time
// }

// type ipRateLimiter struct {
// 	mu       sync.Mutex
// 	visitors map[string]*visitor
// 	rps      rate.Limit
// 	burst    int
// }

// func newIPRateLimiter(rps float64, burst int) *ipRateLimiter {
// 	l := &ipRateLimiter{
// 		visitors: make(map[string]*visitor),
// 		rps:      rate.Limit(rps),
// 		burst:    burst,
// 	}
// 	go l.cleanupLoop()
// 	return l
// }

// func (l *ipRateLimiter) getLimiter(ip string) *rate.Limiter {
// 	l.mu.Lock()
// 	defer l.mu.Unlock()

// 	v, exists := l.visitors[ip]
// 	if !exists {
// 		limiter := rate.NewLimiter(l.rps, l.burst)
// 		l.visitors[ip] = &visitor{limiter: limiter, lastSeen: time.Now()}
// 		return limiter
// 	}
// 	v.lastSeen = time.Now()
// 	return v.limiter
// }

// func (l *ipRateLimiter) cleanupLoop() {
// 	for {
// 		time.Sleep(time.Minute)
// 		l.mu.Lock()
// 		for ip, v := range l.visitors {
// 			if time.Since(v.lastSeen) > 3*time.Minute {
// 				delete(l.visitors, ip)
// 			}
// 		}
// 		l.mu.Unlock()
// 	}
// }

// // RateLimit throttles requests per client IP. rps is the sustained rate
// // (requests/second), burst is the max short-term spike allowed.
// // Suggested usage: a stricter limiter on /auth/login and /auth/register
// // to slow down credential-stuffing/brute-force attempts.
// func RateLimit(rps float64, burst int) gin.HandlerFunc {
// 	limiter := newIPRateLimiter(rps, burst)

// 	return func(c *gin.Context) {
// 		ip := c.ClientIP()
// 		if !limiter.getLimiter(ip).Allow() {
// 			response.Err(c, http.StatusTooManyRequests, "RATE_LIMITED", "too many requests, please slow down")
// 			c.Abort()
// 			return
// 		}
// 		c.Next()
// 	}
// }



package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"

	"backend/pkg/logger"
	"backend/pkg/response"
)

// RateLimit implements a fixed-window counter in Redis, keyed by client IP +
// route. Redis is the single source of truth on purpose — no in-memory
// fallback — so the limit is enforced consistently across every replica of
// the API, not per-process.
//
// Known trade-off: INCR + EXPIRE are two round-trips, so a crash between them
// could leave a key without a TTL. Acceptable for auth endpoints; if you need
// exact atomicity, replace the body with a Lua script (EVAL) that does both
// in one round-trip.
//
// Failure mode: if Redis is unreachable, requests are allowed through
// (fail-open) and the error is logged — we do not silently switch to
// in-memory counting, since that would defeat the point of a shared limiter.
func RateLimit(rdb *redis.Client, limit int, window time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		key := fmt.Sprintf("ratelimit:%s:%s", c.ClientIP(), c.FullPath())

		count, err := rdb.Incr(ctx, key).Result()
		if err != nil {
			logger.L.Errorw("rate limiter: redis unavailable, failing open", "error", err)
			c.Next()
			return
		}
		if count == 1 {
			if err := rdb.Expire(ctx, key, window).Err(); err != nil {
				logger.L.Errorw("rate limiter: failed to set TTL", "error", err)
			}
		}

		if count > int64(limit) {
			response.Err(c, http.StatusTooManyRequests, "RATE_LIMITED", "too many requests, please slow down")
			c.Abort()
			return
		}
		c.Next()
	}
}
