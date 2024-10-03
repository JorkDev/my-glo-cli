# Use the official Golang image as the base image
FROM golang:1.18 as builder

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files and download the dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application files
COPY . .

# Build the Go application
RUN go build -o mycli

# Create a new minimal image to ship the application
FROM alpine:latest

# Set the working directory in the final container
WORKDIR /

# Copy the compiled binary from the builder stage
COPY --from=builder /app/mycli .

# Run the CLI when the container starts
ENTRYPOINT ["./mycli"]
