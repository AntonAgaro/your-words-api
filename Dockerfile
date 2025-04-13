# Use the official Go image as the base image
FROM golang:1.23 as builder

# Set environment variables for Go build
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

# Set workspace directory in the container
WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code into the container
COPY . .


# Build the Go application
RUN go build -o main .

# Use a smaller base image for the final Docker image
FROM alpine:latest

# Set working directory in the runtime container
WORKDIR /root/

# Copy the built binary from the builder phase
COPY --from=builder /app/main .

# Expose the application port
EXPOSE 8080

# Command to run the app
CMD ["./main"]
