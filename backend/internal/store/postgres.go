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
    dsn := fmt.Sprintf(
        "host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName,
    )

    db, err := sql.Open("postgres", dsn)
    if err != nil {
        return nil, err
    }
    if err := db.Ping(); err != nil {
        return nil, err
    }

    // Опционально ограничить пул
    db.SetMaxOpenConns(10)
    db.SetMaxIdleConns(10)
    db.SetConnMaxLifetime(30 * time.Minute)

    log.Println("Connected to PostgreSQL")
    return &Store{db: db}, nil
}

func (s *Store) Close() error  { return s.db.Close() }
func (s *Store) GetDB() *sql.DB { return s.db }

// Users
func (s *Store) GetUserByEmail(email string) (*models.User, error) {
    user := &models.User{}

    query := `
        SELECT id, email, password, role, created_at, updated_at 
        FROM users 
        WHERE email = $1
    `
    log.Printf("GetUserByEmail: searching for email '%s'", email)

    err := s.db.QueryRow(query, email).Scan(
        &user.ID,
        &user.Email,
        &user.Password,
        &user.Role,
        &user.CreatedAt,
        &user.UpdatedAt,
    )

    if err != nil {
        if err == sql.ErrNoRows {
            log.Printf("GetUserByEmail: user '%s' not found", email)
            return nil, err
        }
        log.Printf("GetUserByEmail error: %v", err)
        return nil, err
    }

    user.IsActive = true

    log.Printf("GetUserByEmail: found user ID=%d, email=%s, role=%s",
        user.ID, user.Email, user.Role)

    return user, nil
}

func (s *Store) CreateUser(user *models.User) error {
    query := `
        INSERT INTO users (email, password, role, is_active) 
        VALUES ($1, $2, $3, $4) 
        RETURNING id, created_at, updated_at
    `
    return s.db.QueryRow(
        query,
        user.Email,
        user.Password,
        user.Role,
        user.IsActive,
    ).Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)
}

func (s *Store) UpdateUserPassword(userID int, hashedPassword string) error {
    query := `
        UPDATE users 
        SET password = $1, updated_at = CURRENT_TIMESTAMP 
        WHERE id = $2
    `
    _, err := s.db.Exec(query, hashedPassword, userID)
    return err
}

// Fines (оставляем time.Time)
func (s *Store) GetFines() ([]models.Fine, error) {
    query := `
        SELECT id, date, violations_total, orders_total, fines_amount_total, collected_amount_total,
               COALESCE(created_at, CURRENT_TIMESTAMP) AS created_at,
               COALESCE(updated_at, CURRENT_TIMESTAMP) AS updated_at
        FROM public.fines
        ORDER BY date DESC
    `
    rows, err := s.db.Query(query)
    if err != nil { log.Printf("GetFines query err: %v", err); return nil, err }
    defer rows.Close()

    var out []models.Fine
    for rows.Next() {
        var f models.Fine
        if err := rows.Scan(
            &f.ID, &f.Date, &f.ViolationsTotal, &f.OrdersTotal, &f.FinesAmountTotal, &f.CollectedAmountTotal,
            &f.CreatedAt, &f.UpdatedAt,
        ); err != nil {
            log.Printf("GetFines scan err: %v", err)
            return nil, err
        }
        out = append(out, f)
    }
    return out, nil
}

func (s *Store) CreateFine(f *models.Fine) error {
    query := `
        INSERT INTO public.fines (date, violations_total, orders_total, fines_amount_total, collected_amount_total, created_at, updated_at)
        VALUES ($1,$2,$3,$4,$5,$6,$7)
        RETURNING id
    `
    now := time.Now()
    f.CreatedAt = now
    f.UpdatedAt = now
    if err := s.db.QueryRow(query, f.Date, f.ViolationsTotal, f.OrdersTotal, f.FinesAmountTotal, f.CollectedAmountTotal, f.CreatedAt, f.UpdatedAt).Scan(&f.ID); err != nil {
        log.Printf("CreateFine err: %v", err)
        return err
    }
    return nil
}

func (s *Store) UpdateFine(id int, f *models.Fine) error {
    query := `
        UPDATE public.fines
        SET date=$2, violations_total=$3, orders_total=$4, fines_amount_total=$5, collected_amount_total=$6, updated_at=$7
        WHERE id=$1
    `
    f.UpdatedAt = time.Now()
    if _, err := s.db.Exec(query, id, f.Date, f.ViolationsTotal, f.OrdersTotal, f.FinesAmountTotal, f.CollectedAmountTotal, f.UpdatedAt); err != nil {
        log.Printf("UpdateFine err: %v", err)
        return err
    }
    return nil
}

func (s *Store) DeleteFine(id int) error {
    if _, err := s.db.Exec(`DELETE FROM public.fines WHERE id=$1`, id); err != nil {
        log.Printf("DeleteFine err: %v", err)
        return err
    }
    return nil
}

// Evacuations (оставляем time.Time)
func (s *Store) GetEvacuations() ([]models.Evacuation, error) {
    query := `
        SELECT id, date, evacuators_count, trips_count, evacuations_count, fine_lot_income,
               COALESCE(created_at, CURRENT_TIMESTAMP) AS created_at,
               COALESCE(updated_at, CURRENT_TIMESTAMP) AS updated_at
        FROM public.evacuations
        ORDER BY date DESC
    `
    rows, err := s.db.Query(query)
    if err != nil { log.Printf("GetEvacuations query err: %v", err); return nil, err }
    defer rows.Close()

    var out []models.Evacuation
    for rows.Next() {
        var e models.Evacuation
        if err := rows.Scan(
            &e.ID, &e.Date, &e.EvacuatorsCount, &e.TripsCount, &e.EvacuationsCount, &e.FineLotIncome, &e.CreatedAt, &e.UpdatedAt,
        ); err != nil {
            log.Printf("GetEvacuations scan err: %v", err)
            return nil, err
        }
        out = append(out, e)
    }
    return out, nil
}

func (s *Store) CreateEvacuation(e *models.Evacuation) error {
    query := `
        INSERT INTO public.evacuations (date, evacuators_count, trips_count, evacuations_count, fine_lot_income, created_at, updated_at)
        VALUES ($1,$2,$3,$4,$5,$6,$7)
        RETURNING id
    `
    now := time.Now()
    e.CreatedAt = now
    e.UpdatedAt = now
    if err := s.db.QueryRow(query, e.Date, e.EvacuatorsCount, e.TripsCount, e.EvacuationsCount, e.FineLotIncome, e.CreatedAt, e.UpdatedAt).Scan(&e.ID); err != nil {
        log.Printf("CreateEvacuation err: %v", err)
        return err
    }
    return nil
}

// Evacuation routes (оставляем time.Time)
func (s *Store) GetEvacuationRoutes() ([]models.EvacuationRoute, error) {
    query := `
        SELECT id, year, month, route,
               COALESCE(created_at, CURRENT_TIMESTAMP) AS created_at,
               COALESCE(updated_at, CURRENT_TIMESTAMP) AS updated_at
        FROM public.evacuation_routes
        ORDER BY year DESC, month
    `
    rows, err := s.db.Query(query)
    if err != nil { log.Printf("GetEvacuationRoutes query err: %v", err); return nil, err }
    defer rows.Close()

    var out []models.EvacuationRoute
    for rows.Next() {
        var r models.EvacuationRoute
        if err := rows.Scan(&r.ID, &r.Year, &r.Month, &r.Route, &r.CreatedAt, &r.UpdatedAt); err != nil {
            log.Printf("GetEvacuationRoutes scan err: %v", err)
            return nil, err
        }
        out = append(out, r)
    }
    return out, nil
}

func (s *Store) CreateEvacuationRoute(r *models.EvacuationRoute) error {
    query := `
        INSERT INTO public.evacuation_routes (year, month, route, created_at, updated_at)
        VALUES ($1,$2,$3,$4,$5)
        RETURNING id
    `
    now := time.Now()
    r.CreatedAt = now
    r.UpdatedAt = now
    if err := s.db.QueryRow(query, r.Year, r.Month, r.Route, r.CreatedAt, r.UpdatedAt).Scan(&r.ID); err != nil {
        log.Printf("CreateEvacuationRoute err: %v", err)
        return err
    }
    return nil
}

// Traffic lights (оставляем time.Time)
func (s *Store) GetTrafficLights() ([]models.TrafficLight, error) {
    query := `
        SELECT id, address, light_type, install_year, status,
               COALESCE(created_at, CURRENT_TIMESTAMP) AS created_at,
               COALESCE(updated_at, CURRENT_TIMESTAMP) AS updated_at
        FROM public.traffic_lights
        ORDER BY install_year DESC
    `
    rows, err := s.db.Query(query)
    if err != nil { log.Printf("GetTrafficLights query err: %v", err); return nil, err }
    defer rows.Close()

    var out []models.TrafficLight
    for rows.Next() {
        var t models.TrafficLight
        if err := rows.Scan(&t.ID, &t.Address, &t.LightType, &t.InstallYear, &t.Status, &t.CreatedAt, &t.UpdatedAt); err != nil {
            log.Printf("GetTrafficLights scan err: %v", err)
            return nil, err
        }
        out = append(out, t)
    }
    return out, nil
}

func (s *Store) CreateTrafficLight(t *models.TrafficLight) error {
    query := `
        INSERT INTO public.traffic_lights (address, light_type, install_year, status, created_at, updated_at)
        VALUES ($1,$2,$3,$4,$5,$6)
        RETURNING id
    `
    now := time.Now()
    t.CreatedAt = now
    t.UpdatedAt = now
    if err := s.db.QueryRow(query, t.Address, t.LightType, t.InstallYear, t.Status, t.CreatedAt, t.UpdatedAt).Scan(&t.ID); err != nil {
        log.Printf("CreateTrafficLight err: %v", err)
        return err
    }
    return nil
}

func (s *Store) UpdateTrafficLight(id int, t *models.TrafficLight) error {
    query := `
        UPDATE public.traffic_lights
        SET address=$2, light_type=$3, install_year=$4, status=$5, updated_at=$6
        WHERE id=$1
    `
    t.UpdatedAt = time.Now()
    if _, err := s.db.Exec(query, id, t.Address, t.LightType, t.InstallYear, t.Status, t.UpdatedAt); err != nil {
        log.Printf("UpdateTrafficLight err: %v", err)
        return err
    }
    return nil
}

func (s *Store) DeleteTrafficLight(id int) error {
    if _, err := s.db.Exec(`DELETE FROM public.traffic_lights WHERE id=$1`, id); err != nil {
        log.Printf("DeleteTrafficLight err: %v", err)
        return err
    }
    return nil
}

// News (оставляем time.Time)
func (s *Store) GetNews() ([]models.News, error) {
    query := `
        SELECT id, title, content, tag, date,
               COALESCE(created_at, CURRENT_TIMESTAMP) AS created_at,
               COALESCE(updated_at, CURRENT_TIMESTAMP) AS updated_at
        FROM public.news
        ORDER BY date DESC
    `
    rows, err := s.db.Query(query)
    if err != nil { log.Printf("GetNews query err: %v", err); return nil, err }
    defer rows.Close()

    var out []models.News
    for rows.Next() {
        var n models.News
        if err := rows.Scan(&n.ID, &n.Title, &n.Content, &n.Tag, &n.Date, &n.CreatedAt, &n.UpdatedAt); err != nil {
            log.Printf("GetNews scan err: %v", err)
            return nil, err
        }
        out = append(out, n)
    }
    return out, nil
}

func (s *Store) GetNewsByID(id int) (*models.News, error) {
    query := `
        SELECT id, title, content, tag, date,
               COALESCE(created_at, CURRENT_TIMESTAMP) AS created_at,
               COALESCE(updated_at, CURRENT_TIMESTAMP) AS updated_at
        FROM public.news
        WHERE id = $1
    `
    var n models.News
    err := s.db.QueryRow(query, id).Scan(
        &n.ID, &n.Title, &n.Content, &n.Tag, &n.Date, &n.CreatedAt, &n.UpdatedAt,
    )
    if err != nil {
        if err == sql.ErrNoRows {
            log.Printf("GetNewsByID: news with id=%d not found", id)
        } else {
            log.Printf("GetNewsByID err: %v", err)
        }
        return nil, err
    }
    return &n, nil
}

func (s *Store) CreateNews(n *models.News) error {
    query := `
        INSERT INTO public.news (title, content, tag, date, created_at, updated_at)
        VALUES ($1,$2,$3,$4,$5,$6)
        RETURNING id
    `
    now := time.Now()
    n.Date = now
    n.CreatedAt = now
    n.UpdatedAt = now
    if err := s.db.QueryRow(query, n.Title, n.Content, n.Tag, n.Date, n.CreatedAt, n.UpdatedAt).Scan(&n.ID); err != nil {
        log.Printf("CreateNews err: %v", err)
        return err
    }
    return nil
}

func (s *Store) UpdateNews(id int, n *models.News) error {
    query := `
        UPDATE public.news
        SET title=$2, content=$3, tag=$4, updated_at=$5
        WHERE id=$1
    `
    n.UpdatedAt = time.Now()
    if _, err := s.db.Exec(query, id, n.Title, n.Content, n.Tag, n.UpdatedAt); err != nil {
        log.Printf("UpdateNews err: %v", err)
        return err
    }
    return nil
}

func (s *Store) DeleteNews(id int) error {
    if _, err := s.db.Exec(`DELETE FROM public.news WHERE id=$1`, id); err != nil {
        log.Printf("DeleteNews err: %v", err)
        return err
    }
    return nil
}

// Services (оставляем time.Time)
func (s *Store) GetServices() ([]models.Service, error) {
    query := `
        SELECT id, title, description, price, category,
               COALESCE(icon_url, '') AS icon_url,
               COALESCE(created_at, CURRENT_TIMESTAMP) AS created_at,
               COALESCE(updated_at, CURRENT_TIMESTAMP) AS updated_at
        FROM public.services
    `
    rows, err := s.db.Query(query)
    if err != nil { log.Printf("GetServices query err: %v", err); return nil, err }
    defer rows.Close()

    var out []models.Service
    for rows.Next() {
        var srv models.Service
        if err := rows.Scan(&srv.ID, &srv.Title, &srv.Description, &srv.Price, &srv.Category, &srv.IconURL, &srv.CreatedAt, &srv.UpdatedAt); err != nil {
            log.Printf("GetServices scan err: %v", err)
            return nil, err
        }
        out = append(out, srv)
    }
    return out, nil
}

func (s *Store) GetServiceByID(id int) (*models.Service, error) {
    query := `
        SELECT id, title, description, price, category,
               COALESCE(icon_url, '') AS icon_url,
               COALESCE(created_at, CURRENT_TIMESTAMP) AS created_at,
               COALESCE(updated_at, CURRENT_TIMESTAMP) AS updated_at
        FROM public.services
        WHERE id = $1
    `
    var srv models.Service
    err := s.db.QueryRow(query, id).Scan(
        &srv.ID, &srv.Title, &srv.Description, &srv.Price, &srv.Category, &srv.IconURL, &srv.CreatedAt, &srv.UpdatedAt,
    )
    if err != nil {
        if err == sql.ErrNoRows {
            log.Printf("GetServiceByID: service with id=%d not found", id)
        } else {
            log.Printf("GetServiceByID err: %v", err)
        }
        return nil, err
    }
    return &srv, nil
}

func (s *Store) CreateService(srv *models.Service) error {
    query := `
        INSERT INTO public.services (title, description, price, category, icon_url, created_at, updated_at)
        VALUES ($1,$2,$3,$4,$5,$6,$7)
        RETURNING id
    `
    now := time.Now()
    srv.CreatedAt = now
    srv.UpdatedAt = now
    if err := s.db.QueryRow(query, srv.Title, srv.Description, srv.Price, srv.Category, srv.IconURL, srv.CreatedAt, srv.UpdatedAt).Scan(&srv.ID); err != nil {
        log.Printf("CreateService err: %v", err)
        return err
    }
    return nil
}

func (s *Store) UpdateService(id int, srv *models.Service) error {
    query := `
        UPDATE public.services
        SET title=$2, description=$3, price=$4, category=$5, icon_url=$6, updated_at=$7
        WHERE id=$1
    `
    srv.UpdatedAt = time.Now()
    if _, err := s.db.Exec(query, id, srv.Title, srv.Description, srv.Price, srv.Category, srv.IconURL, srv.UpdatedAt); err != nil {
        log.Printf("UpdateService err: %v", err)
        return err
    }
    return nil
}

func (s *Store) DeleteService(id int) error {
    if _, err := s.db.Exec(`DELETE FROM public.services WHERE id=$1`, id); err != nil {
        log.Printf("DeleteService err: %v", err)
        return err
    }
    return nil
}

// Team (реализуем создание и обновление; предполагаем timestamps как *time.Time)
func (s *Store) GetTeam() ([]models.TeamMember, error) {
    query := `
        SELECT id, name, position, experience,
               photo_url,
               created_at,
               updated_at
        FROM public.team
        ORDER BY id
    `
    rows, err := s.db.Query(query)
    if err != nil { log.Printf("GetTeam query err: %v", err); return nil, err }
    defer rows.Close()

    var out []models.TeamMember
    for rows.Next() {
        var m models.TeamMember
        if err := rows.Scan(&m.ID, &m.Name, &m.Position, &m.Experience, &m.PhotoURL, &m.CreatedAt, &m.UpdatedAt); err != nil {
            log.Printf("GetTeam scan err: %v", err)
            return nil, err
        }
        out = append(out, m)
    }
    return out, nil
}

func (s *Store) GetTeamMemberByID(id int) (*models.TeamMember, error) {
    query := `
        SELECT id, name, position, experience,
               photo_url,
               created_at,
               updated_at
        FROM public.team
        WHERE id = $1
    `
    var m models.TeamMember
    err := s.db.QueryRow(query, id).Scan(
        &m.ID, &m.Name, &m.Position, &m.Experience, &m.PhotoURL, &m.CreatedAt, &m.UpdatedAt,
    )
    if err != nil {
        if err == sql.ErrNoRows {
            log.Printf("GetTeamMemberByID: team member with id=%d not found", id)
        } else {
            log.Printf("GetTeamMemberByID err: %v", err)
        }
        return nil, err
    }
    return &m, nil
}

func (s *Store) CreateTeamMember(m *models.TeamMember) error {
    query := `
        INSERT INTO public.team (name, position, experience, photo_url, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5, $6)
        RETURNING id
    `
    now := time.Now()
    // если в модели *time.Time:
    m.CreatedAt = &now
    m.UpdatedAt = &now

    if err := s.db.QueryRow(query, m.Name, m.Position, m.Experience, m.PhotoURL, m.CreatedAt, m.UpdatedAt).Scan(&m.ID); err != nil {
        log.Printf("CreateTeamMember err: %v", err)
        return err
    }
    return nil
}

func (s *Store) UpdateTeam(id int, m *models.TeamMember) (int64, error) {
    query := `
        UPDATE public.team
        SET name=$2, position=$3, experience=$4, photo_url=$5, updated_at=$6
        WHERE id=$1
    `
    now := time.Now()
    // если в модели *time.Time:
    m.UpdatedAt = &now

    res, err := s.db.Exec(query, id, m.Name, m.Position, m.Experience, m.PhotoURL, m.UpdatedAt)
    if err != nil {
        log.Printf("UpdateTeam err: %v", err)
        return 0, err
    }
    n, _ := res.RowsAffected()
    if n == 0 {
        log.Printf("UpdateTeam: team member id=%d not found", id)
    }
    return n, nil
}

func (s *Store) DeleteTeam(id int) error {
    if _, err := s.db.Exec(`DELETE FROM public.team WHERE id=$1`, id); err != nil {
        log.Printf("DeleteTeam err: %v", err)
        return err
    }
    return nil
}

func (s *Store) DeleteTeamByID(id int) (int64, error) {
    res, err := s.db.Exec(`DELETE FROM public.team WHERE id=$1`, id)
    if err != nil {
        log.Printf("DeleteTeamByID err: %v", err)
        return 0, err
    }
    n, _ := res.RowsAffected()
    if n == 0 {
        log.Printf("DeleteTeamByID: team member id=%d not found", id)
    }
    return n, nil
}

// Projects
func (s *Store) GetProjects() ([]models.Project, error) {
    query := `
        SELECT id, title, description, category, status,
               COALESCE(created_at, CURRENT_TIMESTAMP) AS created_at,
               COALESCE(updated_at, CURRENT_TIMESTAMP) AS updated_at
        FROM public.projects
    `
    rows, err := s.db.Query(query)
    if err != nil { log.Printf("GetProjects query err: %v", err); return nil, err }
    defer rows.Close()

    var out []models.Project
    for rows.Next() {
        var p models.Project
        if err := rows.Scan(&p.ID, &p.Title, &p.Description, &p.Category, &p.Status, &p.CreatedAt, &p.UpdatedAt); err != nil {
            log.Printf("GetProjects scan err: %v", err)
            return nil, err
        }
        out = append(out, p)
    }
    return out, nil
}

// CreateProject
func (s *Store) CreateProject(p *models.Project) error {
    query := `
        INSERT INTO public.projects (title, description, category, status, created_at, updated_at)
        VALUES ($1,$2,$3,$4,$5,$6)
        RETURNING id
    `
    now := time.Now()
    p.CreatedAt = now
    p.UpdatedAt = now
    if err := s.db.QueryRow(query, p.Title, p.Description, p.Category, p.Status, p.CreatedAt, p.UpdatedAt).Scan(&p.ID); err != nil {
        log.Printf("CreateProject err: %v", err)
        return err
    }
    return nil
}

// UpdateProject
func (s *Store) UpdateProject(id int, p *models.Project) error {
    query := `
        UPDATE public.projects
        SET title=$2, description=$3, category=$4, status=$5, updated_at=$6
        WHERE id=$1
    `
    p.UpdatedAt = time.Now()
    if _, err := s.db.Exec(query, id, p.Title, p.Description, p.Category, p.Status, p.UpdatedAt); err != nil {
        log.Printf("UpdateProject err: %v", err)
        return err
    }
    return nil
}

// DeleteProject
func (s *Store) DeleteProject(id int) error {
    if _, err := s.db.Exec(`DELETE FROM public.projects WHERE id=$1`, id); err != nil {
        log.Printf("DeleteProject err: %v", err)
        return err
    }
    return nil
}

// Stats
func (s *Store) GetStats() (map[string]interface{}, error) {
    stats := make(map[string]interface{})

    // fines last
    var f models.Fine
    fq := `
        SELECT violations_total, orders_total, fines_amount_total, collected_amount_total
        FROM public.fines ORDER BY date DESC LIMIT 1
    `
    if err := s.db.QueryRow(fq).Scan(&f.ViolationsTotal, &f.OrdersTotal, &f.FinesAmountTotal, &f.CollectedAmountTotal); err == nil {
        stats["violations_total"] = f.ViolationsTotal
        stats["orders_total"] = f.OrdersTotal
        stats["fines_amount_total"] = f.FinesAmountTotal
        stats["collected_amount_total"] = f.CollectedAmountTotal
    }

    // evacuations last
    var e models.Evacuation
    eq := `
        SELECT evacuators_count, trips_count, evacuations_count, fine_lot_income
        FROM public.evacuations ORDER BY date DESC LIMIT 1
    `
    if err := s.db.QueryRow(eq).Scan(&e.EvacuatorsCount, &e.TripsCount, &e.EvacuationsCount, &e.FineLotIncome); err == nil {
        stats["evacuators_count"] = e.EvacuatorsCount
        stats["trips_count"] = e.TripsCount
        stats["evacuations_count"] = e.EvacuationsCount
        stats["fine_lot_income"] = e.FineLotIncome
    }

    // traffic lights count
    var tlActive int
    if err := s.db.QueryRow(`SELECT COUNT(*) FROM public.traffic_lights WHERE status='active'`).Scan(&tlActive); err == nil {
        stats["traffic_lights_active"] = tlActive
    }

    return stats, nil
}

// Traffic
func (s *Store) GetTraffic() (map[string]interface{}, error) {
    res := make(map[string]interface{})

    // by type
    typeQuery := `SELECT light_type, COUNT(*) FROM public.traffic_lights GROUP BY light_type`
    rows, err := s.db.Query(typeQuery)
    if err != nil { log.Printf("GetTraffic types err: %v", err); return nil, err }
    defer rows.Close()

    byType := make(map[string]int)
    for rows.Next() {
        var lt string
        var c int
        if err := rows.Scan(&lt, &c); err != nil {
            log.Printf("GetTraffic types scan err: %v", err)
            continue
        }
        byType[lt] = c
    }
    res["light_types"] = byType

    // by year
    yearQuery := `SELECT install_year, COUNT(*) FROM public.traffic_lights GROUP BY install_year ORDER BY install_year DESC`
    rows, err = s.db.Query(yearQuery)
    if err != nil {
        return res, nil
    }
    defer rows.Close()

    byYear := make(map[int]int)
    for rows.Next() {
        var y, c int
        if err := rows.Scan(&y, &c); err != nil {
            log.Printf("GetTraffic years scan err: %v", err)
            continue
        }
        byYear[y] = c
    }
    res["install_years"] = byYear

    return res, nil
}

// Vacancies
func (s *Store) GetVacancies() ([]models.Vacancy, error) {
    query := `
        SELECT id,
               position,
               experience,
               salary,
               COALESCE(created_at, CURRENT_TIMESTAMP) AS created_at,
               COALESCE(updated_at, CURRENT_TIMESTAMP) AS updated_at
        FROM public.vacancies
        ORDER BY id DESC
    `
    rows, err := s.db.Query(query)
    if err != nil {
        log.Printf("GetVacancies query err: %v", err)
        return nil, err
    }
    defer rows.Close()

    var out []models.Vacancy
    for rows.Next() {
        var v models.Vacancy
        if err := rows.Scan(&v.ID, &v.Position, &v.Experience, &v.Salary, &v.CreatedAt, &v.UpdatedAt); err != nil {
            log.Printf("GetVacancies scan err: %v", err)
            return nil, err
        }
        out = append(out, v)
    }
    return out, nil
}

func (s *Store) GetVacancyByID(id int) (*models.Vacancy, error) {
    query := `
        SELECT id,
               position,
               experience,
               salary,
               COALESCE(created_at, CURRENT_TIMESTAMP) AS created_at,
               COALESCE(updated_at, CURRENT_TIMESTAMP) AS updated_at
        FROM public.vacancies
        WHERE id = $1
    `
    var v models.Vacancy
    if err := s.db.QueryRow(query, id).Scan(&v.ID, &v.Position, &v.Experience, &v.Salary, &v.CreatedAt, &v.UpdatedAt); err != nil {
        if err == sql.ErrNoRows {
            log.Printf("GetVacancyByID: vacancy id=%d not found", id)
        } else {
            log.Printf("GetVacancyByID err: %v", err)
        }
        return nil, err
    }
    return &v, nil
}

func (s *Store) CreateVacancy(v *models.Vacancy) error {
    query := `
        INSERT INTO public.vacancies (position, experience, salary, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5)
        RETURNING id
    `
    now := time.Now()
    // В модели Vacancy timestamps как *time.Time? В предыдущем файле модели — да.
    // Если у тебя *time.Time — устанавливаем адреса значений:
    v.CreatedAt = &now
    v.UpdatedAt = &now

    if err := s.db.QueryRow(query, v.Position, v.Experience, v.Salary, v.CreatedAt, v.UpdatedAt).Scan(&v.ID); err != nil {
        log.Printf("CreateVacancy err: %v", err)
        return err
    }
    return nil
}

func (s *Store) UpdateVacancy(id int, v *models.Vacancy) error {
    query := `
        UPDATE public.vacancies
        SET position = $2,
            experience = $3,
            salary = $4,
            updated_at = $5
        WHERE id = $1
    `
    now := time.Now()
    v.UpdatedAt = &now

    if _, err := s.db.Exec(query, id, v.Position, v.Experience, v.Salary, v.UpdatedAt); err != nil {
        log.Printf("UpdateVacancy err: %v", err)
        return err
    }
    return nil
}

func (s *Store) DeleteVacancy(id int) error {
    if _, err := s.db.Exec(`DELETE FROM public.vacancies WHERE id = $1`, id); err != nil {
        log.Printf("DeleteVacancy err: %v", err)
        return err
    }
    return nil
}
