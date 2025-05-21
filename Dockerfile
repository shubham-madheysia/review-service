# === Stage 1: Build ===
FROM golang:1.20-alpine AS builder

# Install git and ca-certificates for Go modules and HTTPS
RUN apk add --no-cache git ca-certificates

# Set workdir
WORKDIR /app

# Copy go.mod and go.sum and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the binary with optimizations and strip debug info
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o review-service ./cmd/main.go

# === Stage 2: Final Image ===
FROM alpine:latest

# Install ca-certificates for HTTPS if needed
RUN apk add --no-cache ca-certificates

# Set non-root user (optional but recommended)
RUN addgroup -S appgroup && adduser -S appuser -G appgroup
USER appuser

WORKDIR /app

# Copy binary from builder
COPY --from=builder /app/review-service .

# Expose ports if applicable (adjust as needed)
EXPOSE 8080

# Healthcheck: simple curl to localhost port 8080 (adjust endpoint if you have a health endpoint)
HEALTHCHECK --interval=30s --timeout=5s --start-period=10s --retries=3 \
  CMD wget --spider --quiet http://localhost:8080/health || exit 1

# Command to run your app
# Using exec form to get proper signal handling and logs in Docker
CMD ["./review-service"]

# === Environment Variables ===
# You can override these when running the container or via docker-compose:
# Example: 
# docker run -e AWS_REGION=us-west-2 -e DB_HOST=your-db-host ...
#
# Make sure your app reads these from environment or config files.
