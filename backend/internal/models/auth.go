package models


type AccessKey struct {
    ID        int    `json:"id" db:"id"`
    Key       string `json:"key" db:"key"`
    Role      string `json:"role" db:"role"`
    IsActive  bool   `json:"is_active" db:"is_active"`
    CreatedBy string `json:"created_by" db:"created_by"`
    CreatedAt string `json:"created_at" db:"created_at"`
}
