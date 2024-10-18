# Driver Service

**Note:** This service is under active development. Some features in this document are yet to be implemented, and unit tests have not been written.

## Table of Contents

- [Overview](#overview)
- [API Endpoints](#api-endpoints)
- [Dependencies](#dependencies)
- [Setup](#setup)
- [Running the Service](#running-the-service)
- [Testing](#testing)
- [Deployment](#deployment)

## Overview

The Driver Service is responsible for managing driver-related operations such as registration, status updates, and vehicle assignments. It interacts with a database to store driver information and uses Redis for caching to improve performance.

## API Endpoints

- **POST /drivers**: Register a new driver.
- **GET /drivers/:id**: Retrieve driver profile by ID.
- **PUT /drivers/:id/status**: Update driver status by ID.
- **POST /drivers/:id/vehicles**: Assign vehicle to driver by ID.

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
sqlite3 drivers.db
```

Create the `drivers` table:

```sql
CREATE TABLE drivers (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    username TEXT NOT NULL UNIQUE,
    email TEXT NOT NULL UNIQUE,
    password_hash TEXT NOT NULL,
    salt TEXT NOT NULL,
    first_name TEXT,
    last_name TEXT,
    status TEXT,
    vehicle_id INTEGER,
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

The service will be available at `http://localhost:8082`.

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

RUN go build -o /driver-service

EXPOSE 8082

CMD ["/driver-service"]
```

### Kubernetes Deployment

Create a Kubernetes deployment file:

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: driver-service
spec:
  replicas: 3
  selector:
    matchLabels:
      app: driver-service
  template:
    metadata:
      labels:
        app: driver-service
    spec:
      containers:
        - name: driver-service
          image: driver-service:latest
          ports:
            - containerPort: 8082
```

Create a Kubernetes service file:

```yaml
apiVersion: v1
kind: Service
metadata:
  name: driver-service
spec:
  selector:
    app: driver-service
  ports:
    - protocol: TCP
      port: 8082
      targetPort: 8082
```

### Deploy to Kubernetes

```bash
kubectl apply -f deployments/driver-service-deployment.yaml
kubectl apply -f deployments/driver-service-service.yaml
```
