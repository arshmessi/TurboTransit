package model

import "time"

type AuthToken struct {
    ID        int       `json:"id"`
    UserID    int       `json:"user_id"`
    Token     string    `json:"token"`
    ExpiresAt time.Time `json:"expires_at"`
    CreatedAt time.Time `json:"created_at"`
}