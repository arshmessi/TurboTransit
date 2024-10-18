package repository

import (
	"TurboTransit/booking-service/internal/model"
	"database/sql"
	"time"
)

type BookingRepository struct {
    db *sql.DB
}

func NewBookingRepository(db *sql.DB) *BookingRepository {
    return &BookingRepository{db: db}
}

func (r *BookingRepository) CreateBooking(booking *model.Booking) error {
    stmt, err := r.db.Prepare("INSERT INTO bookings (user_id, driver_id, pickup_location, dropoff_location, status) VALUES (?, ?, ?, ?, ?)")
    if err != nil {
        return err
    }
    defer stmt.Close()

    _, err = stmt.Exec(booking.UserID, booking.DriverID, booking.PickupLocation, booking.DropoffLocation, booking.Status)
    if err != nil {
        return err
    }
    return nil
}

func (r *BookingRepository) GetBookingByID(id int) (*model.Booking, error) {
    row := r.db.QueryRow("SELECT id, user_id, driver_id, pickup_location, dropoff_location, status, created_at, updated_at FROM bookings WHERE id = ?", id)
    booking := &model.Booking{}
    err := row.Scan(&booking.ID, &booking.UserID, &booking.DriverID, &booking.PickupLocation, &booking.DropoffLocation, &booking.Status, &booking.CreatedAt, &booking.UpdatedAt)
    if err != nil {
        return nil, err
    }
    return booking, nil
}

func (r *BookingRepository) UpdateBookingStatus(id int, status string) error {
    stmt, err := r.db.Prepare("UPDATE bookings SET status = ?, updated_at = ? WHERE id = ?")
    if err != nil {
        return err
    }
    defer stmt.Close()

    _, err = stmt.Exec(status, time.Now(), id)
    if err != nil {
        return err
    }
    return nil
}

func (r *BookingRepository) GetUserBookings(userID int) ([]*model.Booking, error) {
    rows, err := r.db.Query("SELECT id, user_id, driver_id, pickup_location, dropoff_location, status, created_at, updated_at FROM bookings WHERE user_id = ?", userID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var bookings []*model.Booking
    for rows.Next() {
        booking := &model.Booking{}
        err := rows.Scan(&booking.ID, &booking.UserID, &booking.DriverID, &booking.PickupLocation, &booking.DropoffLocation, &booking.Status, &booking.CreatedAt, &booking.UpdatedAt)
        if err != nil {
            return nil, err
        }
        bookings = append(bookings, booking)
    }
    return bookings, nil
}