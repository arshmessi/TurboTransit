# Booking Service

**Note:** This service is under development. Some features mentioned in this document may not be fully implemented yet, and unit tests are yet to be written.

## Table of Contents

- [Overview](#overview)
- [API Endpoints](#api-endpoints)
- [Dependencies](#dependencies)
- [Setup](#setup)
- [Running the Service](#running-the-service)
- [Testing](#testing)
- [Deployment](#deployment)

## Overview

The Booking Service manages booking operations such as creation, updates, and cancellations. It interacts with a database to store booking information and uses Redis for caching to optimize performance.

## API Endpoints

- **POST /bookings**: Create a new booking.
- **GET /bookings/:id**: Retrieve booking by ID.
- **PUT /bookings/:id**: Update booking by ID.
- **DELETE /bookings/:id**: Cancel booking by ID.

## Dependencies

- Go 1.17+
- SQLite3
- Redis
- NATS
- Gin (for routing)

## Setup

### Install Dependencies

```bash
go mod download
```

### Database Setup

Create a SQLite database:

```bash
sqlite3 bookings.db
```

Create the `bookings` table:

```sql
CREATE TABLE bookings (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER NOT NULL,
    driver_id INTEGER,
    start_location TEXT,
    end_location TEXT,
    status TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

### Redis Setup

Ensure Redis is running on `localhost:6379`.

### NATS Setup

Ensure NATS is running on `localhost:4222`.

## Running the Service

### Start the Service

```bash
go run cmd/main.go
```

The service will be available at `http://localhost:8083`.

## Testing

### Run Unit Tests

```bash
go test ./...
```

### Run Integration Tests

Ensure the service is running, then run:

```bash
go test ./internal/controller
```

## Deployment

### Docker Setup

Create a Dockerfile:

```dockerfile
FROM golang:1.17-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o /booking-service

EXPOSE 8083

CMD ["/booking-service"]
```

### Kubernetes Deployment

Create a Kubernetes deployment file:

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: booking-service
spec:
  replicas: 3
  selector:
    matchLabels:
      app: booking-service
  template:
    metadata:
      labels:
        app: booking-service
    spec:
      containers:
        - name: booking-service
          image: booking-service:latest
          ports:
            - containerPort: 8083
```

Create a Kubernetes service file:

```yaml
apiVersion: v1
kind: Service
metadata:
  name: booking-service
spec:
  selector:
    app: booking-service
  ports:
    - protocol: TCP
      port: 8083
      targetPort: 8083
```

### Deploy to Kubernetes

```bash
kubectl apply -f deployments/booking-service-deployment.yaml
kubectl apply -f deployments/booking-service-service.yaml
```
