package models

type AdminLoginRequest struct {
    AccessKey string `json:"access_key" binding:"required"`
}

type EditorLoginRequest struct {
    AccessKey string `json:"access_key" binding:"required"`
}

type LoginResponse struct {
    Token string `json:"token"`
    Role  string `json:"role"`
    Valid bool   `json:"valid"`
}

type AccessKey struct {
    ID        int    `json:"id" db:"id"`
    Key       string `json:"key" db:"key"`
    Role      string `json:"role" db:"role"`
    IsActive  bool   `json:"is_active" db:"is_active"`
    CreatedBy string `json:"created_by" db:"created_by"`
    CreatedAt string `json:"created_at" db:"created_at"`
}
