package config

import (
	"log"
	"os"
	"strconv"

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

	JWTAccessSecret     string
	JWTRefreshSecret    string
	JWTAccessTTLMin     int
	JWTRefreshTTLHours  int
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
	}

	if cfg.DBURL == "" {
		log.Fatal("DB_URL is required")
	}
	if cfg.JWTAccessSecret == "" || cfg.JWTRefreshSecret == "" {
		log.Fatal("JWT_ACCESS_SECRET and JWT_REFRESH_SECRET are required")
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
