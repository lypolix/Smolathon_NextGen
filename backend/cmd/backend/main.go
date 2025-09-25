package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

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

// validateCfg — минимальная валидация критичных полей
func validateCfg(cfg *config.Config) error {
	if cfg == nil {
		return fmt.Errorf("nil config")
	}
	if cfg.Port == "" {
		return fmt.Errorf("empty API port")
	}
	// при необходимости добавить проверки cfg.DB.* и т.д.
	return nil
}

func main() {
	// 1) .env (не фатально при отсутствии)
	if err := godotenv.Load(); err != nil {
		log.Printf("No .env file found or failed to load: %v", err)
	}

	// 2) Логи ENV для БД
	logDBEnv()

	// 3) Конфиг и валидация
	cfg := config.Load() // предположительно возвращает *config.Config
	if err := validateCfg(cfg); err != nil {
		log.Fatalf("Invalid config: %v", err)
	}

	// 4) Режим Gin
	// Если поля Debug нет в конфиге — используем значение из окружения GIN_MODE.
	// Можно явно зафиксировать Debug Mode в деве:
	if os.Getenv("GIN_MODE") == "" {
		// Чтобы видеть подробные логи локально
		gin.SetMode(gin.DebugMode)
	}
	log.Printf("Gin mode: %s", gin.Mode())

	// 5) Подключение к БД
	s, err := store.NewStore(cfg)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer func() {
		if cerr := s.Close(); cerr != nil {
			log.Printf("DB close error: %v", cerr)
		}
	}()

	// 6) Роутер с Logger/Recovery
	r := gin.Default()

	// Если RegisterRoutes сам добавляет префикс /api — оставляем как есть.
	// Если нет — можно обернуть в группу: g := r.Group("/api"); api.RegisterRoutes(g, s, cfg)
	api.RegisterRoutes(r, s, cfg)

	// 7) HTTP-сервер с таймаутами
	srv := &http.Server{
		Addr:           ":" + cfg.Port,
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   30 * time.Second,
		IdleTimeout:    60 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Printf("Server starting on port %s", cfg.Port)
	if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatal("Failed to start server:", err)
	}
}
