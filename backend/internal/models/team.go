package models

import "time"

type TeamMember struct {
    ID         int        `json:"id" db:"id"`
    Name       string     `json:"name" db:"name"`
    Position   string     `json:"position" db:"position"`
    Experience string     `json:"experience" db:"experience"`
    PhotoURL   *string    `json:"photo_url,omitempty" db:"photo_url"`
    CreatedAt  *time.Time `json:"created_at,omitempty" db:"created_at"`
    UpdatedAt  *time.Time `json:"updated_at,omitempty" db:"updated_at"`
}

// Запрос на создание участника.
// Поле PhotoURL опционально.
type CreateTeamMemberRequest struct {
    Name       string  `json:"name" binding:"required"`
    Position   string  `json:"position" binding:"required"`
    Experience string  `json:"experience" binding:"required"`
    PhotoURL   *string `json:"photo_url"`
}

// Запрос на обновление участника.
// Все поля опциональны.
type UpdateTeamMemberRequest struct {
    Name       *string `json:"name"`
    Position   *string `json:"position"`
    Experience *string `json:"experience"`
    PhotoURL   *string `json:"photo_url"`
}
