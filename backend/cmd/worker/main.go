package main

import (
	"backend/internal/config"
	"backend/pkg/database"
	"backend/pkg/logger"
)

// Placeholder entrypoint for the Phase 5 execution worker: consume Code
// Runner jobs from RabbitMQ, run inside a sandbox, persist the result back
// to Postgres via internal/submission. Wire internal/execution/* here once
// that module is implemented.
func main() {
	cfg := config.Load()
	logger.Init(cfg.AppEnv)
	defer logger.Sync()

	db, err := database.NewPostgres(cfg)
	if err != nil {
		logger.L.Fatalw("failed to connect to postgres", "error", err)
	}
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	conn, ch, err := database.NewRabbitMQ(cfg)
	if err != nil {
		logger.L.Fatalw("failed to connect to rabbitmq", "error", err)
	}
	defer ch.Close()
	defer conn.Close()

	logger.L.Info("execution worker starting (consumer not yet implemented)")

	// TODO(Phase 5): ch.Consume("code_execution_jobs", ...), run inside
	// sandbox, write result + score back to Postgres.
	select {}
}
