# Use the official Go image
FROM golang:1.23 AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the go mod and sum files
COPY go.mod ./

# Download all dependencies
RUN go mod download

# Copy the source code into the container
COPY main.go ./

# Build the Go apps
RUN CGO_ENABLED=0 GOOS=linux go build -o welcome-app main.go

# Start a new stage from scratch
FROM alpine:latest

# Set the Current Working Directory inside the container
WORKDIR /root/

# Copy the Pre-built binary files from the previous stage
COPY --from=builder /app/welcome-app .

# Expose ports 
EXPOSE 3060

# Command to run the welcome-app executable
CMD ["./welcome-app"]
