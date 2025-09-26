package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"

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
	if cfg.DBHost == "" || cfg.DBPort == "" || cfg.DBUser == "" || cfg.DBName == "" {
		return fmt.Errorf("incomplete DB config")
	}
	if cfg.JWTSecret == "" {
		return fmt.Errorf("empty JWT secret")
	}
	return nil
}

func main() {
	// Конфиг: читает переменные окружения и при наличии .env — подхватывает
	cfg := config.Load()

	// Логи ENV для БД (после Load, чтобы видеть итоговые значения)
	logDBEnv()

	// Валидация конфига
	if err := validateCfg(cfg); err != nil {
		log.Fatalf("Invalid config: %v", err)
	}

	// Режим Gin (используем GIN_MODE из окружения; по умолчанию debug)
	if os.Getenv("GIN_MODE") == "" {
		gin.SetMode(gin.DebugMode)
	}
	log.Printf("Gin mode: %s", gin.Mode())

	// Подключение к БД
	s, err := store.NewStore(cfg)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer func() {
		if cerr := s.Close(); cerr != nil {
			log.Printf("DB close error: %v", cerr)
		}
	}()

	// Роутер
	r := gin.Default()
	api.RegisterRoutes(r, s, cfg)

	// HTTP-сервер с таймаутами
	srv := &http.Server{
		Addr:           ":" + cfg.Port,
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   30 * time.Second,
		IdleTimeout:    60 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	// Грейсфул-шатдаун
	go func() {
		log.Printf("Server starting on port %s", cfg.Port)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal("Failed to start server:", err)
		}
	}()

	// Ожидание сигнала для корректного завершения
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	log.Println("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("Server shutdown error: %v", err)
	}
	log.Println("Server stopped gracefully")
}
