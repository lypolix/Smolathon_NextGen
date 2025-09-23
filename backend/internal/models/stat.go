package models

import "time"

type Stat struct {
    ID          int       `json:"id" db:"id"`
    Type        string    `json:"type" db:"type"`
    Value       int       `json:"value" db:"value"`
    Title       string    `json:"title" db:"title"`
    Description string    `json:"description" db:"description"`
    Date        time.Time `json:"date" db:"date"`
    CreatedAt   time.Time `json:"created_at" db:"created_at"`
    UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

type CreateStatRequest struct {
    Type        string `json:"type" binding:"required"`
    Value       int    `json:"value" binding:"required"`
    Title       string `json:"title" binding:"required"`
    Description string `json:"description"`
}

type UpdateStatRequest struct {
    Type        string `json:"type"`
    Value       int    `json:"value"`
    Title       string `json:"title"`
    Description string `json:"description"`
}
