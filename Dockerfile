# Use official Golang image as builder
FROM golang:1.22 AS builder

# Set the default working directory similar to cd command on Linux
WORKDIR /app

# Copy go.mod files first for better caching
COPY go.mod go.sum .
RUN go mod download

# Copy startup script
COPY . .

# Build binary
RUN go build -o main .

# Final lightweight image
FROM debian:bookworm-slim

WORKDIR /app
COPY --from=builder /app/main .

# Install postgres client for pg_isready
RUN apt-get update && apt-get install -y postgresql-client && rm -rf /var/lib/apt/lists/*

# Use custom script as entrypoint
COPY wait-and-run.sh .
RUN chmod +x wait-and-run.sh

ENTRYPOINT ["./wait-and-run.sh"]