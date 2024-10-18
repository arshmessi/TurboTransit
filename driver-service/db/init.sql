CREATE TABLE drivers (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    username TEXT NOT NULL UNIQUE,
    email TEXT NOT NULL UNIQUE,
    password_hash TEXT NOT NULL,
    salt TEXT NOT NULL,
    first_name TEXT,
    last_name TEXT,
    license_number TEXT NOT NULL UNIQUE,
    status TEXT CHECK(status IN ('available', 'busy', 'offline')) DEFAULT 'offline',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE vehicles (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    driver_id INTEGER,
    vehicle_type TEXT NOT NULL,
    license_plate TEXT NOT NULL UNIQUE,
    model TEXT NOT NULL,
    FOREIGN KEY (driver_id) REFERENCES drivers(id)
);