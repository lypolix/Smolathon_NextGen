package main

import (
	"log"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv" // импорт для загрузки .env

	"backend/config"
	"backend/internal/api"
	"backend/internal/store"
)

// mask скрывает большую часть пароля в логах
func mask(s string) string {
	if s == "" {
		return ""
	}
	if len(s) <= 2 {
		return "*"
	}
	return s[:2] + strings.Repeat("*", len(s)-2)
}

// logDBEnv печатает ключевые переменные окружения БД с маскированным паролем
func logDBEnv() {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	name := os.Getenv("DB_NAME")
	pass := os.Getenv("DB_PASSWORD")

	log.Printf("DB env -> DB_HOST=%q DB_PORT=%q DB_USER=%q DB_NAME=%q DB_PASSWORD=%q",
		host, port, user, name, mask(pass))
}

func main() {
	// 1) Автоподгрузка .env (из корня проекта)
	// Пытаемся загрузить .env бесшумно: если файла нет, продолжаем с системным окружением.
	// Можно явно указать путь: godotenv.Load(".env")
	if err := godotenv.Load(); err != nil {
		log.Printf("No .env file found or failed to load: %v", err)
	}

	// 2) Лог текущих env (после загрузки .env)
	logDBEnv()

	// 3) Загрузка конфигурации приложения
	cfg := config.Load()

	// 4) Лог API порта
	log.Printf("API PORT=%q", cfg.Port)

	// 5) Подключение к БД
	s, err := store.NewStore(cfg)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer s.Close()

	// 6) HTTP роутер
	r := gin.Default()
	api.RegisterRoutes(r, s, cfg)

	// 7) Старт сервера
	log.Printf("Server starting on port %s", cfg.Port)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
