package models

import "time"

// Vacancy — сущность вакансии.
type Vacancy struct {
	ID         int        `json:"id" db:"id"`
	Position   string     `json:"position" db:"position"`
	Experience string     `json:"experience" db:"experience"`
	Salary     string     `json:"salary" db:"salary"` // строка, чтобы гибко хранить "от… до…", "по договорённости" и т.п.
	CreatedAt  *time.Time `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt  *time.Time `json:"updated_at,omitempty" db:"updated_at"`
}


type CreateVacancyRequest struct {
	Position   string `json:"position" binding:"required"`
	Experience string `json:"experience" binding:"required"`
	Salary     string `json:"salary" binding:"required"`
}

type UpdateVacancyRequest struct {
	Position   *string `json:"position"`
	Experience *string `json:"experience"`
	Salary     *string `json:"salary"`
}
