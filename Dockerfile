# -------------------------
# Build stage
# -------------------------
FROM golang:1.26 AS builder

WORKDIR /app

# Copy dependency files first (better caching)
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build binary (static)
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# -------------------------
# Runtime stage (small image)
# -------------------------
FROM alpine:latest

WORKDIR /app

# Install CA certificates (needed for HTTPS calls if any)
RUN apk --no-cache add ca-certificates

# Copy binary, static files and logo from builder
COPY --from=builder /app/main .
COPY --from=builder /app/static ./static
COPY --from=builder /app/logo.png .

# Expose port
EXPOSE 8081

# Run app
CMD ["./main"]
