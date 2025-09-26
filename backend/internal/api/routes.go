package api

import (
    "github.com/gin-gonic/gin"
    "backend/config"
    "backend/internal/store"
)

func RegisterRoutes(r *gin.Engine, s *store.Store, cfg *config.Config) {
    r.Use(CORSMiddleware())

    h := NewHandler(s, cfg)

    // Аутентификация — раздельные эндпоинты + общий
    auth := r.Group("/api/auth")
    {
        auth.POST("/admin/login", h.AdminLogin)
        auth.POST("/editor/login", h.EditorLogin)
        auth.POST("/login", h.Login) // общий логин
    }

    // Публичные маршруты (без авторизации)
    api := r.Group("/api")
    {
        api.GET("/news", h.GetNews)
        api.GET("/news/:id", h.GetNewsByID)
        api.GET("/services", h.GetServices)
        api.GET("/services/:id", h.GetServiceByID)
        api.GET("/team", h.GetTeam)
        api.GET("/team/:id", h.GetTeamMemberByID)
        api.GET("/projects", h.GetProjects)
        api.GET("/stats", h.GetStats)
        api.GET("/traffic", h.GetTraffic)

        // Данные из Excel — публичные
        api.GET("/fines", h.GetFines)
        api.GET("/evacuations", h.GetEvacuations)
        api.GET("/evacuation-routes", h.GetEvacuationRoutes)
        api.GET("/traffic-lights", h.GetTrafficLights)

        // Вакансии — публичное чтение
        api.GET("/vacancies", h.GetVacancies)
        api.GET("/vacancies/:id", h.GetVacancyByID)
    }

    // Админские маршруты (только админ; в дальнейшем можно расширять отдельно)
    admin := r.Group("/api/admin", AuthMiddleware(cfg), RequireAdmin())
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
        admin.POST("/team", h.CreateTeam)        // если реализовано
        admin.PUT("/team/:id", h.UpdateTeam)     // если реализовано
        admin.DELETE("/team/:id", h.DeleteTeam)

        // Проекты
        admin.POST("/projects", h.CreateProject)
        admin.PUT("/projects/:id", h.UpdateProject)
        admin.DELETE("/projects/:id", h.DeleteProject)

        // Вакансии — полный CRUD
        admin.POST("/vacancies", h.CreateVacancy)
        admin.PUT("/vacancies/:id", h.UpdateVacancy)
        admin.DELETE("/vacancies/:id", h.DeleteVacancy)
    }

    // Редакторские маршруты: на текущий момент совпадают с админскими
    editor := r.Group("/api/editor", AuthMiddleware(cfg), RequireEditor())
    {
        // Новости
        editor.POST("/news", h.CreateNews)
        editor.PUT("/news/:id", h.UpdateNews)
        editor.DELETE("/news/:id", h.DeleteNews)

        // Услуги
        editor.POST("/services", h.CreateService)
        editor.PUT("/services/:id", h.UpdateService)
        editor.DELETE("/services/:id", h.DeleteService)

        // Штрафы
        editor.POST("/fines", h.CreateFine)
        editor.PUT("/fines/:id", h.UpdateFine)
        editor.DELETE("/fines/:id", h.DeleteFine)

        // Эвакуация
        editor.POST("/evacuations", h.CreateEvacuation)
        editor.POST("/evacuation-routes", h.CreateEvacuationRoute)

        // Светофоры
        editor.POST("/traffic-lights", h.CreateTrafficLight)
        editor.PUT("/traffic-lights/:id", h.UpdateTrafficLight)
        editor.DELETE("/traffic-lights/:id", h.DeleteTrafficLight)

        // Команда
        editor.POST("/team", h.CreateTeam)        // если реализовано
        editor.PUT("/team/:id", h.UpdateTeam)     // если реализовано
        editor.DELETE("/team/:id", h.DeleteTeam)

        // Проекты
        editor.POST("/projects", h.CreateProject)
        editor.PUT("/projects/:id", h.UpdateProject)
        editor.DELETE("/projects/:id", h.DeleteProject)

        // Вакансии — полный CRUD
        editor.POST("/vacancies", h.CreateVacancy)
        editor.PUT("/vacancies/:id", h.UpdateVacancy)
        editor.DELETE("/vacancies/:id", h.DeleteVacancy)
    }
}
