package models

import "time"

type TeamMember struct {
    ID        int       `json:"id" db:"id"`
    Name      string    `json:"name" db:"name"`
    Position  string    `json:"position" db:"position"`
    Experience string   `json:"experience" db:"experience"`
    PhotoURL  string    `json:"photo_url" db:"photo_url"`
    CreatedAt time.Time `json:"created_at" db:"created_at"`
    UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type CreateTeamMemberRequest struct {
    Name       string `json:"name" binding:"required"`
    Position   string `json:"position" binding:"required"`
    Experience string `json:"experience" binding:"required"`
    PhotoURL   string `json:"photo_url"`
}

type UpdateTeamMemberRequest struct {
    Name       string `json:"name"`
    Position   string `json:"position"`
    Experience string `json:"experience"`
    PhotoURL   string `json:"photo_url"`
}
