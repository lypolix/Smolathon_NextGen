package models

import "time"

type Service struct {
    ID          int       `json:"id" db:"id"`
    Title       string    `json:"title" db:"title"`
    Description string    `json:"description" db:"description"`
    Price       int       `json:"price" db:"price"`
    Category    string    `json:"category" db:"category"`
    IconURL     string    `json:"icon_url" db:"icon_url"`
    CreatedAt   time.Time `json:"created_at" db:"created_at"`
    UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

type CreateServiceRequest struct {
    Title       string `json:"title" binding:"required"`
    Description string `json:"description" binding:"required"`
    Price       int    `json:"price" binding:"required"`
    Category    string `json:"category" binding:"required"`
    IconURL     string `json:"icon_url"`
}

type UpdateServiceRequest struct {
    Title       string `json:"title"`
    Description string `json:"description"`
    Price       int    `json:"price"`
    Category    string `json:"category"`
    IconURL     string `json:"icon_url"`
}
