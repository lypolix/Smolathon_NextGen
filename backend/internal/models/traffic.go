package models

import "time"

type TrafficReport struct {
    ID          int       `json:"id" db:"id"`
    Type        string    `json:"type" db:"type"`
    Count       int       `json:"count" db:"count"`
    Location    string    `json:"location" db:"location"`
    Status      string    `json:"status" db:"status"`
    Description string    `json:"description" db:"description"`
    CreatedAt   time.Time `json:"created_at" db:"created_at"`
    UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

type CreateTrafficReportRequest struct {
    Type        string `json:"type" binding:"required"`
    Count       int    `json:"count" binding:"required"`
    Location    string `json:"location" binding:"required"`
    Status      string `json:"status"`
    Description string `json:"description"`
}

type UpdateTrafficReportRequest struct {
    Type        string `json:"type"`
    Count       int    `json:"count"`
    Location    string `json:"location"`
    Status      string `json:"status"`
    Description string `json:"description"`
}
