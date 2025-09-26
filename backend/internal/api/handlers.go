package api

import (
    "log"
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"

    "backend/config"
    "backend/internal/auth"
    "backend/internal/models"
    "backend/internal/store"
)

type Handler struct {
    store *store.Store
    cfg   *config.Config
}

func NewHandler(store *store.Store, cfg *config.Config) *Handler {
    return &Handler{store: store, cfg: cfg}
}

// AdminLogin - логин для админа
func (h *Handler) AdminLogin(c *gin.Context) {
    var req models.AdminLoginRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
        return
    }

    user, err := h.store.GetUserByEmail(req.Email)
    if err != nil {
        log.Printf("Admin login failed for %s: user not found", req.Email)
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
        return
    }

    if user.Role != "admin" {
        log.Printf("Admin login failed for %s: not admin role (role=%s)", req.Email, user.Role)
        c.JSON(http.StatusForbidden, gin.H{"error": "Admin access required"})
        return
    }

    if user.Password != req.Password {
        log.Printf("Admin login failed for %s: password mismatch", req.Email)
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
        return
    }

    token, err := auth.GenerateToken(*user, h.cfg.JWTSecret)
    if err != nil {
        log.Printf("Failed to generate token for admin %s: %v", req.Email, err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
        return
    }

    log.Printf("Admin login successful for %s", req.Email)

    user.Password = ""

    c.JSON(http.StatusOK, models.LoginResponse{
        Token: token,
        User:  *user,
    })
}

// EditorLogin - логин для редактора
func (h *Handler) EditorLogin(c *gin.Context) {
    var req models.EditorLoginRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
        return
    }

    user, err := h.store.GetUserByEmail(req.Email)
    if err != nil {
        log.Printf("Editor login failed for %s: user not found", req.Email)
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
        return
    }

    if user.Role != "editor" && user.Role != "admin" {
        log.Printf("Editor login failed for %s: insufficient permissions (role=%s)", req.Email, user.Role)
        c.JSON(http.StatusForbidden, gin.H{"error": "Editor access required"})
        return
    }

    if user.Password != req.Password {
        log.Printf("Editor login failed for %s: password mismatch", req.Email)
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
        return
    }

    token, err := auth.GenerateToken(*user, h.cfg.JWTSecret)
    if err != nil {
        log.Printf("Failed to generate token for editor %s: %v", req.Email, err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
        return
    }

    log.Printf("Editor login successful for %s", req.Email)

    user.Password = ""

    c.JSON(http.StatusOK, models.LoginResponse{
        Token: token,
        User:  *user,
    })
}

// Обычный логин
func (h *Handler) Login(c *gin.Context) {
    var req models.LoginRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
        return
    }

    user, err := h.store.GetUserByEmail(req.Email)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
        return
    }

    if user.Password != req.Password {
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

// Fines
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
        OrdersTotal:          req.OrdersTotal,
        FinesAmountTotal:     req.FinesAmountTotal,
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
        OrdersTotal:          req.OrdersTotal,
        FinesAmountTotal:     req.FinesAmountTotal,
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

    c.Status(http.StatusNoContent)
}

// Evacuations
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
        Date:             req.Date,
        EvacuatorsCount:  req.EvacuatorsCount,
        TripsCount:       req.TripsCount,
        EvacuationsCount: req.EvacuationsCount,
        FineLotIncome:    req.FineLotIncome,
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

// Traffic lights
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

    c.Status(http.StatusNoContent)
}

// News
func (h *Handler) GetNews(c *gin.Context) {
    news, err := h.store.GetNews()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get news"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"news": news})
}

func (h *Handler) GetNewsByID(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    news, err := h.store.GetNewsByID(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "News not found"})
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

    c.Status(http.StatusNoContent)
}

// Services
func (h *Handler) GetServices(c *gin.Context) {
    services, err := h.store.GetServices()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get services"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"services": services})
}

func (h *Handler) GetServiceByID(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    service, err := h.store.GetServiceByID(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Service not found"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"service": service})
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

    c.Status(http.StatusNoContent)
}

// Team
func (h *Handler) GetTeam(c *gin.Context) {
    team, err := h.store.GetTeam()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get team"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"team": team})
}

func (h *Handler) GetTeamMemberByID(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    member, err := h.store.GetTeamMemberByID(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Team member not found"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"team_member": member})
}

// Update team member (admin/editor)
func (h *Handler) UpdateTeam(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil || id <= 0 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    var req models.TeamMember
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    member := &models.TeamMember{
        Name:       req.Name,
        Position:   req.Position,
        Experience: req.Experience,
        PhotoURL:   req.PhotoURL,
    }

    affected, err := h.store.UpdateTeam(id, member)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update team member"})
        return
    }
    if affected == 0 {
        c.JSON(http.StatusNotFound, gin.H{"error": "Team member not found"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Team member updated successfully"})
}

// Delete team member (admin/editor)
func (h *Handler) DeleteTeam(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil || id <= 0 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    affected, err := h.store.DeleteTeamByID(id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete team member"})
        return
    }
    if affected == 0 {
        c.JSON(http.StatusNotFound, gin.H{"error": "Team member not found"})
        return
    }

    c.Status(http.StatusNoContent)
}

// CreateTeam — создать участника команды (admin/editor)
func (h *Handler) CreateTeam(c *gin.Context) {
    var req models.TeamMember
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    member := &models.TeamMember{
        Name:       req.Name,
        Position:   req.Position,
        Experience: req.Experience,
        PhotoURL:   req.PhotoURL,
    }

    if err := h.store.CreateTeamMember(member); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create team member"})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"team_member": member})
}

// Projects
func (h *Handler) GetProjects(c *gin.Context) {
    projects, err := h.store.GetProjects()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get projects"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"projects": projects})
}

// CreateProject
func (h *Handler) CreateProject(c *gin.Context) {
    var req models.Project
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    p := &models.Project{
        Title:       req.Title,
        Description: req.Description,
        Category:    req.Category,
        Status:      req.Status,
    }

    if err := h.store.CreateProject(p); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create project"})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"project": p})
}

// UpdateProject
func (h *Handler) UpdateProject(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil || id <= 0 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    var req models.Project
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    p := &models.Project{
        Title:       req.Title,
        Description: req.Description,
        Category:    req.Category,
        Status:      req.Status,
    }

    if err := h.store.UpdateProject(id, p); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update project"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Project updated successfully"})
}

// DeleteProject
func (h *Handler) DeleteProject(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil || id <= 0 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    if err := h.store.DeleteProject(id); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete project"})
        return
    }

    c.Status(http.StatusNoContent)
}

// Stats
func (h *Handler) GetStats(c *gin.Context) {
    stats, err := h.store.GetStats()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get stats"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"stats": stats})
}

// Traffic
func (h *Handler) GetTraffic(c *gin.Context) {
    traffic, err := h.store.GetTraffic()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get traffic"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"traffic": traffic})
}

func (h *Handler) GetVacancies(c *gin.Context) {
    vacancies, err := h.store.GetVacancies()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get vacancies"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"vacancies": vacancies})
}

// Публичное получение вакансии по ID
func (h *Handler) GetVacancyByID(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil || id <= 0 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    vacancy, err := h.store.GetVacancyByID(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Vacancy not found"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"vacancy": vacancy})
}

// Создание вакансии (admin/editor)
func (h *Handler) CreateVacancy(c *gin.Context) {
    var req models.CreateVacancyRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    v := &models.Vacancy{
        Position:   req.Position,
        Experience: req.Experience,
        Salary:     req.Salary,
    }

    if err := h.store.CreateVacancy(v); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create vacancy"})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"vacancy": v})
}

// Обновление вакансии (admin/editor)
func (h *Handler) UpdateVacancy(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil || id <= 0 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    var req models.UpdateVacancyRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    v := &models.Vacancy{
        Position:   ptrOrEmpty(req.Position),
        Experience: ptrOrEmpty(req.Experience),
        Salary:     ptrOrEmpty(req.Salary),
    }

    if err := h.store.UpdateVacancy(id, v); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update vacancy"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Vacancy updated successfully"})
}

// Удаление вакансии (admin/editor)
func (h *Handler) DeleteVacancy(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil || id <= 0 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    if err := h.store.DeleteVacancy(id); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete vacancy"})
        return
    }

    c.Status(http.StatusNoContent)
}

// утилита: вернуть значение из *string либо пустую строку
func ptrOrEmpty(s *string) string {
    if s == nil {
        return ""
    }
    return *s
}