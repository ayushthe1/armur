# Start with the official Golang image for building the binary
FROM golang:1.22 as builder

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /clients-service .

# Final stage: a minimal base image
FROM alpine:3.18

# Set up a working directory
WORKDIR /app

# Copy the built binary from the previous stage
COPY --from=builder /clients-service /app/clients-service

# Expose the service port
EXPOSE 50053

# Command to run the service
CMD ["/app/clients-service"]
