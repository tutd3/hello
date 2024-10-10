# Use the official Go image
FROM golang:1.23 AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code into the container
COPY main.go ./
COPY handlers/ ./handlers/ 

# Copy the .env file into the container
COPY .env ./

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -o welcome-app main.go

# Start a new stage from scratch
FROM alpine:latest

# Set the Current Working Directory inside the container
WORKDIR /root/

# Copy the Pre-built binary files from the previous stage
COPY --from=builder /app/welcome-app .

# Copy the .env file into the container for runtime
COPY --from=builder /app/.env ./

# Expose ports
EXPOSE 8080

# Command to run the welcome-app executable
CMD ["./welcome-app"]

