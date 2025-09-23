package api

import (
    "net/http"
    "strings"

    "github.com/gin-gonic/gin"
    "backend/config"
    "backend/internal/auth"
)

func AuthMiddleware(cfg *config.Config) gin.HandlerFunc {
    return func(c *gin.Context) {
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
            c.Abort()
            return
        }

        bearerToken := strings.Split(authHeader, " ")
        if len(bearerToken) != 2 || bearerToken[0] != "Bearer" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
            c.Abort()
            return
        }

        claims, err := auth.ValidateToken(bearerToken[1], cfg.JWTSecret)
        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            c.Abort()
            return
        }

        c.Set("user_id", claims.UserID)
        c.Set("role", claims.Role)
        c.Next()
    }
}

func RequireRole(role string) gin.HandlerFunc {
    return func(c *gin.Context) {
        userRole, exists := c.Get("role")
        if !exists {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Role not found in context"})
            c.Abort()
            return
        }

        if userRole != role && userRole != "admin" {
            c.JSON(http.StatusForbidden, gin.H{"error": "Insufficient permissions"})
            c.Abort()
            return
        }

        c.Next()
    }
}

// Обновленный CORS мидлвар с правильными настройками
func CORSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        origin := c.Request.Header.Get("Origin")
        
        // Разрешенные origins (можно настроить через конфиг)
        allowedOrigins := []string{
            "http://localhost:3000",  // React dev server
            "http://localhost:5173",  // Vite dev server
            "http://localhost:8080",  // Возможный фронт на том же порту
            "http://127.0.0.1:3000",
            "http://127.0.0.1:5173",
            "http://127.0.0.1:8080",
            // Добавь свой домен в продакшн
            // "https://yourdomain.com",
        }

        // Проверка origin
        originAllowed := false
        for _, allowed := range allowedOrigins {
            if origin == allowed {
                originAllowed = true
                break
            }
        }

        // Устанавливаем заголовки CORS
        if originAllowed {
            c.Header("Access-Control-Allow-Origin", origin)
        } else {
            // В разработке можно разрешить все origins
            c.Header("Access-Control-Allow-Origin", "*")
        }

        c.Header("Access-Control-Allow-Credentials", "true")
        c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, PATCH")
        c.Header("Access-Control-Allow-Headers", 
            "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, "+
            "Authorization, accept, origin, Cache-Control, X-Requested-With, "+
            "X-HTTP-Method-Override, X-Forwarded-For")
        c.Header("Access-Control-Expose-Headers", 
            "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, "+
            "Content-Type, Cache-Control, Expires, Last-Modified")
        c.Header("Access-Control-Max-Age", "3600")

        // Обработка preflight запросов
        if c.Request.Method == "OPTIONS" {
            c.Header("Access-Control-Allow-Origin", origin)
            c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, PATCH")
            c.Header("Access-Control-Allow-Headers", 
                "Origin, Content-Type, Accept, Authorization, X-Requested-With")
            c.AbortWithStatus(http.StatusNoContent)
            return
        }

        c.Next()
    }
}

// Альтернативный CORS мидлвар для продакшн (более строгий)
func ProductionCORSMiddleware(allowedOrigins []string) gin.HandlerFunc {
    return func(c *gin.Context) {
        origin := c.Request.Header.Get("Origin")
        
        // Проверяем только разрешенные origins
        originAllowed := false
        for _, allowed := range allowedOrigins {
            if origin == allowed {
                originAllowed = true
                c.Header("Access-Control-Allow-Origin", origin)
                break
            }
        }

        if !originAllowed {
            c.AbortWithStatus(http.StatusForbidden)
            return
        }

        c.Header("Access-Control-Allow-Credentials", "true")
        c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
        c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")
        c.Header("Access-Control-Max-Age", "86400")

        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(http.StatusNoContent)
            return
        }

        c.Next()
    }
}

// Мидлвар для логирования запросов (полезен для отладки CORS)
func RequestLogger() gin.HandlerFunc {
    return func(c *gin.Context) {
        // Логируем CORS-related заголовки
        origin := c.Request.Header.Get("Origin")
        method := c.Request.Method
        
        if origin != "" || method == "OPTIONS" {
            // Можно добавить логирование в файл или stdout
            // log.Printf("CORS Request: %s %s from %s", method, c.Request.URL.Path, origin)
        }
        
        c.Next()
    }
}
