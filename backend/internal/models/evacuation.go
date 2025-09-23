package models

import "time"

type Evacuation struct {
    ID               int       `json:"id" db:"id"`
    Date             time.Time `json:"date" db:"date"`
    EvacuatorsCount  int       `json:"evacuators_count" db:"evacuators_count"`
    TripsCount       int       `json:"trips_count" db:"trips_count"`
    EvacuationsCount int       `json:"evacuations_count" db:"evacuations_count"`
    FineLotIncome    int       `json:"fine_lot_income" db:"fine_lot_income"`
    CreatedAt        time.Time `json:"created_at" db:"created_at"`
    UpdatedAt        time.Time `json:"updated_at" db:"updated_at"`
}

type EvacuationRoute struct {
    ID        int       `json:"id" db:"id"`
    Year      int       `json:"year" db:"year"`
    Month     string    `json:"month" db:"month"`
    Route     string    `json:"route" db:"route"`
    CreatedAt time.Time `json:"created_at" db:"created_at"`
    UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type CreateEvacuationRequest struct {
    Date             time.Time `json:"date" binding:"required"`
    EvacuatorsCount  int       `json:"evacuators_count" binding:"required"`
    TripsCount       int       `json:"trips_count" binding:"required"`
    EvacuationsCount int       `json:"evacuations_count" binding:"required"`
    FineLotIncome    int       `json:"fine_lot_income" binding:"required"`
}

type CreateEvacuationRouteRequest struct {
    Year  int    `json:"year" binding:"required"`
    Month string `json:"month" binding:"required"`
    Route string `json:"route" binding:"required"`
}
