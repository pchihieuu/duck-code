package auth

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"time"

	"github.com/redis/go-redis/v9"
)

// Repository abstracts refresh-token revocation storage in Redis, so the
// service layer never touches the redis client directly.
type Repository interface {
	// RevokeToken marks a refresh token as revoked for the given TTL.
	// Idempotent — safe to call more than once for the same token (used by
	// Logout, which should succeed even if the token was already revoked).
	RevokeToken(ctx context.Context, token string, ttl time.Duration) error
	// TryRevoke atomically revokes a token IF it wasn't already revoked
	// (Redis SETNX under the hood — a single round-trip, not a separate
	// check-then-set). Returns alreadyRevoked=true if some other request
	// won the race and revoked it first; the caller must treat that as a
	// rejection, not proceed to issue new tokens. This is what makes
	// refresh-token rotation safe against two concurrent /auth/refresh
	// calls with the same token (see docs/PLANNING.md BR-004 addendum).
	//
	// Deliberately the ONLY way to check-and-revoke — there is no separate
	// IsRevoked() method, so future code can't accidentally reintroduce the
	// check-then-set race this method exists to close.
	TryRevoke(ctx context.Context, token string, ttl time.Duration) (alreadyRevoked bool, err error)
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

func (r *repository) TryRevoke(ctx context.Context, token string, ttl time.Duration) (bool, error) {
	// SetNX returns true if the key was set by THIS call (we won the race,
	// token was not revoked before) — false if the key already existed
	// (someone else revoked it first, or it's a genuine reuse attempt).
	set, err := r.rdb.SetNX(ctx, revokedKey(token), "1", ttl).Result()
	if err != nil {
		return false, err
	}
	return !set, nil
}

// revokedKey hashes the raw token before using it as a Redis key. Raw JWTs
// are valid bearer credentials on their own — storing them verbatim as
// Redis keys means anyone with read access to Redis (ops tooling, a log
// line, an RDB dump) sees live, usable refresh tokens. Hashing means the
// key leaks nothing usable; we don't need to reverse it, only compare.
func revokedKey(token string) string {
	sum := sha256.Sum256([]byte(token))
	return "auth:revoked_refresh:" + hex.EncodeToString(sum[:])
}
