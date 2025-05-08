FROM golang:1.24-alpine AS builder

RUN apk update && apk add --no-cache git tzdata

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Create app directory
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the full source code
COPY . .

# Build the Go binary
RUN go build -o main ./cmd/main.go

FROM alpine:latest

# Install CA certificates and timezone info
RUN apk --no-cache add ca-certificates tzdata

# Set working directory
WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/main .

# Copy .env file (optional — prefer mounting in prod)
COPY .env .env

# Expose the server port (from .env)
EXPOSE 8080

# Run the binary
CMD ["./main"]
