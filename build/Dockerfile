# Build stage
FROM golang:1.16-alpine AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY src/go.mod src/go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY src/ .

# Build the Go app
RUN go build -o service-exporter main.go metrics.go service_discovery.go

# Final stage
FROM alpine:latest

# Set the Current Working Directory inside the container
WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/service-exporter .
COPY config/config.yaml /etc/service-exporter/config.yaml

# Expose port 5555 to the outside world
EXPOSE 5555

# Command to run the executable
CMD ["./service-exporter"]
