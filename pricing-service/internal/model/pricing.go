package model

import "time"

type PricingRule struct {
    ID             int       `json:"id"`
    VehicleType    string    `json:"vehicle_type"`
    BaseFare       float64   `json:"base_fare"`
    PerKmRate      float64   `json:"per_km_rate"`
    PerMinuteRate  float64   `json:"per_minute_rate"`
    CreatedAt      time.Time `json:"created_at"`
    UpdatedAt      time.Time `json:"updated_at"`
}

type SurgePricingHistory struct {
    ID         int       `json:"id"`
    Multiplier float64   `json:"multiplier"`
    StartTime  time.Time `json:"start_time"`
    EndTime    time.Time `json:"end_time"`
    Reason     string    `json:"reason"`
}