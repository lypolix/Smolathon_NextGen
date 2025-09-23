package main

import (
    "log"

    "github.com/gin-gonic/gin"
    "backend/config"
    "backend/internal/api"
    "backend/internal/store"
)

func main() {
    cfg := config.Load()

    s, err := store.NewStore(cfg)
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }
    defer s.Close()

    r := gin.Default()
    
    api.RegisterRoutes(r, s, cfg)

    log.Printf("Server starting on port %s", cfg.Port)
    if err := r.Run(":" + cfg.Port); err != nil {
        log.Fatal("Failed to start server:", err)
    }
}
