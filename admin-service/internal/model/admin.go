package model

import "time"

type AdminUser struct {
    ID        int       `json:"id"`
    Username  string    `json:"username"`
    Email     string    `json:"email"`
    Password  string    `json:"-"`
    Role      string    `json:"role"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}

type SystemSetting struct {
    Key       string    `json:"key"`
    Value     string    `json:"value"`
    UpdatedAt time.Time `json:"updated_at"`
}