package api

import (
    "net/http"
    "strings"

    "github.com/gin-gonic/gin"
    "backend/config"
    "backend/internal/auth"
)

// AuthMiddleware валидирует JWT, извлекает user_id и role и кладёт их в контекст.
func AuthMiddleware(cfg *config.Config) gin.HandlerFunc {
    return func(c *gin.Context) {
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
            c.Abort()
            return
        }

        parts := strings.Split(authHeader, " ")
        if len(parts) != 2 || parts[0] != "Bearer" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
            c.Abort()
            return
        }

        claims, err := auth.ValidateToken(parts[1], cfg.JWTSecret)
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

// RequireAdmin — доступ только для роли admin (строго).
// В дальнейшем сюда можно добавлять дополнительные условия/атрибуты.
func RequireAdmin() gin.HandlerFunc {
    return func(c *gin.Context) {
        role, ok := c.Get("role")
        if !ok {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Role not found in context"})
            c.Abort()
            return
        }
        if roleStr, _ := role.(string); roleStr != "admin" {
            c.JSON(http.StatusForbidden, gin.H{"error": "Admin permissions required"})
            c.Abort()
            return
        }
        c.Next()
    }
}

// RequireEditor — на текущий момент пускает и admin, и editor,
// чтобы у редактора были те же права, что и у админа.
// Позже можно будет сузить до editor‑списка (или сделать матрицу прав).
func RequireEditor() gin.HandlerFunc {
    return func(c *gin.Context) {
        role, ok := c.Get("role")
        if !ok {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Role not found in context"})
            c.Abort()
            return
        }
        roleStr, _ := role.(string)
        if roleStr != "editor" && roleStr != "admin" {
            c.JSON(http.StatusForbidden, gin.H{"error": "Editor or Admin permissions required"})
            c.Abort()
            return
        }
        c.Next()
    }
}

// Обновлённый CORS‑мидлвар
func CORSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        origin := c.Request.Header.Get("Origin")

        allowedOrigins := []string{
            "http://localhost:3000",
            "http://localhost:5173",
            "http://localhost:8080",
            "http://127.0.0.1:3000",
            "http://127.0.0.1:5173",
            "http://127.0.0.1:8080",
        }

        originAllowed := false
        for _, allowed := range allowedOrigins {
            if origin == allowed {
                originAllowed = true
                break
            }
        }

        if originAllowed {
            c.Header("Access-Control-Allow-Origin", origin)
        } else {
            c.Header("Access-Control-Allow-Origin", "*")
        }

        c.Header("Access-Control-Allow-Credentials", "true")
        c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, PATCH")
        c.Header("Access-Control-Allow-Headers",
            "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, " +
                "Authorization, accept, origin, Cache-Control, X-Requested-With, " +
                "X-HTTP-Method-Override, X-Forwarded-For")
        c.Header("Access-Control-Expose-Headers",
            "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, " +
                "Content-Type, Cache-Control, Expires, Last-Modified")
        c.Header("Access-Control-Max-Age", "3600")

        if c.Request.Method == "OPTIONS" {
            c.Header("Access-Control-Allow-Origin", origin)
            c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, PATCH")
            c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Authorization, X-Requested-With")
            c.AbortWithStatus(http.StatusNoContent)
            return
        }

        c.Next()
    }
}

// Более строгий CORS для продакшена
func ProductionCORSMiddleware(allowedOrigins []string) gin.HandlerFunc {
    return func(c *gin.Context) {
        origin := c.Request.Header.Get("Origin")

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

// Логирование запросов (полезно при отладке CORS)
func RequestLogger() gin.HandlerFunc {
    return func(c *gin.Context) {
        // origin := c.Request.Header.Get("Origin")
        // method := c.Request.Method
        // log.Printf("CORS Request: %s %s from %s", method, c.Request.URL.Path, origin)
        c.Next()
    }
}
