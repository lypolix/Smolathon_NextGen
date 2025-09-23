package models

import "time"

type News struct {
    ID        int       `json:"id" db:"id"`
    Title     string    `json:"title" db:"title"`
    Content   string    `json:"content" db:"content"`
    Tag       string    `json:"tag" db:"tag"`
    Date      time.Time `json:"date" db:"date"`
    CreatedAt time.Time `json:"created_at" db:"created_at"`
    UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type CreateNewsRequest struct {
    Title   string `json:"title" binding:"required"`
    Content string `json:"content" binding:"required"`
    Tag     string `json:"tag" binding:"required"`
}

type UpdateNewsRequest struct {
    Title   string `json:"title"`
    Content string `json:"content"`
    Tag     string `json:"tag"`
}
