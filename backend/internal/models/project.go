package models

import "time"

type Project struct {
    ID          int       `json:"id" db:"id"`
    Title       string    `json:"title" db:"title"`
    Description string    `json:"description" db:"description"`
    Category    string    `json:"category" db:"category"`
    Status      string    `json:"status" db:"status"`
    CreatedAt   time.Time `json:"created_at" db:"created_at"`
    UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

type CreateProjectRequest struct {
    Title       string `json:"title" binding:"required"`
    Description string `json:"description" binding:"required"`
    Category    string `json:"category" binding:"required"`
    Status      string `json:"status"`
}

type UpdateProjectRequest struct {
    Title       string `json:"title"`
    Description string `json:"description"`
    Category    string `json:"category"`
    Status      string `json:"status"`
}
