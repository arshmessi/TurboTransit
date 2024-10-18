CREATE TABLE locations (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    driver_id INTEGER NOT NULL,
    booking_id INTEGER,
    latitude REAL NOT NULL,
    longitude REAL NOT NULL,
    timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (driver_id) REFERENCES drivers(id),
    FOREIGN KEY (booking_id) REFERENCES bookings(id)
);