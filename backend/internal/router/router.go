// package router

// import (
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// 	"github.com/redis/go-redis/v9"
// 	"gorm.io/gorm"

// 	"backend/internal/auth"
// 	"backend/internal/config"
// 	"backend/internal/middleware"
// 	"backend/internal/user"
// 	"backend/pkg/response"
// )

// // New builds the full Gin engine with all modules wired together.
// func New(db *gorm.DB, rdb *redis.Client, cfg *config.Config) *gin.Engine {
// 	if cfg.AppEnv == "production" {
// 		gin.SetMode(gin.ReleaseMode)
// 	}

// 	r := gin.New()
// 	r.Use(middleware.Recover(), middleware.Logger(), middleware.CORS())

// 	r.GET("/health", func(c *gin.Context) {
// 		response.OK(c, http.StatusOK, gin.H{"status": "ok"})
// 	})

// 	v1 := r.Group("/api/v1")

// 	// --- auth (public, rate-limited against brute force) ---
// 	userRepo := user.NewRepository(db)
// 	authRepo := auth.NewRepository(rdb)
// 	authSvc := auth.NewService(userRepo, authRepo, cfg)
// 	authHandler := auth.NewHandler(authSvc)

// 	authGroup := v1.Group("")
// 	authGroup.Use(middleware.RateLimit(2, 10)) // 2 req/s sustained, burst 10 per IP
// 	auth.RegisterRoutes(authGroup, authHandler)

// 	// --- protected routes ---
// 	protected := v1.Group("")
// 	protected.Use(middleware.RequireAuth(cfg))

// 	userSvc := user.NewService(userRepo)
// 	userHandler := user.NewHandler(userSvc)
// 	user.RegisterRoutes(protected, userHandler)

// 	// Phase 2+ modules (progress, activity, exercise, quiz, submission,
// 	// gamification, achievement, leaderboard, admin) plug in here the same
// 	// way: build repo -> service -> handler -> RegisterRoutes(protected, h).

// 	return r
// }


package router

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	"backend/internal/auth"
	"backend/internal/config"
	"backend/internal/middleware"
	"backend/internal/user"
	"backend/pkg/response"
)

// New builds the full Gin engine with all modules wired together.
func New(db *gorm.DB, rdb *redis.Client, cfg *config.Config) *gin.Engine {
	if cfg.AppEnv == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()
	r.Use(middleware.Recover(), middleware.Logger(), middleware.CORS())

	r.GET("/health", func(c *gin.Context) {
		response.OK(c, http.StatusOK, gin.H{"status": "ok"})
	})

	v1 := r.Group("/api/v1")

	// --- auth (public, rate-limited against brute force) ---
	userRepo := user.NewRepository(db)
	authRepo := auth.NewRepository(rdb)
	authSvc := auth.NewService(userRepo, authRepo, cfg)
	authHandler := auth.NewHandler(authSvc)

	authGroup := v1.Group("")
	authGroup.Use(middleware.RateLimit(rdb, 10, time.Minute)) // 10 req/min per IP per route
	auth.RegisterRoutes(authGroup, authHandler)

	// --- protected routes ---
	protected := v1.Group("")
	protected.Use(middleware.RequireAuth(cfg))

	userSvc := user.NewService(userRepo)
	userHandler := user.NewHandler(userSvc)
	user.RegisterRoutes(protected, userHandler)

	// Phase 2+ modules (progress, activity, exercise, quiz, submission,
	// gamification, achievement, leaderboard, admin) plug in here the same
	// way: build repo -> service -> handler -> RegisterRoutes(protected, h).

	return r
}
