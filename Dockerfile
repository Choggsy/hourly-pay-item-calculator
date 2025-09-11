# Stage 1: Build the Go binary, matches my Go.Mod file version
FROM golang:1.25 AS builder

# Set working directory inside container
WORKDIR /app

# Copy go.mod and go.sum first to leverage caching - helps Docker cache dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the Go app (targeting your web entry point)
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/hourly-pay-item-calculator ./cmd/web

# Stage 2: Create a lightweight container
FROM alpine:latest

# Set working directory in final container
WORKDIR /app

# Copy the compiled binary
COPY --from=builder /app/hourly-pay-item-calculator .

# Copy the HTML template
COPY --from=builder /app/web/templates /app/web/templates

EXPOSE 8080

CMD ["/app/hourly-pay-item-calculator"]