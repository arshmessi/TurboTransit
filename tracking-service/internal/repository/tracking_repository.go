package repository

import (
	"TurboTransit/tracking-service/internal/model"
	"database/sql"
)

type TrackingRepository struct {
    db *sql.DB
}

func NewTrackingRepository(db *sql.DB) *TrackingRepository {
    return &TrackingRepository{db: db}
}

func (r *TrackingRepository) UpdateLocation(location *model.Location) error {
    stmt, err := r.db.Prepare("INSERT INTO locations (driver_id, booking_id, latitude, longitude) VALUES (?, ?, ?, ?)")
    if err != nil {
        return err
    }
    defer stmt.Close()

    _, err = stmt.Exec(location.DriverID, location.BookingID, location.Latitude, location.Longitude)
    if err != nil {
        return err
    }
    return nil
}

func (r *TrackingRepository) GetDriverLocation(driverID int) (*model.Location, error) {
    row := r.db.QueryRow("SELECT id, driver_id, booking_id, latitude, longitude, timestamp FROM locations WHERE driver_id = ? ORDER BY timestamp DESC LIMIT 1", driverID)
    location := &model.Location{}
    err := row.Scan(&location.ID, &location.DriverID, &location.BookingID, &location.Latitude, &location.Longitude, &location.Timestamp)
    if err != nil {
        return nil, err
    }
    return location, nil
}

func (r *TrackingRepository) GetBookingLocations(bookingID int) ([]*model.Location, error) {
    rows, err := r.db.Query("SELECT id, driver_id, booking_id, latitude, longitude, timestamp FROM locations WHERE booking_id = ? ORDER BY timestamp DESC", bookingID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var locations []*model.Location
    for rows.Next() {
        location := &model.Location{}
        err := rows.Scan(&location.ID, &location.DriverID, &location.BookingID, &location.Latitude, &location.Longitude, &location.Timestamp)
        if err != nil {
            return nil, err
        }
        locations = append(locations, location)
    }
    return locations, nil
}