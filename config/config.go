package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Database struct {
		Host     string
		Port     int
		User     string
		Password string
		Name     string
	}
	Server struct {
		Port int
	}
}

// LoadConfig loads configuration from environment variables
func LoadConfig() (*Config, error) {
	// Carregar variáveis do .env
	if err := godotenv.Load(); err != nil {
		log.Println("Arquivo .env não encontrado, usando variáveis de ambiente padrão")
	}

	cfg := &Config{}
	// Mapear as variáveis do .env para o struct
	cfg.Database.Host = os.Getenv("DB_HOST")
	cfg.Database.Port = getEnvAsInt("DB_PORT", 5432)
	cfg.Database.User = os.Getenv("DB_USER")
	cfg.Database.Password = os.Getenv("DB_PASSWORD")
	cfg.Database.Name = os.Getenv("DB_NAME")
	cfg.Server.Port = getEnvAsInt("SERVER_PORT", 8080)

	return cfg, nil
}

// Helper functions
func getEnvAsInt(key string, fallback int) int {
	if value, exists := os.LookupEnv(key); exists {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return fallback
}
