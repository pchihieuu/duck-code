package jwtutil

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type TokenType string

const (
	TokenTypeAccess  TokenType = "access"
	TokenTypeRefresh TokenType = "refresh"
)

var (
	ErrUnexpectedSigningMethod = errors.New("jwtutil: unexpected signing method")
	ErrInvalidToken            = errors.New("jwtutil: invalid token")
	ErrWrongTokenType          = errors.New("jwtutil: token type does not match expected type")
)
// Claims embedded in every access/refresh token issued by the platform.
type Claims struct {
	UserID uuid.UUID `json:"user_id"`
	Role   string    `json:"role"`
	Type  TokenType `json:"type"`
	jwt.RegisteredClaims
}

// GenerateToken signs a JWT for the given user with the given TTL and secret.
func GenerateToken(userID uuid.UUID, role string, tokenType TokenType, secret string, ttl time.Duration) (string, error) {
	claims := Claims{
		UserID: userID,
		Role:   role,
		Type: tokenType,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(ttl)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

// ParseToken verifies signature and expiry, returning the embedded claims.
func ParseToken(tokenStr, secret string, expectedType TokenType) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrUnexpectedSigningMethod
		}
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, ErrInvalidToken
	}
	if expectedType != "" && claims.Type != expectedType {
		return nil, ErrWrongTokenType
	}
	return claims, nil
}
