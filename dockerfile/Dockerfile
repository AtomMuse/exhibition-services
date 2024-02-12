# Use the official Golang image as a base image
FROM golang:1.19-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules and the Go files into the container
COPY go.mod go.sum ./
COPY . .

# Build the Go application
RUN go build -o /my-exhibition-app ./cmd/exhibition

# Start a new stage for a small image
FROM alpine:3.14

# Set the working directory inside the container
WORKDIR /app

# Copy the built executable from the builder stage
COPY --from=builder /my-exhibition-app .

# Copy the .env file
COPY .env .

# Expose the port the application runs on
EXPOSE 8080

# Command to run the executable
CMD ["./my-exhibition-app"]
