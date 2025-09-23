package api

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "golang.org/x/crypto/bcrypt"
    "backend/config"
    "backend/internal/auth"
    "backend/internal/models"
    "backend/internal/store"
)

type Handler struct {
    store *store.Store
    cfg   *config.Config
}

func NewHandler(s *store.Store, cfg *config.Config) *Handler {
    return &Handler{
        store: s,
        cfg:   cfg,
    }
}

// Auth handlers (обновленные - обычная авторизация по email/password)
func (h *Handler) AdminLogin(c *gin.Context) {
    var req models.AdminLoginRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    user, err := h.store.GetUserByEmail(req.Email)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
        return
    }

    if user.Role != "admin" {
        c.JSON(http.StatusForbidden, gin.H{"error": "Access denied: admin role required"})
        return
    }

    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
        return
    }

    token, err := auth.GenerateToken(*user, h.cfg.JWTSecret)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
        return
    }

    user.Password = ""
    c.JSON(http.StatusOK, models.LoginResponse{
        Token: token,
        User:  *user,
    })
}

func (h *Handler) EditorLogin(c *gin.Context) {
    var req models.EditorLoginRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    user, err := h.store.GetUserByEmail(req.Email)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
        return
    }

    if user.Role != "editor" {
        c.JSON(http.StatusForbidden, gin.H{"error": "Access denied: editor role required"})
        return
    }

    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
        return
    }

    token, err := auth.GenerateToken(*user, h.cfg.JWTSecret)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
        return
    }

    user.Password = ""
    c.JSON(http.StatusOK, models.LoginResponse{
        Token: token,
        User:  *user,
    })
}

// Общий логин (опционально, если нужен единый эндпоинт)
func (h *Handler) Login(c *gin.Context) {
    var req models.LoginRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    user, err := h.store.GetUserByEmail(req.Email)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
        return
    }

    if user.Role != "admin" && user.Role != "editor" {
        c.JSON(http.StatusForbidden, gin.H{"error": "Access denied: admin or editor role required"})
        return
    }

    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
        return
    }

    token, err := auth.GenerateToken(*user, h.cfg.JWTSecret)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
        return
    }

    user.Password = ""
    c.JSON(http.StatusOK, models.LoginResponse{
        Token: token,
        User:  *user,
    })
}

// Fine handlers (без изменений)
func (h *Handler) GetFines(c *gin.Context) {
    fines, err := h.store.GetFines()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get fines"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"fines": fines})
}

func (h *Handler) CreateFine(c *gin.Context) {
    var req models.CreateFineRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    fine := &models.Fine{
        Date:                 req.Date,
        ViolationsTotal:      req.ViolationsTotal,
        OrdersTotal:         req.OrdersTotal,
        FinesAmountTotal:    req.FinesAmountTotal,
        CollectedAmountTotal: req.CollectedAmountTotal,
    }

    if err := h.store.CreateFine(fine); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create fine"})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"fine": fine})
}

func (h *Handler) UpdateFine(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    var req models.UpdateFineRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    fine := &models.Fine{
        Date:                 req.Date,
        ViolationsTotal:      req.ViolationsTotal,
        OrdersTotal:         req.OrdersTotal,
        FinesAmountTotal:    req.FinesAmountTotal,
        CollectedAmountTotal: req.CollectedAmountTotal,
    }

    if err := h.store.UpdateFine(id, fine); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update fine"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Fine updated successfully"})
}

func (h *Handler) DeleteFine(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    if err := h.store.DeleteFine(id); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete fine"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Fine deleted successfully"})
}

// Evacuation handlers (без изменений)
func (h *Handler) GetEvacuations(c *gin.Context) {
    evacuations, err := h.store.GetEvacuations()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get evacuations"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"evacuations": evacuations})
}

func (h *Handler) GetEvacuationRoutes(c *gin.Context) {
    routes, err := h.store.GetEvacuationRoutes()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get evacuation routes"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"evacuation_routes": routes})
}

func (h *Handler) CreateEvacuation(c *gin.Context) {
    var req models.CreateEvacuationRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    evacuation := &models.Evacuation{
        Date:            req.Date,
        EvacuatorsCount: req.EvacuatorsCount,
        TripsCount:      req.TripsCount,
        EvacuationsCount: req.EvacuationsCount,
        FineLotIncome:   req.FineLotIncome,
    }

    if err := h.store.CreateEvacuation(evacuation); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create evacuation"})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"evacuation": evacuation})
}

func (h *Handler) CreateEvacuationRoute(c *gin.Context) {
    var req models.CreateEvacuationRouteRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    route := &models.EvacuationRoute{
        Year:  req.Year,
        Month: req.Month,
        Route: req.Route,
    }

    if err := h.store.CreateEvacuationRoute(route); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create evacuation route"})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"evacuation_route": route})
}

// Traffic Light handlers (без изменений)
func (h *Handler) GetTrafficLights(c *gin.Context) {
    lights, err := h.store.GetTrafficLights()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get traffic lights"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"traffic_lights": lights})
}

func (h *Handler) CreateTrafficLight(c *gin.Context) {
    var req models.CreateTrafficLightRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if req.Status == "" {
        req.Status = "active"
    }

    light := &models.TrafficLight{
        Address:     req.Address,
        LightType:   req.LightType,
        InstallYear: req.InstallYear,
        Status:      req.Status,
    }

    if err := h.store.CreateTrafficLight(light); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create traffic light"})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"traffic_light": light})
}

func (h *Handler) UpdateTrafficLight(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    var req models.UpdateTrafficLightRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    light := &models.TrafficLight{
        Address:     req.Address,
        LightType:   req.LightType,
        InstallYear: req.InstallYear,
        Status:      req.Status,
    }

    if err := h.store.UpdateTrafficLight(id, light); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update traffic light"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Traffic light updated successfully"})
}

func (h *Handler) DeleteTrafficLight(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    if err := h.store.DeleteTrafficLight(id); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete traffic light"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Traffic light deleted successfully"})
}

// News handlers (без изменений)
func (h *Handler) GetNews(c *gin.Context) {
    news, err := h.store.GetNews()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get news"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"news": news})
}

func (h *Handler) CreateNews(c *gin.Context) {
    var req models.CreateNewsRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    news := &models.News{
        Title:   req.Title,
        Content: req.Content,
        Tag:     req.Tag,
    }

    if err := h.store.CreateNews(news); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create news"})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"news": news})
}

func (h *Handler) UpdateNews(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    var req models.UpdateNewsRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    news := &models.News{
        Title:   req.Title,
        Content: req.Content,
        Tag:     req.Tag,
    }

    if err := h.store.UpdateNews(id, news); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update news"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "News updated successfully"})
}

func (h *Handler) DeleteNews(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    if err := h.store.DeleteNews(id); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete news"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "News deleted successfully"})
}

// Services handlers (без изменений)
func (h *Handler) GetServices(c *gin.Context) {
    services, err := h.store.GetServices()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get services"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"services": services})
}

func (h *Handler) CreateService(c *gin.Context) {
    var req models.CreateServiceRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    service := &models.Service{
        Title:       req.Title,
        Description: req.Description,
        Price:       req.Price,
        Category:    req.Category,
        IconURL:     req.IconURL,
    }

    if err := h.store.CreateService(service); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create service"})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"service": service})
}

func (h *Handler) UpdateService(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    var req models.UpdateServiceRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    service := &models.Service{
        Title:       req.Title,
        Description: req.Description,
        Price:       req.Price,
        Category:    req.Category,
        IconURL:     req.IconURL,
    }

    if err := h.store.UpdateService(id, service); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update service"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Service updated successfully"})
}

func (h *Handler) DeleteService(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    if err := h.store.DeleteService(id); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete service"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Service deleted successfully"})
}

// Team handlers (без изменений)
func (h *Handler) GetTeam(c *gin.Context) {
    team, err := h.store.GetTeam()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get team"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"team": team})
}

// Projects handlers (без изменений)
func (h *Handler) GetProjects(c *gin.Context) {
    projects, err := h.store.GetProjects()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get projects"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"projects": projects})
}

// Stats handlers (без изменений)
func (h *Handler) GetStats(c *gin.Context) {
    stats, err := h.store.GetStats()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get stats"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"stats": stats})
}

// Traffic handlers (без изменений)
func (h *Handler) GetTraffic(c *gin.Context) {
    traffic, err := h.store.GetTraffic()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get traffic"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"traffic": traffic})
}
