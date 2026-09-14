package auth

import (
	"context"
	"time"

	"github.com/google/uuid"

	"backend/internal/config"
	"backend/internal/user"
	apperrors "backend/pkg/errors"
	"backend/pkg/jwtutil"
	"backend/pkg/passwordhash"
)

type Service interface {
	Register(ctx context.Context, req RegisterRequest) (*user.User, error)
	Login(ctx context.Context, req LoginRequest) (*user.User, TokenPair, error)
	Refresh(ctx context.Context, refreshToken string) (TokenPair, error)
	Logout(ctx context.Context, refreshToken string) error
}

type service struct {
	users user.Repository
	repo  Repository
	cfg   *config.Config
}

func NewService(users user.Repository, repo Repository, cfg *config.Config) Service {
	return &service{users: users, repo: repo, cfg: cfg}
}

func (s *service) Register(ctx context.Context, req RegisterRequest) (*user.User, error) {
	if _, err := s.users.FindByEmail(ctx, req.Email); err == nil {
		return nil, apperrors.Conflict("email already registered")
	} else if !apperrors.Is(err, user.ErrNotFound) {
		return nil, apperrors.Internal(err)
	}

	if _, err := s.users.FindByUsername(ctx, req.Username); err == nil {
		return nil, apperrors.Conflict("username already taken")
	} else if !apperrors.Is(err, user.ErrNotFound) {
		return nil, apperrors.Internal(err)
	}

	hash, err := passwordhash.Hash(req.Password)
	if err != nil {
		return nil, apperrors.Internal(err)
	}

	u := &user.User{
		ID:           uuid.New(),
		Email:        req.Email,
		Username:     req.Username,
		PasswordHash: hash,
		DisplayName:  req.DisplayName,
		Role:         "student",
	}

	if err := s.users.Create(ctx, u); err != nil {
		return nil, apperrors.Internal(err)
	}
	return u, nil
}

// Login always returns the SAME error for "email doesn't exist" and "wrong
// password" (BR-003, docs/PLANNING.md §7.1) — never leak which one it was.
func (s *service) Login(ctx context.Context, req LoginRequest) (*user.User, TokenPair, error) {
	u, err := s.users.FindByEmail(ctx, req.Email)
	if err != nil {
		if apperrors.Is(err, user.ErrNotFound) {
			return nil, TokenPair{}, apperrors.Unauthorized("invalid email or password")
		}
		return nil, TokenPair{}, apperrors.Internal(err)
	}

	ok, err := passwordhash.Verify(req.Password, u.PasswordHash)
	if err != nil {
		return nil, TokenPair{}, apperrors.Internal(err)
	}
	if !ok {
		return nil, TokenPair{}, apperrors.Unauthorized("invalid email or password")
	}

	pair, err := s.issueTokenPair(u)
	if err != nil {
		return nil, TokenPair{}, apperrors.Internal(err)
	}
	return u, pair, nil
}

// Refresh implements rotate-on-use: the presented refresh token is atomically
// revoked (Repository.TryRevoke, a single Redis SETNX round-trip) before a
// new pair is issued. If TryRevoke reports the token was ALREADY revoked —
// either a genuine reuse attempt, or a second concurrent /auth/refresh call
// racing the first with the same token — this request is rejected outright
// rather than also issuing a token pair. That closes the
// check-then-set race a naive IsRevoked()-then-RevokeToken() sequence would
// have (two concurrent requests could both see "not revoked" and both
// successfully rotate, doubling the number of valid sessions from one
// stolen token).
func (s *service) Refresh(ctx context.Context, refreshToken string) (TokenPair, error) {
	claims, err := jwtutil.ParseToken(refreshToken, s.cfg.JWTRefreshSecret, jwtutil.TokenTypeRefresh)
	if err != nil {
		return TokenPair{}, apperrors.Unauthorized("invalid or expired refresh token")
	}

	u, err := s.users.FindByID(ctx, claims.UserID)
	if err != nil {
		// Covers both "user truly gone" and "soft-deleted" — GORM's
		// DeletedAt scoping already excludes soft-deleted rows from
		// FindByID, so this one check is sufficient (docs/PLANNING.md §15).
		return TokenPair{}, apperrors.Unauthorized("user no longer exists")
	}

	refreshTTL := time.Duration(s.cfg.JWTRefreshTTLHours) * time.Hour
	alreadyRevoked, err := s.repo.TryRevoke(ctx, refreshToken, refreshTTL)
	if err != nil {
		// Fail closed: a Redis error must NOT silently allow the refresh.
		return TokenPair{}, apperrors.Internal(err)
	}
	if alreadyRevoked {
		return TokenPair{}, apperrors.Unauthorized("refresh token already used or revoked")
	}

	pair, err := s.issueTokenPair(u)
	if err != nil {
		return TokenPair{}, apperrors.Internal(err)
	}
	return pair, nil
}

// Logout is intentionally idempotent — calling it twice (or with an
// already-revoked token) must still return success, since the goal state
// ("this token is revoked") is already achieved either way.
func (s *service) Logout(ctx context.Context, refreshToken string) error {
	refreshTTL := time.Duration(s.cfg.JWTRefreshTTLHours) * time.Hour
	if err := s.repo.RevokeToken(ctx, refreshToken, refreshTTL); err != nil {
		return apperrors.Internal(err)
	}
	return nil
}

func (s *service) issueTokenPair(u *user.User) (TokenPair, error) {
	accessTTL := time.Duration(s.cfg.JWTAccessTTLMin) * time.Minute
	refreshTTL := time.Duration(s.cfg.JWTRefreshTTLHours) * time.Hour

	access, err := jwtutil.GenerateToken(u.ID, u.Role, jwtutil.TokenTypeAccess, s.cfg.JWTAccessSecret, accessTTL)
	if err != nil {
		return TokenPair{}, err
	}
	refresh, err := jwtutil.GenerateToken(u.ID, u.Role, jwtutil.TokenTypeRefresh, s.cfg.JWTRefreshSecret, refreshTTL)
	if err != nil {
		return TokenPair{}, err
	}

	return TokenPair{
		AccessToken:  access,
		RefreshToken: refresh,
		TokenType:    "Bearer",
		ExpiresInSec: int(accessTTL.Seconds()),
	}, nil
}
