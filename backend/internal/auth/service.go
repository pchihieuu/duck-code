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

func (s *service) Refresh(ctx context.Context, refreshToken string) (TokenPair, error) {
	claims, err := jwtutil.ParseToken(refreshToken, s.cfg.JWTRefreshSecret)
	if err != nil {
		return TokenPair{}, apperrors.Unauthorized("invalid or expired refresh token")
	}

	revoked, err := s.repo.IsRevoked(ctx, refreshToken)
	if err != nil {
		return TokenPair{}, apperrors.Internal(err)
	}
	if revoked {
		return TokenPair{}, apperrors.Unauthorized("refresh token revoked")
	}

	u, err := s.users.FindByID(ctx, claims.UserID)
	if err != nil {
		return TokenPair{}, apperrors.Unauthorized("user no longer exists")
	}

	// Rotate: invalidate the old refresh token, issue a fresh pair.
	refreshTTL := time.Duration(s.cfg.JWTRefreshTTLHours) * time.Hour
	if err := s.repo.RevokeToken(ctx, refreshToken, refreshTTL); err != nil {
		return TokenPair{}, apperrors.Internal(err)
	}

	pair, err := s.issueTokenPair(u)
	if err != nil {
		return TokenPair{}, apperrors.Internal(err)
	}
	return pair, nil
}

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

	access, err := jwtutil.GenerateToken(u.ID, u.Role, s.cfg.JWTAccessSecret, accessTTL)
	if err != nil {
		return TokenPair{}, err
	}
	refresh, err := jwtutil.GenerateToken(u.ID, u.Role, s.cfg.JWTRefreshSecret, refreshTTL)
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
