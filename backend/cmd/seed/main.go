// cmd/seed bootstraps the very first admin account by writing directly to
// the database — this deliberately bypasses the HTTP API, since the public
// /auth/register endpoint must never accept a "role" field (see
// .agents/skills/backend-architecture/SKILL.md §5 and internal/auth BR-005
// in docs/PLANNING.md). Run once per environment:
//
//	ADMIN_EMAIL=admin@example.com ADMIN_USERNAME=admin ADMIN_PASSWORD=... \
//	ADMIN_DISPLAY_NAME="Platform Admin" make seed-admin
//
// Safe to re-run: if the email already exists, it only promotes that
// account to admin (idempotent) instead of failing or creating a duplicate.
package main

import (
	"context"
	"os"

	"github.com/google/uuid"

	"backend/internal/config"
	"backend/internal/user"
	apperrors "backend/pkg/errors"
	"backend/pkg/database"
	"backend/pkg/logger"
	"backend/pkg/passwordhash"
)

func main() {
	cfg := config.Load()
	logger.Init(cfg.AppEnv)
	defer logger.Sync()

	email := requireEnv("ADMIN_EMAIL")
	username := requireEnv("ADMIN_USERNAME")
	password := requireEnv("ADMIN_PASSWORD")
	displayName := getEnvOr("ADMIN_DISPLAY_NAME", "Admin")

	db, err := database.NewPostgres(cfg)
	if err != nil {
		logger.L.Fatalw("failed to connect to postgres", "error", err)
	}
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	repo := user.NewRepository(db)
	ctx := context.Background()

	existing, err := repo.FindByEmail(ctx, email)
	if err == nil {
		// Already exists — just make sure it's an admin. Never log the
		// password/hash here, only the outcome.
		if existing.Role == "admin" {
			logger.L.Infow("admin already exists, nothing to do", "email", email)
			return
		}
		if err := repo.UpdateRole(ctx, existing.ID, "admin"); err != nil {
			logger.L.Fatalw("failed to promote existing user to admin", "error", err)
		}
		logger.L.Infow("promoted existing user to admin", "email", email)
		return
	}
	if !apperrors.Is(err, user.ErrNotFound) {
		logger.L.Fatalw("failed to check for existing admin", "error", err)
	}

	hash, err := passwordhash.Hash(password)
	if err != nil {
		logger.L.Fatalw("failed to hash admin password", "error", err)
	}

	admin := &user.User{
		ID:           uuid.New(),
		Email:        email,
		Username:     username,
		PasswordHash: hash,
		DisplayName:  displayName,
		Role:         "admin",
	}
	if err := repo.Create(ctx, admin); err != nil {
		logger.L.Fatalw("failed to create admin user", "error", err)
	}

	logger.L.Infow("admin user created", "email", email, "username", username)
}

func requireEnv(key string) string {
	v := os.Getenv(key)
	if v == "" {
		panic("seed: required env var " + key + " is not set")
	}
	return v
}

func getEnvOr(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
