package model

import "time"

type Driver struct {
    ID            int       `json:"id"`
    Username      string    `json:"username"`
    Email         string    `json:"email"`
    Password      string    `json:"-"`
    FirstName     string    `json:"first_name"`
    LastName      string    `json:"last_name"`
    LicenseNumber string    `json:"license_number"`
    Status        string    `json:"status"`
    CreatedAt     time.Time `json:"created_at"`
    UpdatedAt     time.Time `json:"updated_at"`
}

type Vehicle struct {
    ID            int    `json:"id"`
    DriverID      int    `json:"driver_id"`
    VehicleType   string `json:"vehicle_type"`
    LicensePlate  string `json:"license_plate"`
    Model         string `json:"model"`
}