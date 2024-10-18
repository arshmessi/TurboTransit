package model

import "time"

type Booking struct {
    ID              int       `json:"id"`
    UserID          int       `json:"user_id"`
    DriverID        int       `json:"driver_id"`
    PickupLocation  string    `json:"pickup_location"`
    DropoffLocation string    `json:"dropoff_location"`
    Status          string    `json:"status"`
    CreatedAt       time.Time `json:"created_at"`
    UpdatedAt       time.Time `json:"updated_at"`
}