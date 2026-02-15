package config

import (
	"os"

	"github.com/joho/godotenv"
)

// ConfigurationService holds all application configuration
type ConfigurationService struct {
	ServerPort       string
	DatabaseURL      string
	GinMode          string
	FrontendURL       string
	JWTSecret        string
	PaymentServiceURL string
	AuthAPIURL       string
}

var instance *ConfigurationService

// GetInstance returns the singleton configuration instance
func GetInstance() *ConfigurationService {
	if instance == nil {
		_ = godotenv.Load()

		instance = &ConfigurationService{
			ServerPort:       getEnvOrDefault("SERVER_PORT", "8080"),
			DatabaseURL:      getEnvOrDefault("DATABASE_URL", "postgres://postgres:postgres@localhost:5432/wappi?sslmode=disable"),
			GinMode:          getEnvOrDefault("GIN_MODE", "debug"),
			FrontendURL:       getEnvOrDefault("FRONTEND_URL", "http://localhost:5173"),
			JWTSecret:        getEnvOrDefault("JWT_SECRET", ""),
			PaymentServiceURL: getEnvOrDefault("PAYMENT_SERVICE_URL", "http://localhost:8008"),
			AuthAPIURL:       getEnvOrDefault("AUTH_API_URL", "http://localhost:8082"),
		}
	}
	return instance
}

func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
