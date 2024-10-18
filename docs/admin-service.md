# Admin Service

## Table of Contents

- [Overview](#overview)
- [API Endpoints](#api-endpoints)
- [Dependencies](#dependencies)
- [Setup](#setup)
- [Running the Service](#running-the-service)
- [Testing](#testing)
- [Deployment](#deployment)

## Overview

The Admin Service provides administrative functionalities such as managing users, drivers, and bookings. It interacts with the User Service, Driver Service, and Booking Service to perform administrative tasks like retrieving and deleting records.

## API Endpoints

- **GET /admin/users**: Retrieve all users.
- **GET /admin/drivers**: Retrieve all drivers.
- **GET /admin/bookings**: Retrieve all bookings.
- **DELETE /admin/users/:id**: Delete user by ID.
- **DELETE /admin/drivers/:id**: Delete driver by ID.
- **DELETE /admin/bookings/:id**: Delete booking by ID.

## Dependencies

- Go 1.17+
- Redis
- NATS
- Gin (for routing)

## Setup

### Install Dependencies

```bash
go mod download
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

The service will be accessible at `http://localhost:8087`.

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

RUN go build -o /admin-service

EXPOSE 8087

CMD ["/admin-service"]
```

### Kubernetes Deployment

Create a Kubernetes deployment file:

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: admin-service
spec:
  replicas: 3
  selector:
    matchLabels:
      app: admin-service
  template:
    metadata:
      labels:
        app: admin-service
    spec:
      containers:
        - name: admin-service
          image: admin-service:latest
          ports:
            - containerPort: 8087
```

Create a Kubernetes service file:

```yaml
apiVersion: v1
kind: Service
metadata:
  name: admin-service
spec:
  selector:
    app: admin-service
  ports:
    - protocol: TCP
      port: 8087
      targetPort: 8087
```

### Deploy to Kubernetes

```bash
kubectl apply -f deployments/admin-service-deployment.yaml
kubectl apply -f deployments/admin-service-service.yaml
```
