package api

import (
    "github.com/gin-gonic/gin"
    "backend/config"
    "backend/internal/store"
)

func RegisterRoutes(r *gin.Engine, s *store.Store, cfg *config.Config) {
    r.Use(CORSMiddleware())
    
    h := NewHandler(s, cfg)

    // Аутентификация - раздельные эндпоинты + общий
    auth := r.Group("/api/auth")
    {
        auth.POST("/admin/login", h.AdminLogin)
        auth.POST("/editor/login", h.EditorLogin)
        auth.POST("/login", h.Login) // общий логин
    }

    // Публичные маршруты (доступны всем без авторизации)
    api := r.Group("/api")
    {
        api.GET("/news", h.GetNews)
        api.GET("/news/:id", h.GetNewsByID)           // Получить новость по ID
        api.GET("/services", h.GetServices)
        api.GET("/services/:id", h.GetServiceByID)    // Получить услугу по ID
        api.GET("/team", h.GetTeam)
        api.GET("/team/:id", h.GetTeamMemberByID)     // Получить участника команды по ID
        api.GET("/projects", h.GetProjects)
        api.GET("/stats", h.GetStats)
        api.GET("/traffic", h.GetTraffic)
        
        // Данные из Excel - публичные
        api.GET("/fines", h.GetFines)
        api.GET("/evacuations", h.GetEvacuations)
        api.GET("/evacuation-routes", h.GetEvacuationRoutes)
        api.GET("/traffic-lights", h.GetTrafficLights)
    }

    // Админские маршруты
    admin := r.Group("/api/admin", AuthMiddleware(cfg), RequireRole("admin"))
    {
        // Новости
        admin.POST("/news", h.CreateNews)
        admin.PUT("/news/:id", h.UpdateNews)
        admin.DELETE("/news/:id", h.DeleteNews)

        // Услуги
        admin.POST("/services", h.CreateService)
        admin.PUT("/services/:id", h.UpdateService)
        admin.DELETE("/services/:id", h.DeleteService)

        // Штрафы
        admin.POST("/fines", h.CreateFine)
        admin.PUT("/fines/:id", h.UpdateFine)
        admin.DELETE("/fines/:id", h.DeleteFine)

        // Эвакуация
        admin.POST("/evacuations", h.CreateEvacuation)
        admin.POST("/evacuation-routes", h.CreateEvacuationRoute)

        // Светофоры
        admin.POST("/traffic-lights", h.CreateTrafficLight)
        admin.PUT("/traffic-lights/:id", h.UpdateTrafficLight)
        admin.DELETE("/traffic-lights/:id", h.DeleteTrafficLight)

        // Команда
        admin.POST("/team", func(c *gin.Context) { c.JSON(200, gin.H{"message": "Create team member"}) })
        admin.PUT("/team/:id", func(c *gin.Context) { c.JSON(200, gin.H{"message": "Update team member"}) })
        admin.DELETE("/team/:id", func(c *gin.Context) { c.JSON(200, gin.H{"message": "Delete team member"}) })

        // Проекты
        admin.POST("/projects", func(c *gin.Context) { c.JSON(200, gin.H{"message": "Create project"}) })
        admin.PUT("/projects/:id", func(c *gin.Context) { c.JSON(200, gin.H{"message": "Update project"}) })
        admin.DELETE("/projects/:id", func(c *gin.Context) { c.JSON(200, gin.H{"message": "Delete project"}) })
    }

    // Редакторские маршруты (ограниченные права)
    editor := r.Group("/api/editor", AuthMiddleware(cfg), RequireRole("editor"))
    {
        editor.POST("/news", h.CreateNews)
        editor.PUT("/news/:id", h.UpdateNews)
        editor.POST("/fines", h.CreateFine)
        editor.PUT("/fines/:id", h.UpdateFine)
        editor.POST("/evacuations", h.CreateEvacuation)
        editor.POST("/evacuation-routes", h.CreateEvacuationRoute)
        editor.POST("/traffic-lights", h.CreateTrafficLight)
        editor.PUT("/traffic-lights/:id", h.UpdateTrafficLight)
    }
}
