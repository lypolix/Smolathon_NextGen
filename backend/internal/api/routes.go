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
        // Новости
        api.GET("/news", h.GetNews)
        api.GET("/news/:id", h.GetNewsByID)

        // Услуги
        api.GET("/services", h.GetServices)
        api.GET("/services/:id", h.GetServiceByID)

        // Команда
        api.GET("/team", h.GetTeam)
        api.GET("/team/:id", h.GetTeamMemberByID)

        // Проекты
        api.GET("/projects", h.GetProjects)

        // Статистика/трафик
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

    // Админские маршруты (только админ)
    admin := r.Group("/api/admin", AuthMiddleware(cfg), RequireAdmin())
    {
        // Зеркальные GET для админских страниц (чтение с авторизацией)
        admin.GET("/news", h.GetNews)
        admin.GET("/news/:id", h.GetNewsByID)

        admin.GET("/services", h.GetServices)
        admin.GET("/services/:id", h.GetServiceByID)

        admin.GET("/team", h.GetTeam)
        admin.GET("/team/:id", h.GetTeamMemberByID)

        admin.GET("/projects", h.GetProjects)

        admin.GET("/fines", h.GetFines)
        admin.GET("/evacuations", h.GetEvacuations)
        admin.GET("/evacuation-routes", h.GetEvacuationRoutes)
        admin.GET("/traffic-lights", h.GetTrafficLights)

        admin.GET("/vacancies", h.GetVacancies)
        admin.GET("/vacancies/:id", h.GetVacancyByID)

        // Новости — CRUD
        admin.POST("/news", h.CreateNews)
        admin.PUT("/news/:id", h.UpdateNews)
        admin.DELETE("/news/:id", h.DeleteNews)

        // Услуги — CRUD
        admin.POST("/services", h.CreateService)
        admin.PUT("/services/:id", h.UpdateService)
        admin.DELETE("/services/:id", h.DeleteService)

        // Штрафы — CRUD
        admin.POST("/fines", h.CreateFine)
        admin.PUT("/fines/:id", h.UpdateFine)
        admin.DELETE("/fines/:id", h.DeleteFine)

        // Эвакуация — CRUD
        admin.POST("/evacuations", h.CreateEvacuation)
        admin.POST("/evacuation-routes", h.CreateEvacuationRoute)

        // Светофоры — CRUD
        admin.POST("/traffic-lights", h.CreateTrafficLight)
        admin.PUT("/traffic-lights/:id", h.UpdateTrafficLight)
        admin.DELETE("/traffic-lights/:id", h.DeleteTrafficLight)

        // Команда — CRUD
        admin.POST("/team", h.CreateTeam)        // если реализовано
        admin.PUT("/team/:id", h.UpdateTeam)     // если реализовано
        admin.DELETE("/team/:id", h.DeleteTeam)

        // Проекты — CRUD
        admin.POST("/projects", h.CreateProject)
        admin.PUT("/projects/:id", h.UpdateProject)
        admin.DELETE("/projects/:id", h.DeleteProject)

        // Вакансии — CRUD
        admin.POST("/vacancies", h.CreateVacancy)
        admin.PUT("/vacancies/:id", h.UpdateVacancy)
        admin.DELETE("/vacancies/:id", h.DeleteVacancy)
    }

    // Редакторские маршруты (редактор/админ)
    editor := r.Group("/api/editor", AuthMiddleware(cfg), RequireEditor())
    {
        // Зеркальные GET для редакторских страниц (чтение с авторизацией)
        editor.GET("/news", h.GetNews)
        editor.GET("/news/:id", h.GetNewsByID)

        editor.GET("/services", h.GetServices)
        editor.GET("/services/:id", h.GetServiceByID)

        editor.GET("/team", h.GetTeam)
        editor.GET("/team/:id", h.GetTeamMemberByID)

        editor.GET("/projects", h.GetProjects)

        editor.GET("/fines", h.GetFines)
        editor.GET("/evacuations", h.GetEvacuations)
        editor.GET("/evacuation-routes", h.GetEvacuationRoutes)
        editor.GET("/traffic-lights", h.GetTrafficLights)

        editor.GET("/vacancies", h.GetVacancies)
        editor.GET("/vacancies/:id", h.GetVacancyByID)

        // Новости — CRUD
        editor.POST("/news", h.CreateNews)
        editor.PUT("/news/:id", h.UpdateNews)
        editor.DELETE("/news/:id", h.DeleteNews)

        // Услуги — CRUD
        editor.POST("/services", h.CreateService)
        editor.PUT("/services/:id", h.UpdateService)
        editor.DELETE("/services/:id", h.DeleteService)

        // Штрафы — CRUD
        editor.POST("/fines", h.CreateFine)
        editor.PUT("/fines/:id", h.UpdateFine)
        editor.DELETE("/fines/:id", h.DeleteFine)

        // Эвакуация — CRUD
        editor.POST("/evacuations", h.CreateEvacuation)
        editor.POST("/evacuation-routes", h.CreateEvacuationRoute)

        // Светофоры — CRUD
        editor.POST("/traffic-lights", h.CreateTrafficLight)
        editor.PUT("/traffic-lights/:id", h.UpdateTrafficLight)
        editor.DELETE("/traffic-lights/:id", h.DeleteTrafficLight)

        // Команда — CRUD
        editor.POST("/team", h.CreateTeam)        // если реализовано
        editor.PUT("/team/:id", h.UpdateTeam)     // если реализовано
        editor.DELETE("/team/:id", h.DeleteTeam)

        // Проекты — CRUD
        editor.POST("/projects", h.CreateProject)
        editor.PUT("/projects/:id", h.UpdateProject)
        editor.DELETE("/projects/:id", h.DeleteProject)

        // Вакансии — CRUD
        editor.POST("/vacancies", h.CreateVacancy)
        editor.PUT("/vacancies/:id", h.UpdateVacancy)
        editor.DELETE("/vacancies/:id", h.DeleteVacancy)
    }
}
