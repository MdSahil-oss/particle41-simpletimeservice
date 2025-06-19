# ===== STAGE 1: Build the Go binary =====
FROM golang:1.24 AS builder

WORKDIR /app

COPY go.mod ./
COPY main.go ./

RUN go mod tidy

# Build the Go binary
# Disables CGO (C bindings). Tells the Go compiler not to use any C libraries or dependencies. Important when you're using minimal Docker base images like alpine or scratch, which don’t hav>
RUN CGO_ENABLED=0 GOOS=linux go build -o simple-time-service

# ===== STAGE 2: Minimal runtime with non-root user =====
FROM alpine:latest

# Create a non-root user and group
RUN addgroup -S appgroup && adduser -S appuser -G appgroup

WORKDIR /app

# Copy binary from builder
COPY --from=builder /app/simple-time-service .

# Use non-root user
USER appuser

# Expose the service port
EXPOSE 8080

# Start the app
ENTRYPOINT ["./simple-time-service"]
