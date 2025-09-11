# Stage 1: Build the Go binary
FROM golang:1.21 AS builder

# Set working directory inside container
WORKDIR /app

# Copy go.mod and go.sum first to leverage caching - helps Docker cache dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the Go app (targeting your web entry point)
RUN go build -o hourly-pay-item-calculator ./cmd/web

# Stage 2: Create a lightweight container
FROM alpine:latest

# Set working directory in final container
WORKDIR /app

# Copy the compiled binary from builder stage
COPY --from=builder /app/hourly-pay-item-calculator .

# Expose port 8080 to the host
EXPOSE 8080

# Run the binary
CMD ["./hourly-pay-item-calculator"]