# Stage 1: Build the binary
FROM golang:1.17 AS builder

WORKDIR /app

COPY . .

# Build the Go binary
RUN CGO_ENABLED=0 GOOS=linux go build -o myapp .

# Stage 2: Create a minimal image
FROM alpine:latest

WORKDIR /app

# Copy only the necessary files from the previous stage
COPY --from=builder /app/myapp .

# Expose the port your application will run on
EXPOSE 8080

# Command to run your application
CMD ["./myapp"]