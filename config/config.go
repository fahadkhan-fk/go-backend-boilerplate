package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort       string
	DatabaseURL      string
	DatabaseHost     string
	DatabasePort     string
	DatabaseUsername string
	DatabasePassword string
	DatabaseDb       string
	DatabaseSslMode  string
}

func Load() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using default environment variables")
	}

	return &Config{
		ServerPort:       getEnv("SERVER_PORT", "8080"),
		DatabaseURL:      getEnv("DATABASE_URL", "postgres://postgres:postgres@localhost:5432/go_boilerplate?sslmode=disable"),
		DatabaseHost:     getEnv("DATABASE_HOST", "localhost"),
		DatabasePort:     getEnv("DATABASE_PORT", "5432"),
		DatabaseUsername: getEnv("DATABASE_USERNAME", "postgres"),
		DatabasePassword: getEnv("DATABASE_PASSWORD", "postgres"),
		DatabaseDb:       getEnv("DATABASE_DB", "go_boilerplate"),
		DatabaseSslMode:  getEnv("DATABASE_SSL_MODE", "disable"),
	}
}

func getEnv(key, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}
