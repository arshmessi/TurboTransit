# Authentication Service

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

The Authentication Service is responsible for managing user authentication and token generation. It interacts with the User Service to validate user credentials and generates JWT tokens for authenticated users.

## API Endpoints

- **POST /auth/login**: User login and token generation.
- **POST /auth/logout**: User logout.
- **POST /auth/refresh**: Refresh authentication token.

## Dependencies

- Go 1.17+
- Redis
- Gin (for routing)

## Setup

### Install Dependencies

```bash
go mod download
```

### Redis Setup

Ensure Redis is running on `localhost:6379`.

## Running the Service

### Start the Service

```bash
go run cmd/main.go
```

The service will be available at `http://localhost:8081`.

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

RUN go build -o /auth-service

EXPOSE 8081

CMD ["/auth-service"]
```

### Kubernetes Deployment

Create a Kubernetes deployment file:

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: auth-service
spec:
  replicas: 3
  selector:
    matchLabels:
      app: auth-service
  template:
    metadata:
      labels:
        app: auth-service
    spec:
      containers:
        - name: auth-service
          image: auth-service:latest
          ports:
            - containerPort: 8081
```

Create a Kubernetes service file:

```yaml
apiVersion: v1
kind: Service
metadata:
  name: auth-service
spec:
  selector:
    app: auth-service
  ports:
    - protocol: TCP
      port: 8081
      targetPort: 8081
```

### Deploy to Kubernetes

```bash
kubectl apply -f deployments/auth-service-deployment.yaml
kubectl apply -f deployments/auth-service-service.yaml
```

```

```
