package model

import "time"

type Location struct {
    ID        int       `json:"id"`
    DriverID  int       `json:"driver_id"`
    BookingID int       `json:"booking_id"`
    Latitude  float64   `json:"latitude"`
    Longitude float64   `json:"longitude"`
    Timestamp time.Time `json:"timestamp"`
}