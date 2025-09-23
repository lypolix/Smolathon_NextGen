package store

import (
    "database/sql"
    "fmt"
    "log"
    "time"

    _ "github.com/lib/pq"
    "backend/config"
    "backend/internal/models"
)

type Store struct {
    db *sql.DB
}

func NewStore(cfg *config.Config) (*Store, error) {
    dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)
    
    db, err := sql.Open("postgres", dsn)
    if err != nil {
        return nil, err
    }

    if err := db.Ping(); err != nil {
        return nil, err
    }

    log.Println("Connected to PostgreSQL")
    
    return &Store{db: db}, nil
}

func (s *Store) Close() error {
    return s.db.Close()
}

func (s *Store) GetDB() *sql.DB {
    return s.db
}

// User methods (вернули обратно)
func (s *Store) GetUserByEmail(email string) (*models.User, error) {
    query := `SELECT id, email, password, role, is_active, created_at, updated_at FROM users WHERE email = $1 AND is_active = true`
    
    user := &models.User{}
    err := s.db.QueryRow(query, email).Scan(
        &user.ID, &user.Email, &user.Password, &user.Role, &user.IsActive, &user.CreatedAt, &user.UpdatedAt,
    )
    
    if err != nil {
        return nil, err
    }
    
    return user, nil
}

func (s *Store) CreateUser(user *models.User) error {
    query := `INSERT INTO users (email, password, role, is_active, created_at, updated_at) 
              VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
    
    now := time.Now()
    user.CreatedAt = now
    user.UpdatedAt = now
    
    return s.db.QueryRow(query, user.Email, user.Password, user.Role, user.IsActive, user.CreatedAt, user.UpdatedAt).Scan(&user.ID)
}

// Fine methods (без изменений)
func (s *Store) GetFines() ([]models.Fine, error) {
    query := `SELECT id, date, violations_total, orders_total, fines_amount_total, collected_amount_total, created_at, updated_at 
              FROM fines ORDER BY date DESC`
    
    rows, err := s.db.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var fines []models.Fine
    for rows.Next() {
        var f models.Fine
        if err := rows.Scan(&f.ID, &f.Date, &f.ViolationsTotal, &f.OrdersTotal, &f.FinesAmountTotal, &f.CollectedAmountTotal, &f.CreatedAt, &f.UpdatedAt); err != nil {
            return nil, err
        }
        fines = append(fines, f)
    }
    
    return fines, nil
}

func (s *Store) CreateFine(fine *models.Fine) error {
    query := `INSERT INTO fines (date, violations_total, orders_total, fines_amount_total, collected_amount_total, created_at, updated_at) 
              VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`
    
    now := time.Now()
    fine.CreatedAt = now
    fine.UpdatedAt = now
    
    return s.db.QueryRow(query, fine.Date, fine.ViolationsTotal, fine.OrdersTotal, fine.FinesAmountTotal, fine.CollectedAmountTotal, fine.CreatedAt, fine.UpdatedAt).Scan(&fine.ID)
}

func (s *Store) UpdateFine(id int, fine *models.Fine) error {
    query := `UPDATE fines SET date = $2, violations_total = $3, orders_total = $4, fines_amount_total = $5, collected_amount_total = $6, updated_at = $7 WHERE id = $1`
    
    fine.UpdatedAt = time.Now()
    _, err := s.db.Exec(query, id, fine.Date, fine.ViolationsTotal, fine.OrdersTotal, fine.FinesAmountTotal, fine.CollectedAmountTotal, fine.UpdatedAt)
    return err
}

func (s *Store) DeleteFine(id int) error {
    query := `DELETE FROM fines WHERE id = $1`
    _, err := s.db.Exec(query, id)
    return err
}

// Evacuation methods (без изменений)
func (s *Store) GetEvacuations() ([]models.Evacuation, error) {
    query := `SELECT id, date, evacuators_count, trips_count, evacuations_count, fine_lot_income, created_at, updated_at 
              FROM evacuations ORDER BY date DESC`
    
    rows, err := s.db.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var evacuations []models.Evacuation
    for rows.Next() {
        var e models.Evacuation
        if err := rows.Scan(&e.ID, &e.Date, &e.EvacuatorsCount, &e.TripsCount, &e.EvacuationsCount, &e.FineLotIncome, &e.CreatedAt, &e.UpdatedAt); err != nil {
            return nil, err
        }
        evacuations = append(evacuations, e)
    }
    
    return evacuations, nil
}

func (s *Store) CreateEvacuation(evacuation *models.Evacuation) error {
    query := `INSERT INTO evacuations (date, evacuators_count, trips_count, evacuations_count, fine_lot_income, created_at, updated_at) 
              VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`
    
    now := time.Now()
    evacuation.CreatedAt = now
    evacuation.UpdatedAt = now
    
    return s.db.QueryRow(query, evacuation.Date, evacuation.EvacuatorsCount, evacuation.TripsCount, evacuation.EvacuationsCount, evacuation.FineLotIncome, evacuation.CreatedAt, evacuation.UpdatedAt).Scan(&evacuation.ID)
}

func (s *Store) GetEvacuationRoutes() ([]models.EvacuationRoute, error) {
    query := `SELECT id, year, month, route, created_at, updated_at FROM evacuation_routes ORDER BY year DESC, month`
    
    rows, err := s.db.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var routes []models.EvacuationRoute
    for rows.Next() {
        var r models.EvacuationRoute
        if err := rows.Scan(&r.ID, &r.Year, &r.Month, &r.Route, &r.CreatedAt, &r.UpdatedAt); err != nil {
            return nil, err
        }
        routes = append(routes, r)
    }
    
    return routes, nil
}

func (s *Store) CreateEvacuationRoute(route *models.EvacuationRoute) error {
    query := `INSERT INTO evacuation_routes (year, month, route, created_at, updated_at) 
              VALUES ($1, $2, $3, $4, $5) RETURNING id`
    
    now := time.Now()
    route.CreatedAt = now
    route.UpdatedAt = now
    
    return s.db.QueryRow(query, route.Year, route.Month, route.Route, route.CreatedAt, route.UpdatedAt).Scan(&route.ID)
}

// Traffic Light methods (без изменений)
func (s *Store) GetTrafficLights() ([]models.TrafficLight, error) {
    query := `SELECT id, address, light_type, install_year, status, created_at, updated_at 
              FROM traffic_lights ORDER BY install_year DESC`
    
    rows, err := s.db.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var lights []models.TrafficLight
    for rows.Next() {
        var t models.TrafficLight
        if err := rows.Scan(&t.ID, &t.Address, &t.LightType, &t.InstallYear, &t.Status, &t.CreatedAt, &t.UpdatedAt); err != nil {
            return nil, err
        }
        lights = append(lights, t)
    }
    
    return lights, nil
}

func (s *Store) CreateTrafficLight(light *models.TrafficLight) error {
    query := `INSERT INTO traffic_lights (address, light_type, install_year, status, created_at, updated_at) 
              VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
    
    now := time.Now()
    light.CreatedAt = now
    light.UpdatedAt = now
    
    return s.db.QueryRow(query, light.Address, light.LightType, light.InstallYear, light.Status, light.CreatedAt, light.UpdatedAt).Scan(&light.ID)
}

func (s *Store) UpdateTrafficLight(id int, light *models.TrafficLight) error {
    query := `UPDATE traffic_lights SET address = $2, light_type = $3, install_year = $4, status = $5, updated_at = $6 WHERE id = $1`
    
    light.UpdatedAt = time.Now()
    _, err := s.db.Exec(query, id, light.Address, light.LightType, light.InstallYear, light.Status, light.UpdatedAt)
    return err
}

func (s *Store) DeleteTrafficLight(id int) error {
    query := `DELETE FROM traffic_lights WHERE id = $1`
    _, err := s.db.Exec(query, id)
    return err
}

// News methods (без изменений)
func (s *Store) GetNews() ([]models.News, error) {
    query := `SELECT id, title, content, tag, date, created_at, updated_at FROM news ORDER BY date DESC`
    
    rows, err := s.db.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var news []models.News
    for rows.Next() {
        var n models.News
        if err := rows.Scan(&n.ID, &n.Title, &n.Content, &n.Tag, &n.Date, &n.CreatedAt, &n.UpdatedAt); err != nil {
            return nil, err
        }
        news = append(news, n)
    }
    
    return news, nil
}

func (s *Store) CreateNews(news *models.News) error {
    query := `INSERT INTO news (title, content, tag, date, created_at, updated_at) 
              VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
    
    now := time.Now()
    news.Date = now
    news.CreatedAt = now
    news.UpdatedAt = now
    
    return s.db.QueryRow(query, news.Title, news.Content, news.Tag, news.Date, news.CreatedAt, news.UpdatedAt).Scan(&news.ID)
}

func (s *Store) UpdateNews(id int, news *models.News) error {
    query := `UPDATE news SET title = $2, content = $3, tag = $4, updated_at = $5 WHERE id = $1`
    
    news.UpdatedAt = time.Now()
    _, err := s.db.Exec(query, id, news.Title, news.Content, news.Tag, news.UpdatedAt)
    return err
}

func (s *Store) DeleteNews(id int) error {
    query := `DELETE FROM news WHERE id = $1`
    _, err := s.db.Exec(query, id)
    return err
}

// Services methods (без изменений)
func (s *Store) GetServices() ([]models.Service, error) {
    query := `SELECT id, title, description, price, category, icon_url, created_at, updated_at FROM services`
    
    rows, err := s.db.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var services []models.Service
    for rows.Next() {
        var service models.Service
        if err := rows.Scan(&service.ID, &service.Title, &service.Description, &service.Price, &service.Category, &service.IconURL, &service.CreatedAt, &service.UpdatedAt); err != nil {
            return nil, err
        }
        services = append(services, service)
    }
    
    return services, nil
}

func (s *Store) CreateService(service *models.Service) error {
    query := `INSERT INTO services (title, description, price, category, icon_url, created_at, updated_at) 
              VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`
    
    now := time.Now()
    service.CreatedAt = now
    service.UpdatedAt = now
    
    return s.db.QueryRow(query, service.Title, service.Description, service.Price, service.Category, service.IconURL, service.CreatedAt, service.UpdatedAt).Scan(&service.ID)
}

func (s *Store) UpdateService(id int, service *models.Service) error {
    query := `UPDATE services SET title = $2, description = $3, price = $4, category = $5, icon_url = $6, updated_at = $7 WHERE id = $1`
    
    service.UpdatedAt = time.Now()
    _, err := s.db.Exec(query, id, service.Title, service.Description, service.Price, service.Category, service.IconURL, service.UpdatedAt)
    return err
}

func (s *Store) DeleteService(id int) error {
    query := `DELETE FROM services WHERE id = $1`
    _, err := s.db.Exec(query, id)
    return err
}

// Team methods (без изменений)
func (s *Store) GetTeam() ([]models.TeamMember, error) {
    query := `SELECT id, name, position, experience, photo_url, created_at, updated_at FROM team`
    
    rows, err := s.db.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var team []models.TeamMember
    for rows.Next() {
        var member models.TeamMember
        if err := rows.Scan(&member.ID, &member.Name, &member.Position, &member.Experience, &member.PhotoURL, &member.CreatedAt, &member.UpdatedAt); err != nil {
            return nil, err
        }
        team = append(team, member)
    }
    
    return team, nil
}

// Projects methods (без изменений)
func (s *Store) GetProjects() ([]models.Project, error) {
    query := `SELECT id, title, description, category, status, created_at, updated_at FROM projects`
    
    rows, err := s.db.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var projects []models.Project
    for rows.Next() {
        var project models.Project
        if err := rows.Scan(&project.ID, &project.Title, &project.Description, &project.Category, &project.Status, &project.CreatedAt, &project.UpdatedAt); err != nil {
            return nil, err
        }
        projects = append(projects, project)
    }
    
    return projects, nil
}

// Stats methods (без изменений)
func (s *Store) GetStats() (map[string]interface{}, error) {
    stats := make(map[string]interface{})

    // Получаем последние данные по штрафам
    var fineStats models.Fine
    fineQuery := `SELECT violations_total, orders_total, fines_amount_total, collected_amount_total 
                  FROM fines ORDER BY date DESC LIMIT 1`
    err := s.db.QueryRow(fineQuery).Scan(&fineStats.ViolationsTotal, &fineStats.OrdersTotal, &fineStats.FinesAmountTotal, &fineStats.CollectedAmountTotal)
    if err == nil {
        stats["violations_total"] = fineStats.ViolationsTotal
        stats["orders_total"] = fineStats.OrdersTotal
        stats["fines_amount_total"] = fineStats.FinesAmountTotal
        stats["collected_amount_total"] = fineStats.CollectedAmountTotal
    }

    // Получаем последние данные по эвакуации
    var evacStats models.Evacuation
    evacQuery := `SELECT evacuators_count, trips_count, evacuations_count, fine_lot_income 
                  FROM evacuations ORDER BY date DESC LIMIT 1`
    err = s.db.QueryRow(evacQuery).Scan(&evacStats.EvacuatorsCount, &evacStats.TripsCount, &evacStats.EvacuationsCount, &evacStats.FineLotIncome)
    if err == nil {
        stats["evacuators_count"] = evacStats.EvacuatorsCount
        stats["trips_count"] = evacStats.TripsCount
        stats["evacuations_count"] = evacStats.EvacuationsCount
        stats["fine_lot_income"] = evacStats.FineLotIncome
    }

    // Получаем количество светофоров
    var trafficLightsCount int
    lightQuery := `SELECT COUNT(*) FROM traffic_lights WHERE status = 'active'`
    err = s.db.QueryRow(lightQuery).Scan(&trafficLightsCount)
    if err == nil {
        stats["traffic_lights_active"] = trafficLightsCount
    }

    return stats, nil
}

// Traffic methods (без изменений)
func (s *Store) GetTraffic() (map[string]interface{}, error) {
    traffic := make(map[string]interface{})

    // Группируем светофоры по типам
    typeQuery := `SELECT light_type, COUNT(*) FROM traffic_lights GROUP BY light_type`
    rows, err := s.db.Query(typeQuery)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    lightTypes := make(map[string]int)
    for rows.Next() {
        var lightType string
        var count int
        if err := rows.Scan(&lightType, &count); err != nil {
            continue
        }
        lightTypes[lightType] = count
    }
    traffic["light_types"] = lightTypes

    // Группируем по годам установки
    yearQuery := `SELECT install_year, COUNT(*) FROM traffic_lights GROUP BY install_year ORDER BY install_year DESC`
    rows, err = s.db.Query(yearQuery)
    if err != nil {
        return traffic, nil
    }
    defer rows.Close()

    installYears := make(map[int]int)
    for rows.Next() {
        var year, count int
        if err := rows.Scan(&year, &count); err != nil {
            continue
        }
        installYears[year] = count
    }
    traffic["install_years"] = installYears

    return traffic, nil
}
