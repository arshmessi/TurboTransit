# Logistics Platform Architecture

## 1. Microservice Architecture

### 1.1 User Service

#### Responsibilities:

- User registration and profile management
- User preferences and settings

#### Internal Components:

- **User Controller**: Handles HTTP requests for user-related operations
- **User Repository**: Interfaces with the database for user data operations
- **Profile Manager**: Handles user profile updates and retrieval

#### Events:

- Publishes: UserCreated, UserUpdated
- Subscribes to: BookingCreated (for user history updates)

#### API Endpoints:

- POST /users (registration)
- GET /users/{id} (profile retrieval)
- PUT /users/{id} (profile update)

#### Database Schema:

```sql
CREATE TABLE users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    username TEXT NOT NULL UNIQUE,
    email TEXT NOT NULL UNIQUE,
    password_hash TEXT NOT NULL,
    salt TEXT NOT NULL,
    first_name TEXT,
    last_name TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

### 1.2 Authentication Service

#### Responsibilities:

- User authentication and authorization
- Token generation and validation
- Password hashing and verification

#### Internal Components:

- **Auth Controller**: Handles HTTP requests for authentication operations
- **Token Manager**: Generates and validates authentication tokens
- **Password Manager**: Handles password hashing, salting, and verification

#### Events:

- Publishes: UserAuthenticated, UserLoggedOut
- Subscribes to: UserCreated, UserUpdated

#### API Endpoints:

- POST /auth/login (user login)
- POST /auth/logout (user logout)
- POST /auth/refresh (refresh authentication token)

#### Database Schema:

```sql
CREATE TABLE auth_tokens (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER NOT NULL,
    token TEXT NOT NULL,
    expires_at TIMESTAMP NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id)
);
```

### 1.3 Driver Service

#### Responsibilities:

- Driver registration and profile management
- Driver status management (available, busy, offline)
- Vehicle assignment and management

#### Internal Components:

- **Driver Controller**: Handles HTTP requests for driver-related operations
- **Driver Repository**: Interfaces with the database for driver data operations
- **Status Manager**: Manages driver status updates
- **Vehicle Manager**: Handles vehicle assignments and updates

#### Events:

- Publishes: DriverCreated, DriverStatusUpdated, VehicleAssigned
- Subscribes to: BookingCreated, BookingCancelled

#### API Endpoints:

- POST /drivers (registration)
- GET /drivers/{id} (profile retrieval)
- PUT /drivers/{id}/status (status update)
- POST /drivers/{id}/vehicles (vehicle assignment)

#### Database Schema:

```sql
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
```

### 1.4 Booking Service

#### Responsibilities:

- Creating and managing bookings
- Handling booking status updates

#### Internal Components:

- **Booking Controller**: Handles HTTP requests for booking-related operations
- **Booking Repository**: Interfaces with the database for booking data operations
- **Pricing Calculator**: Calculates the price for a booking (might interact with Pricing Service)

#### Events:

- Publishes: BookingCreated, BookingUpdated, BookingCancelled
- Subscribes to: DriverStatusUpdated, UserCreated, LocationUpdated, MatchingCompleted

#### API Endpoints:

- POST /bookings (create booking)
- GET /bookings/{id} (retrieve booking details)
- PUT /bookings/{id}/status (update booking status)
- GET /bookings/user/{userId} (retrieve user's bookings)

#### Database Schema:

```sql
CREATE TABLE bookings (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER NOT NULL,
    driver_id INTEGER,
    pickup_location TEXT NOT NULL,
    dropoff_location TEXT NOT NULL,
    status TEXT CHECK(status IN ('pending', 'accepted', 'in_progress', 'completed', 'cancelled')) DEFAULT 'pending',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (driver_id) REFERENCES drivers(id)
);
```

### 1.5 Matching Service

#### Responsibilities:

- Implementing driver matching algorithm
- Optimizing matches based on various factors (location, driver rating, etc.)

#### Internal Components:

- **Matching Controller**: Handles requests for driver-booking matches
- **Matching Engine**: Implements the algorithm for matching drivers to bookings
- **Optimization Manager**: Continuously improves matching based on historical data

#### Events:

- Publishes: MatchingCompleted, MatchingFailed
- Subscribes to: BookingCreated, DriverStatusUpdated, LocationUpdated

#### API Endpoints:

- POST /matching/find-driver (find a suitable driver for a booking)
- GET /matching/statistics (retrieve matching statistics)

#### Database Schema:

```sql
CREATE TABLE matching_history (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    booking_id INTEGER NOT NULL,
    driver_id INTEGER NOT NULL,
    match_score REAL NOT NULL,
    accepted BOOLEAN,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (booking_id) REFERENCES bookings(id),
    FOREIGN KEY (driver_id) REFERENCES drivers(id)
);
```

### 1.6 Tracking Service

#### Responsibilities:

- Real-time GPS tracking of vehicles
- Updating and storing location data
- Calculating ETAs

#### Internal Components:

- **Tracking Controller**: Handles HTTP requests for location updates and retrieval
- **Location Repository**: Interfaces with the database for location data operations
- **ETA Calculator**: Calculates estimated time of arrival based on current location and destination

#### Events:

- Publishes: LocationUpdated, ETAUpdated
- Subscribes to: BookingCreated, BookingUpdated

#### API Endpoints:

- POST /locations (update location)
- GET /locations/driver/{driverId} (get driver's current location)
- GET /locations/booking/{bookingId} (get location updates for a booking)

#### Database Schema:

```sql
CREATE TABLE locations (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    driver_id INTEGER NOT NULL,
    booking_id INTEGER,
    latitude REAL NOT NULL,
    longitude REAL NOT NULL,
    timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (driver_id) REFERENCES drivers(id),
    FOREIGN KEY (booking_id) REFERENCES bookings(id)
);
```

### 1.7 Pricing Service

#### Responsibilities:

- Calculating fares for bookings
- Managing surge pricing
- Providing price estimates

#### Internal Components:

- **Pricing Controller**: Handles HTTP requests for pricing-related operations
- **Fare Calculator**: Calculates fares based on distance, time, and other factors
- **Surge Pricing Manager**: Implements surge pricing logic based on demand and supply
- **Estimator**: Provides price estimates for potential bookings

#### Events:

- Publishes: PriceUpdated
- Subscribes to: BookingCreated, LocationUpdated

#### API Endpoints:

- POST /pricing/calculate (calculate fare for a booking)
- GET /pricing/estimate (get price estimate for a potential booking)
- GET /pricing/surge-status (get current surge pricing status)

#### Database Schema:

```sql
CREATE TABLE pricing_rules (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    vehicle_type TEXT NOT NULL,
    base_fare REAL NOT NULL,
    per_km_rate REAL NOT NULL,
    per_minute_rate REAL NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE surge_pricing_history (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    multiplier REAL NOT NULL,
    start_time TIMESTAMP NOT NULL,
    end_time TIMESTAMP,
    reason TEXT
);
```

### 1.8 Admin Service

#### Responsibilities:

- Managing system-wide settings
- Generating reports and analytics
- Handling user and driver management tasks

#### Internal Components:

- **Admin Controller**: Handles HTTP requests for admin-related operations
- **Report Generator**: Generates various system reports
- **User Manager**: Handles administrative tasks related to users
- **Driver Manager**: Handles administrative tasks related to drivers

#### Events:

- Publishes: SettingsUpdated, UserBanned, DriverSuspended
- Subscribes to: UserCreated, DriverCreated, BookingCompleted

#### API Endpoints:

- GET /admin/dashboard (retrieve admin dashboard data)
- POST /admin/users/{id}/ban (ban a user)
- POST /admin/drivers/{id}/suspend (suspend a driver)
- GET /admin/reports/bookings (generate booking report)

#### Database Schema:

```sql
CREATE TABLE admin_users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    username TEXT NOT NULL UNIQUE,
    email TEXT NOT NULL UNIQUE,
    password_hash TEXT NOT NULL,
    salt TEXT NOT NULL,
    role TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE system_settings (
    key TEXT PRIMARY KEY,
    value TEXT NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

## 2. System Components

### 2.1 API Gateway

#### Responsibilities:

- Request routing
- Authentication and authorization (in coordination with Auth Service)
- Rate limiting
- Request/response transformation

#### Technology Choice: Go with the gin framework

- Reasons:
  - High performance and low latency
  - Built-in middleware for common tasks
  - Easy to extend and customize
- Pros:
  - Excellent performance characteristics
  - Strong community support
  - Easy to learn and use
- Cons:
  - May require additional libraries for some advanced features
- Alternative: Node.js with Express.js
  - Pros: Large ecosystem, easy to find developers
  - Cons: Potentially lower performance compared to Go

### 2.2 Event Bus (NATS)

#### Responsibilities:

- Asynchronous communication between microservices
- Event publishing and subscription
- Message persistence (optional)

#### Technology Choice: NATS

- Reasons:
  - High performance and low latency
  - Supports various messaging patterns (pub/sub, request/reply)
  - Built-in support for clustering and fault tolerance
- Pros:
  - Extremely fast and lightweight
  - Supports millions of messages per second
  - Easy to set up and operate
- Cons:
  - Less feature-rich compared to some alternatives
  - Limited built-in persistence options
- Alternative: Apache Kafka
  - Pros: Excellent for high-throughput scenarios, strong persistence guarantees
  - Cons: More complex to set up and operate, higher resource requirements

### 2.3 Database Layer

#### Technology Choice: SQLite with sqlc

- Reasons:
  - Simplicity and ease of setup
  - No separate database server required
  - Good performance for moderate workloads
- Pros:
  - Zero configuration required
  - ACID compliant
  - Suitable for development and small to medium-scale deployments
- Cons:
  - Limited concurrency support
  - Not suitable for high write loads or large datasets
- Alternative: PostgreSQL
  - Pros: Excellent for complex queries, high concurrency, and large datasets
  - Cons: Requires separate server setup, more complex to manage

### 2.4 Caching Layer

#### Technology Choice: Redis

- Reasons:
  - In-memory data structure store
  - Support for various data structures
  - Built-in persistence and cluster support
- Pros:
  - Extremely fast read/write operations
  - Versatile (can be used for caching, queues, pub/sub)
  - Supports data structures like lists, sets, and sorted sets
- Cons:
  - Data size limited by available memory
  - Potential for data loss in case of crashes (if persistence is not configured properly)
- Alternative: Memcached
  - Pros: Simpler, potentially faster for basic key-value caching
  - Cons: Less feature-rich, no built-in persistence

## 3. Scalability and Performance Considerations

### 3.1 Horizontal Scaling

- Each microservice can be independently scaled based on load
- Use of container orchestration (e.g., Kubernetes) for easy scaling and management

### 3.2 Database Sharding

- As data grows, implement sharding strategies for SQLite databases
- Consider moving to a distributed database system for very large datasets

### 3.3 Caching Strategy

- Implement multi-level caching (application-level and distributed cache)
- Use cache-aside pattern for reading data
- Implement write-through or write-behind caching for writes

### 3.4 Asynchronous Processing

- Use event-driven architecture to handle high-volume operations asynchronously
- Implement background job processing for time-consuming tasks

### 3.5 Load Balancing

- Implement load balancing at the API Gateway level
- Use service discovery for dynamic routing between services

## 4. Fault Tolerance and Reliability

### 4.1 Circuit Breaker Pattern

- Implement circuit breakers for inter-service communication to prevent cascading failures

### 4.2 Retry Mechanism

- Implement intelligent retry mechanisms with exponential backoff for transient failures

### 4.3 Data Replication

- Implement data replication for critical data to ensure availability

### 4.4 Monitoring and Alerting

- Use Prometheus for metrics collection
- Set up Grafana dashboards for visualization
- Implement alerting for critical system issues

## 5. Security Considerations

### 5.1 Authentication and Authorization

- Implement JWT-based authentication using the dedicated Auth Service
- Use role-based access control (RBAC) for fine-grained permissions
- Implement password salting and hashing for user and driver accounts

### 5.2 Data Encryption

- Encrypt sensitive data at rest and in transit
- Use HTTPS for all external communications

### 5.3 Rate Limiting and DDoS Protection

- Implement rate limiting at the API Gateway level
- Consider using a CDN or specialized DDoS protection service for large-scale deployments

### 5.4 Regular Security Audits

- Conduct regular security audits and penetration testing
- Keep all systems and libraries up to date with security patches

## 6. Deployment and DevOps

### 6.1 Containerization

- Use Docker for containerizing microservices
- Implement Docker Compose for local development environments

### 6.2 Orchestration

- Use Kubernetes for container orchestration in production
- Implement Helm charts for easy deployment and upgrades

### 6.3 Continuous Integration/Continuous Deployment (CI/CD)

- Implement CI/CD pipelines using tools like Jenkins or GitLab CI
- Automate testing, building, and deployment processes

### 6.4 Infrastructure as Code

- Use tools like Terraform for managing infrastructure
- Implement GitOps practices for infrastructure management

## 7. Monitoring and Observability

### 7.1 Logging

- Implement centralized logging using the ELK stack (Elasticsearch, Logstash, Kibana)
- Use structured logging for easier parsing and analysis

### 7.2 Metrics

- Use Prometheus for collecting and storing metrics
- Set up Grafana dashboards for visualizing system performance

### 7.3 Tracing

- Implement distributed tracing using Jaeger or Zipkin
- Trace requests across microservices for performance analysis and debugging

### 7.4 Alerting

- Set up alerting rules in Prometheus
- Integrate with PagerDuty or similar services for on-call management

## 8. Data Management and Analytics

### 8.1 Data Warehousing

- Implement a data warehouse for historical data analysis
- Consider using tools like Apache Airflow for ETL processes

### 8.2 Real-time Analytics

- Use stream processing tools like Apache Flink for real-time data analysis
- Implement real-time dashboards for business intelligence

### 8.3 Machine Learning Integration

- Integrate machine learning models for demand prediction and dynamic pricing
- Use tools like TensorFlow Serving for model deployment

## 9. Conclusion and System Overview

The logistics platform architecture has been thoughtfully designed to be both robust and scalable, leveraging modern microservices and event-driven principles to deliver real-time, efficient, and flexible operations. Key updates and features reflect the improvements and strategic decisions made in this iteration:

### Core Architecture and Design Principles

- **Microservices Architecture**: The platform is built using Go for microservices, ensuring high performance and efficient resource utilization. The event-driven nature, powered by NATS, facilitates real-time updates and loose coupling between services, critical for features like real-time tracking and instant notifications about booking status changes.
- **Service Isolation**:

  - **Authentication Service**: By separating authentication into its own service, the system achieves better modularity, allowing for easier scaling and maintenance of authentication-related functionality.
  - **Matching Service**: The matching engine has been turned into its own dedicated service, improving focus on development, optimization, and scalability for this computationally intensive task.

- **Dedicated Services for Key Operations**:
  - **User and Driver Services**: These manage the core entities of the system, providing clear boundaries and responsibilities.
  - **Booking Service**: Manages critical business logic for creating and handling transportation requests.
  - **Tracking Service**: Provides real-time location updates, enhancing the logistics experience.
  - **Pricing Service**: Implements dynamic and flexible pricing strategies.

### Enhancements and New Additions

- **New Admin Service**: Introduced to centralize system-wide management tasks, reporting, and analytics, without interfering with core business logic. This opens opportunities for business intelligence through advanced reporting and analytics.
- **Updated API Gateway**: The API Gateway, using gin (formerly gin-gonic), serves as the single entry point for all client requests, managing critical operations like authentication, rate limiting, and request routing, while enforcing security policies centrally.

- **Security**:
  - Password salting for users and drivers has been implemented to enhance credential security.
  - Comprehensive security measures, including encryption, authentication, and authorization, have been implemented at various levels, and regular security audits are crucial for maintaining a strong security posture.

### Scalability and Future-Proofing

- **Database Layer**: SQLite with sqlc offers simplicity and efficiency for the initial deployment, but the architecture leaves room for scaling to distributed databases as the system grows.
- **Caching with Redis**: A caching layer is implemented using Redis, improving data retrieval speed for frequently used data and reducing load on the primary databases.

- **Fault Tolerance and Monitoring**: With patterns like circuit breakers and retry mechanisms, the system is designed to handle faults gracefully. Docker and Kubernetes support easy service management and scaling. Additionally, the comprehensive monitoring stack ensures timely issue detection and resolution through logging, metrics, and tracing.

### Future Considerations

While the current architecture provides a solid foundation, there is room for future enhancements:

- **Analytics and Machine Learning**: Future integration of advanced analytics and machine learning can offer valuable insights, optimize matching, pricing, and operational efficiency.
- **Scaling and Optimization**: As the platform scales, a move to distributed database solutions and advanced caching strategies will be necessary. Regular architecture reviews, performance testing, and monitoring will be key to identifying bottlenecks and driving continuous improvements.

In conclusion, this system design provides a solid, scalable, and secure foundation for a logistics platform. As the platform evolves, continuous evaluation, refinement, and iterative improvements based on real-world usage patterns and business needs will ensure long-term success.
