# syntax=docker/dockerfile:1

FROM golang:1.19 AS builder


# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .


RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-gs-ping ./cmd/exhibition

FROM node:16-alpine

EXPOSE 8080

# Run
CMD ["/docker-gs-ping"]