# Matching Service

## Table of Contents

- [Overview](#overview)
- [API Endpoints](#api-endpoints)
- [Dependencies](#dependencies)
- [Setup](#setup)
- [Running the Service](#running-the-service)
- [Testing](#testing)
- [Deployment](#deployment)

## Overview

The Matching Service is responsible for pairing drivers with bookings based on location and availability. It interacts with the Driver Service and Booking Service to retrieve and update relevant data, ensuring efficient driver-booking matches.

## API Endpoints

- **POST /match**: Match a driver with a booking.
- **GET /match/:id**: Retrieve match information by ID.

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

The service will be available at `http://localhost:8084`.

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

RUN go build -o /matching-service

EXPOSE 8084

CMD ["/matching-service"]
```

### Kubernetes Deployment

Create a Kubernetes deployment file:

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: matching-service
spec:
  replicas: 3
  selector:
    matchLabels:
      app: matching-service
  template:
    metadata:
      labels:
        app: matching-service
    spec:
      containers:
        - name: matching-service
          image: matching-service:latest
          ports:
            - containerPort: 8084
```

Create a Kubernetes service file:

```yaml
apiVersion: v1
kind: Service
metadata:
  name: matching-service
spec:
  selector:
    app: matching-service
  ports:
    - protocol: TCP
      port: 8084
      targetPort: 8084
```

### Deploy to Kubernetes

```bash
kubectl apply -f deployments/matching-service-deployment.yaml
kubectl apply -f deployments/matching-service-service.yaml
```
