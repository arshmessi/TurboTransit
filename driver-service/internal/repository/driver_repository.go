package repository

import (
	"TurboTransit/driver-service/internal/model"
	"database/sql"
	"time"
)

type DriverRepository struct {
    db *sql.DB
}

func NewDriverRepository(db *sql.DB) *DriverRepository {
    return &DriverRepository{db: db}
}

func (r *DriverRepository) CreateDriver(driver *model.Driver) error {
    stmt, err := r.db.Prepare("INSERT INTO drivers (username, email, password_hash, salt, first_name, last_name, license_number, status) VALUES (?, ?, ?, ?, ?, ?, ?, ?)")
    if err != nil {
        return err
    }
    defer stmt.Close()

    _, err = stmt.Exec(driver.Username, driver.Email, driver.Password, driver.FirstName, driver.LastName, driver.LicenseNumber, driver.Status)
    if err != nil {
        return err
    }
    return nil
}

func (r *DriverRepository) GetDriverByID(id int) (*model.Driver, error) {
    row := r.db.QueryRow("SELECT id, username, email, first_name, last_name, license_number, status, created_at, updated_at FROM drivers WHERE id = ?", id)
    driver := &model.Driver{}
    err := row.Scan(&driver.ID, &driver.Username, &driver.Email, &driver.FirstName, &driver.LastName, &driver.LicenseNumber, &driver.Status, &driver.CreatedAt, &driver.UpdatedAt)
    if err != nil {
        return nil, err
    }
    return driver, nil
}

func (r *DriverRepository) UpdateDriverStatus(id int, status string) error {
    stmt, err := r.db.Prepare("UPDATE drivers SET status = ?, updated_at = ? WHERE id = ?")
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

func (r *DriverRepository) AssignVehicle(driverID int, vehicle *model.Vehicle) error {
    stmt, err := r.db.Prepare("INSERT INTO vehicles (driver_id, vehicle_type, license_plate, model) VALUES (?, ?, ?, ?)")
    if err != nil {
        return err
    }
    defer stmt.Close()

    _, err = stmt.Exec(driverID, vehicle.VehicleType, vehicle.LicensePlate, vehicle.Model)
    if err != nil {
        return err
    }
    return nil
}