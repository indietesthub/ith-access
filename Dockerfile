# Use the official Golang image as the base image
FROM golang:1.24-alpine AS builder

# Set the working directory
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/bin/api cmd/api/main.go

# Use a minimal alpine image for the final stage
FROM alpine:latest

# Install bash and PostgreSQL client for SSL certificates
RUN apk add --no-cache bash postgresql-client

# Set working directory
WORKDIR /app

RUN mkdir -p /logs && chmod 775 /logs

# Copy the binary from builder
COPY --from=builder /app/bin/api .
COPY --from=builder /app/.env .
COPY --from=builder /app/internal/database/migrations ./internal/database/migrations
COPY entrypoint.sh .

# Make the entrypoint script executable
RUN chmod +x /app/entrypoint.sh

# Expose the port the app runs on
EXPOSE 8080


# Command to run the application
CMD ["/bin/bash", "/app/entrypoint.sh"]
