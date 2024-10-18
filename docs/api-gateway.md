# API Gateway

## Table of Contents

- [Overview](#overview)
- [API Endpoints](#api-endpoints)
- [Dependencies](#dependencies)
- [Setup](#setup)
- [Running the Service](#running-the-service)
- [Testing](#testing)
- [Deployment](#deployment)

## Overview

The API Gateway is responsible for routing incoming requests to the appropriate microservices. It manages authentication, routing, and rate limiting, ensuring secure and efficient API usage across all services.

## API Endpoints

- **POST /users**: Route to User Service.
- **GET /users/:id**: Route to User Service.
- **PUT /users/:id**: Route to User Service.

- **POST /auth/login**: Route to Authentication Service.
- **POST /auth/logout**: Route to Authentication Service.
- **POST /auth/refresh**: Route to Authentication Service.

- **POST /drivers**: Route to Driver Service.
- **GET /drivers/:id**: Route to Driver Service.
- **PUT /drivers/:id/status**: Route to Driver Service.
- **POST /drivers/:id/vehicles**: Route to Driver Service.

- **POST /bookings**: Route to Booking Service.
- **GET /bookings/:id**: Route to Booking Service.
- **PUT /bookings/:id**: Route to Booking Service.
- **DELETE /bookings/:id**: Route to Booking Service.

- **POST /match**: Route to Matching Service.
- **GET /match/:id**: Route to Matching Service.

- **POST /track**: Route to Tracking Service.
- **GET /track/:id**: Route to Tracking Service.

- **POST /price**: Route to Pricing Service.
- **GET /price/:id**: Route to Pricing Service.

- **GET /admin/users**: Route to Admin Service.
- **GET /admin/drivers**: Route to Admin Service.
- **GET /admin/bookings**: Route to Admin Service.
- **DELETE /admin/users/:id**: Route to Admin Service.
- **DELETE /admin/drivers/:id**: Route to Admin Service.
- **DELETE /admin/bookings/:id**: Route to Admin Service.

## Dependencies

- Go 1.17+
- Gin (for routing)

## Setup

### Install Dependencies

```bash
go mod download
```

## Running the Service

### Start the Service

```bash
go run cmd/main.go
```

The service will be accessible at `http://localhost:8080`.

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

RUN go build -o /api-gateway

EXPOSE 8080

CMD ["/api-gateway"]
```

### Kubernetes Deployment

Create a Kubernetes deployment file:

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: api-gateway
spec:
  replicas: 3
  selector:
    matchLabels:
      app: api-gateway
  template:
    metadata:
      labels:
        app: api-gateway
    spec:
      containers:
        - name: api-gateway
          image: api-gateway:latest
          ports:
            - containerPort: 8080
```

Create a Kubernetes service file:

```yaml
apiVersion: v1
kind: Service
metadata:
  name: api-gateway
spec:
  selector:
    app: api-gateway
  ports:
    - protocol: TCP
      port: 8080
      targetPort: 8080
```

### Deploy to Kubernetes

```bash
kubectl apply -f deployments/api-gateway-deployment.yaml
kubectl apply -f deployments/api-gateway-service.yaml
```
