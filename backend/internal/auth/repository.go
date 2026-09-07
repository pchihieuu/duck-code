package auth

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

// Repository abstracts session/refresh-token storage in Redis, so the
// service layer never touches the redis client directly.
type Repository interface {
	// RevokeToken marks a refresh token as revoked for the given TTL
	// (used on logout and on every refresh-token rotation).
	RevokeToken(ctx context.Context, token string, ttl time.Duration) error
	// IsRevoked reports whether a refresh token has already been revoked.
	IsRevoked(ctx context.Context, token string) (bool, error)
}

type repository struct {
	rdb *redis.Client
}

func NewRepository(rdb *redis.Client) Repository {
	return &repository{rdb: rdb}
}

func (r *repository) RevokeToken(ctx context.Context, token string, ttl time.Duration) error {
	return r.rdb.Set(ctx, revokedKey(token), "1", ttl).Err()
}

func (r *repository) IsRevoked(ctx context.Context, token string) (bool, error) {
	n, err := r.rdb.Exists(ctx, revokedKey(token)).Result()
	if err != nil {
		return false, err
	}
	return n > 0, nil
}

func revokedKey(token string) string {
	return "auth:revoked_refresh:" + token
}
