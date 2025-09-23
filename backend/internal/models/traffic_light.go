package models

import "time"

type TrafficLight struct {
    ID          int       `json:"id" db:"id"`
    Address     string    `json:"address" db:"address"`
    LightType   string    `json:"light_type" db:"light_type"`
    InstallYear int       `json:"install_year" db:"install_year"`
    Status      string    `json:"status" db:"status"`
    CreatedAt   time.Time `json:"created_at" db:"created_at"`
    UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

type CreateTrafficLightRequest struct {
    Address     string `json:"address" binding:"required"`
    LightType   string `json:"light_type" binding:"required"`
    InstallYear int    `json:"install_year" binding:"required"`
    Status      string `json:"status"`
}

type UpdateTrafficLightRequest struct {
    Address     string `json:"address"`
    LightType   string `json:"light_type"`
    InstallYear int    `json:"install_year"`
    Status      string `json:"status"`
}
