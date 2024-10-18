# Tracking Service

## Table of Contents

- [Overview](#overview)
- [API Endpoints](#api-endpoints)
- [Dependencies](#dependencies)
- [Setup](#setup)
- [Running the Service](#running-the-service)
- [Testing](#testing)
- [Deployment](#deployment)

## Overview

The Tracking Service is responsible for tracking the real-time location of drivers and bookings. It interacts with the Driver Service and Booking Service to retrieve and update location data for effective tracking and monitoring.

## API Endpoints

- **POST /track**: Update the location of a driver or booking.
- **GET /track/:id**: Retrieve the location information by ID.

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

The service will be available at `http://localhost:8085`.

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

RUN go build -o /tracking-service

EXPOSE 8085

CMD ["/tracking-service"]
```

### Kubernetes Deployment

Create a Kubernetes deployment file:

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: tracking-service
spec:
  replicas: 3
  selector:
    matchLabels:
      app: tracking-service
  template:
    metadata:
      labels:
        app: tracking-service
    spec:
      containers:
        - name: tracking-service
          image: tracking-service:latest
          ports:
            - containerPort: 8085
```

Create a Kubernetes service file:

```yaml
apiVersion: v1
kind: Service
metadata:
  name: tracking-service
spec:
  selector:
    app: tracking-service
  ports:
    - protocol: TCP
      port: 8085
      targetPort: 8085
```

### Deploy to Kubernetes

```bash
kubectl apply -f deployments/tracking-service-deployment.yaml
kubectl apply -f deployments/tracking-service-service.yaml
```
