package router

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	"backend/internal/activity"
	"backend/internal/auth"
	"backend/internal/config"
	"backend/internal/course"
	"backend/internal/gamification"
	"backend/internal/lesson"
	"backend/internal/middleware"
	"backend/internal/profile"
	"backend/internal/progress"
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

	profileRepo := profile.NewRepository(db)
	profileSvc := profile.NewService(profileRepo)
	profileHandler := profile.NewHandler(profileSvc)
	profile.RegisterRoutes(protected, profileHandler)

	// --- admin-only routes (authenticated AND role=admin) ---
	adminGroup := v1.Group("")
	adminGroup.Use(middleware.RequireAuth(cfg), middleware.RequireRole("admin"))
	user.RegisterAdminRoutes(adminGroup, userHandler)

	// --- Phase 2: course, lesson, activity, gamification, progress ---
	// Thứ tự khởi tạo bám theo dependency: course/lesson độc lập; activity
	// độc lập; gamification chỉ cần userRepo; progress cần cả lessonRepo,
	// activitySvc, gamificationSvc nên khởi tạo sau cùng.
	courseRepo := course.NewRepository(db)
	courseSvc := course.NewService(courseRepo)
	courseHandler := course.NewHandler(courseSvc)
	course.RegisterRoutes(v1, courseHandler)              // public, GET only
	course.RegisterAdminRoutes(adminGroup, courseHandler) // admin CRUD

	lessonRepo := lesson.NewRepository(db)
	lessonSvc := lesson.NewService(lessonRepo, courseRepo) // validate course_id tồn tại
	lessonHandler := lesson.NewHandler(lessonSvc)
	lesson.RegisterRoutes(v1, lessonHandler)
	lesson.RegisterAdminRoutes(adminGroup, lessonHandler)

	activityRepo := activity.NewRepository(db)
	activitySvc := activity.NewService(activityRepo)
	activityHandler := activity.NewHandler(activitySvc)
	activity.RegisterRoutes(protected, activityHandler)

	gamificationSvc := gamification.NewService(userRepo) // userRepo đã có sẵn ở trên
	gamificationHandler := gamification.NewHandler(gamificationSvc)
	gamification.RegisterRoutes(protected, gamificationHandler)

	progressRepo := progress.NewRepository(db)
	progressSvc := progress.NewService(progressRepo, lessonRepo, activitySvc, gamificationSvc)
	progressHandler := progress.NewHandler(progressSvc)
	progress.RegisterRoutes(protected, progressHandler)

	return r
}