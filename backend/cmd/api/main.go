package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"backend/internal/config"
	"backend/internal/router"
	"backend/pkg/database"
	"backend/pkg/logger"
	"backend/pkg/validator"
)

func main() {
	cfg := config.Load()
	logger.Init(cfg.AppEnv)
	defer logger.Sync()
	validator.Get() // register custom validation rules before routes bind requests

	db, err := database.NewPostgres(cfg)
	if err != nil {
		logger.L.Fatalw("failed to connect to postgres", "error", err)
	}
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	rdb, err := database.NewRedis(cfg)
	if err != nil {
		logger.L.Fatalw("failed to connect to redis", "error", err)
	}
	defer rdb.Close()

	engine := router.New(db, rdb, cfg)

	srv := &http.Server{
		Addr:              ":" + cfg.Port,
		Handler:           engine,
		ReadHeaderTimeout: 5 * time.Second,
	}

	go func() {
		logger.L.Infow("api server starting", "port", cfg.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.L.Fatalw("server failed", "error", err)
		}
	}()

	// graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.L.Info("shutting down gracefully")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.L.Errorw("forced shutdown", "error", err)
	}
}
