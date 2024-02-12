# syntax=docker/dockerfile:1

FROM golang:1.19 AS builder

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download
COPY .env /app/.env

# Copy the source code
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /my-exhibition-app ./cmd/exhibition

FROM node:16-alpine

EXPOSE 8080

# Copy the binary from the builder stage
COPY --from=builder /my-exhibition-app /my-exhibition-app

# Run
CMD ["/my-exhibition-app"]
