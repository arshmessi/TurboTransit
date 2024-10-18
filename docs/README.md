# TurboTransit

A robust, scalable, and maintainable logistics platform built using microservices architecture.

## Table of Contents

1. [Overview](#overview)
2. [Architecture](#architecture)
3. [Services](#services)
4. [Setup](#setup)
5. [Running the Project](#running-the-project)
6. [Testing](#testing)
7. [Deployment](#deployment)
8. [Monitoring and Observability](#monitoring-and-observability)
9. [Security](#security)
10. [Performance Optimization](#performance-optimization)
## Overview

The Logistics Platform is designed to efficiently handle high loads and real-time updates. Built on a microservices architecture, each service is modular and scalable. The platform manages core logistics functionalities including user management, authentication, driver and vehicle management, booking, matching drivers with bookings, real-time tracking, pricing, and administrative operations.

## Architecture

The platform uses a microservices architecture where services communicate via REST APIs and NATS for event-driven messaging. The **API Gateway** serves as the entry point for external requests, ensuring proper routing, authentication, and rate limiting. Each microservice is responsible for its domain-specific tasks, making the architecture highly scalable and maintainable.

## Services

- **User Service**: Manages user registration, profiles, and updates.
- **Authentication Service**: Handles login, logout, and token management.
- **Driver Service**: Manages drivers, their status, and vehicle assignments.
- **Booking Service**: Handles booking creation, updates, and cancellations.
- **Matching Service**: Matches drivers to bookings.
- **Tracking Service**: Tracks real-time driver and booking locations.
- **Pricing Service**: Calculates prices for bookings.
- **Admin Service**: Provides admin controls for managing users, drivers, and bookings.
- **API Gateway**: Routes requests and enforces security policies.

## Setup

### Clone the Repository

```bash
git clone https://github.com/arshmessi/TurboTransit.git
cd TurboTransit
```

### Install Dependencies

For each microservice, navigate to its directory and run:

```bash
go mod download
```

### Database Setup

For services that need a database (e.g., User, Driver, Booking Services), set up SQLite and create necessary tables. For example, for the User Service:

```bash
sqlite3 users.db
```

```sql
CREATE TABLE users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    username TEXT NOT NULL UNIQUE,
    email TEXT NOT NULL UNIQUE,
    password_hash TEXT NOT NULL,
    first_name TEXT,
    last_name TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

### Redis and NATS Setup

Ensure Redis is running at `localhost:6379` and NATS is running at `localhost:4222`.

## Running the Project

Start each service by navigating to its directory and running:

```bash
go run cmd/main.go
```

For example, to start the User Service:

```bash
cd user-service
go run cmd/main.go
```

Repeat this process for all services.

## Testing

To run unit tests, navigate to a service’s directory and run:

```bash
go test ./...
```

For integration tests:

```bash
go test ./internal/controller
```

Ensure the respective service is running before running integration tests.

## Deployment

### Docker Setup

Create a `Dockerfile` for each service. Here is an example for the User Service:

```dockerfile
FROM golang:1.17-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o /user-service

EXPOSE 8080

CMD ["/user-service"]
```

### Kubernetes Deployment

Create Kubernetes deployment and service files for each service. Here’s an example for the User Service:

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: user-service
spec:
  replicas: 3
  selector:
    matchLabels:
      app: user-service
  template:
    metadata:
      labels:
        app: user-service
    spec:
      containers:
        - name: user-service
          image: user-service:latest
          ports:
            - containerPort: 8080
```

```yaml
apiVersion: v1
kind: Service
metadata:
  name: user-service
spec:
  selector:
    app: user-service
  ports:
    - protocol: TCP
      port: 8080
      targetPort: 8080
```

Deploy using:

```bash
kubectl apply -f deployments/user-service-deployment.yaml
kubectl apply -f deployments/user-service-service.yaml
```

Repeat this for all services.

## Monitoring and Observability

- **Logging**: Use the ELK stack (Elasticsearch, Logstash, Kibana) for centralized logging.
- **Metrics**: Use Prometheus for collecting metrics, and Grafana for visualization.
- **Tracing**: Implement Jaeger for distributed tracing across services.

## Security

- **RBAC (Role-Based Access Control)**: Implemented in the Authentication Service to restrict user access.
- **Data Encryption**: Encrypt sensitive data such as user passwords and personal information using industry-standard algorithms.

## Performance Optimization

- **Caching**: Use Redis for multi-level caching to reduce load on databases.
- **Load Balancing**: Use Kubernetes to manage scaling and load balancing across instances of each service.

