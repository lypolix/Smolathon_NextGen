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

func Load() *Config {
    // Пытаемся загрузить .env файл (не критично если его нет)
    if err := godotenv.Load(); err != nil {
        log.Println("No .env file found, using environment variables")
    }

    config := &Config{
        DBHost:     getEnv("DB_HOST", "localhost"),
        DBPort:     getEnv("DB_PORT", "5432"),
        DBUser:     getEnv("DB_USER", "postgres"),
        DBPassword: getEnv("DB_PASSWORD", "password"),
        DBName:     getEnv("DB_NAME", "smolathon_db"),
        JWTSecret:  getEnv("JWT_SECRET", "your-default-secret-key-change-in-production"),
        Port:       getEnv("PORT", "8080"),
    }

    // Проверяем критические параметры
    if config.JWTSecret == "your-default-secret-key-change-in-production" {
        log.Println("WARNING: Using default JWT secret. Set JWT_SECRET environment variable in production!")
    }

    return config
}

func getEnv(key, defaultVal string) string {
    if value, exists := os.LookupEnv(key); exists {
        return value
    }
    return defaultVal
}
