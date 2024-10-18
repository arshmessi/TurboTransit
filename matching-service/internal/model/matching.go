package model

import "time"

type MatchingHistory struct {
    ID         int       `json:"id"`
    BookingID  int       `json:"booking_id"`
    DriverID   int       `json:"driver_id"`
    MatchScore float64   `json:"match_score"`
    Accepted   bool      `json:"accepted"`
    CreatedAt  time.Time `json:"created_at"`
}