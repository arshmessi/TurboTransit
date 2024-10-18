# Dockerfile for user-service
FROM golang:1.17-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o /driver-service

EXPOSE 8080

CMD ["/driver-service"]