package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	JWTSecret  string
	Port       string
}

// Load читает переменные окружения. Если .env присутствует рядом с бинарником/проектом — подхватывает.
// В Docker рекомендуется полагаться на переменные из docker-compose.
func Load() *Config {
	// .env опционален: в Docker чаще всего не нужен
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	cfg := &Config{
		// В контейнере DB_HOST почти всегда "postgres" — имя сервиса в docker-compose
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", "password"),
		DBName:     getEnv("DB_NAME", "smolathon_db"),
		JWTSecret:  getEnv("JWT_SECRET", "your-default-secret-key-change-in-production"),
		Port:       getEnv("PORT", "8080"),
	}

	if cfg.JWTSecret == "your-default-secret-key-change-in-production" {
		log.Println("WARNING: Using default JWT secret. Set JWT_SECRET environment variable in production!")
	}

	return cfg
}

func getEnv(key, defaultVal string) string {
	if v, ok := os.LookupEnv(key); ok && v != "" {
		return v
	}
	return defaultVal
}
