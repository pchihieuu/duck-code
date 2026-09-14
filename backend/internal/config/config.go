package config

import (
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

// Config holds all environment-driven settings for the service.
type Config struct {
	AppEnv string
	Port   string

	DBURL string

	RedisHost     string
	RedisPort     string
	RedisPassword string
	RedisDB       int

	RabbitMQURL string

	JWTAccessSecret    string
	JWTRefreshSecret   string
	JWTAccessTTLMin    int
	JWTRefreshTTLHours int

	// AllowedOrigins is the CORS whitelist. Required (non-"*") once the
	// refresh token moved to an HttpOnly cookie — browsers refuse
	// "Access-Control-Allow-Origin: *" together with credentialed
	// (cookie-carrying) requests, so this MUST be an explicit origin list,
	// never a wildcard. See internal/middleware/cors.go.
	AllowedOrigins []string

	// Cookie attributes for the refresh_token cookie
	// (internal/auth/cookie.go). Defaults below are safe for local HTTP
	// dev (same registrable domain, different port — e.g.
	// localhost:3000 -> localhost:8080 — which is "same-site" for
	// SameSite purposes even though it's cross-origin). In production,
	// with FE/BE on different domains, set COOKIE_SECURE=true and
	// COOKIE_SAMESITE=None (SameSite=None requires Secure=true or
	// browsers drop the cookie).
	CookieDomain   string
	CookieSecure   bool
	CookieSameSite string
}

// Load reads .env (if present) then environment variables into a Config.
// Environment variables always take precedence over .env values.
func Load() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("no .env file found, relying on real environment variables")
	}

	cfg := &Config{
		AppEnv: getEnv("APP_ENV", "development"),
		Port:   getEnv("PORT", "8080"),

		DBURL: getEnv("DB_URL", ""),

		RedisHost:     getEnv("REDIS_HOST", "localhost"),
		RedisPort:     getEnv("REDIS_PORT", "6379"),
		RedisPassword: getEnv("REDIS_PASSWORD", ""),
		RedisDB:       getEnvInt("REDIS_DB", 0),

		RabbitMQURL: getEnv("RABBITMQ_URL", "amqp://guest:guest@localhost:5672/"),

		JWTAccessSecret:    getEnv("JWT_ACCESS_SECRET", ""),
		JWTRefreshSecret:   getEnv("JWT_REFRESH_SECRET", ""),
		JWTAccessTTLMin:    getEnvInt("JWT_ACCESS_TTL_MIN", 15),
		JWTRefreshTTLHours: getEnvInt("JWT_REFRESH_TTL_HOURS", 168),

		AllowedOrigins: getEnvList("ALLOWED_ORIGINS", []string{"http://localhost:3000"}),

		CookieDomain:   getEnv("COOKIE_DOMAIN", ""),
		CookieSecure:   getEnvBool("COOKIE_SECURE", false),
		CookieSameSite: getEnv("COOKIE_SAMESITE", "Lax"),
	}

	if cfg.DBURL == "" {
		log.Fatal("DB_URL is required")
	}
	if cfg.JWTAccessSecret == "" || cfg.JWTRefreshSecret == "" {
		log.Fatal("JWT_ACCESS_SECRET and JWT_REFRESH_SECRET are required")
	}
	if cfg.JWTAccessSecret == cfg.JWTRefreshSecret {
		log.Fatal("JWT_ACCESS_SECRET and JWT_REFRESH_SECRET must be different values")
	}

	return cfg
}

func getEnv(key, fallback string) string {
	if v, ok := os.LookupEnv(key); ok && v != "" {
		return v
	}
	return fallback
}

func getEnvInt(key string, fallback int) int {
	v, ok := os.LookupEnv(key)
	if !ok || v == "" {
		return fallback
	}
	i, err := strconv.Atoi(v)
	if err != nil {
		return fallback
	}
	return i
}

func getEnvBool(key string, fallback bool) bool {
	v, ok := os.LookupEnv(key)
	if !ok || v == "" {
		return fallback
	}
	b, err := strconv.ParseBool(v)
	if err != nil {
		return fallback
	}
	return b
}

// getEnvList parses a comma-separated env var into a trimmed string slice.
func getEnvList(key string, fallback []string) []string {
	v, ok := os.LookupEnv(key)
	if !ok || v == "" {
		return fallback
	}
	parts := strings.Split(v, ",")
	out := make([]string, 0, len(parts))
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p != "" {
			out = append(out, p)
		}
	}
	if len(out) == 0 {
		return fallback
	}
	return out
}
