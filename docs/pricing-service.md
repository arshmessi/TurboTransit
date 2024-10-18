# Pricing Service

## Table of Contents

- [Overview](#overview)
- [API Endpoints](#api-endpoints)
- [Dependencies](#dependencies)
- [Setup](#setup)
- [Running the Service](#running-the-service)
- [Testing](#testing)
- [Deployment](#deployment)

## Overview

The Pricing Service is responsible for calculating pricing for bookings based on factors such as distance, time, and demand. It interacts with the Booking Service to retrieve booking details and compute the final price.

## API Endpoints

- **POST /price**: Calculate the pricing for a booking.
- **GET /price/:id**: Retrieve the pricing details by booking ID.

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

The service will be accessible at `http://localhost:8086`.

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

RUN go build -o /pricing-service

EXPOSE 8086

CMD ["/pricing-service"]
```

### Kubernetes Deployment

Create a Kubernetes deployment file:

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: pricing-service
spec:
  replicas: 3
  selector:
    matchLabels:
      app: pricing-service
  template:
    metadata:
      labels:
        app: pricing-service
    spec:
      containers:
        - name: pricing-service
          image: pricing-service:latest
          ports:
            - containerPort: 8086
```

Create a Kubernetes service file:

```yaml
apiVersion: v1
kind: Service
metadata:
  name: pricing-service
spec:
  selector:
    app: pricing-service
  ports:
    - protocol: TCP
      port: 8086
      targetPort: 8086
```

### Deploy to Kubernetes

```bash
kubectl apply -f deployments/pricing-service-deployment.yaml
kubectl apply -f deployments/pricing-service-service.yaml
```
