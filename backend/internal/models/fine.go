package models

import "time"

type Fine struct {
    ID                     int       `json:"id" db:"id"`
    Date                   time.Time `json:"date" db:"date"`
    ViolationsTotal        int       `json:"violations_total" db:"violations_total"`
    OrdersTotal           int       `json:"orders_total" db:"orders_total"`
    FinesAmountTotal      int       `json:"fines_amount_total" db:"fines_amount_total"`
    CollectedAmountTotal  int       `json:"collected_amount_total" db:"collected_amount_total"`
    CreatedAt            time.Time `json:"created_at" db:"created_at"`
    UpdatedAt            time.Time `json:"updated_at" db:"updated_at"`
}

type CreateFineRequest struct {
    Date                   time.Time `json:"date" binding:"required"`
    ViolationsTotal        int       `json:"violations_total" binding:"required"`
    OrdersTotal           int       `json:"orders_total" binding:"required"`
    FinesAmountTotal      int       `json:"fines_amount_total" binding:"required"`
    CollectedAmountTotal  int       `json:"collected_amount_total" binding:"required"`
}

type UpdateFineRequest struct {
    Date                   time.Time `json:"date"`
    ViolationsTotal        int       `json:"violations_total"`
    OrdersTotal           int       `json:"orders_total"`
    FinesAmountTotal      int       `json:"fines_amount_total"`
    CollectedAmountTotal  int       `json:"collected_amount_total"`
}
