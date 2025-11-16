# ---------------------
# 1. Build Stage
# ---------------------
FROM golang:1.25 AS builder

WORKDIR /app

# Copy go.mod and go.sum
COPY go.mod go.sum ./
RUN go mod download

# Copy
COPY . .

# Compile main.go (provided by build args)
ARG SERVICE_PATH
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o service-binary ${SERVICE_PATH}

# ---------------------
# 2. Run Stage
# ---------------------
FROM --platform=linux/amd64 alpine:latest

WORKDIR /app

RUN apk add --no-cache tzdata
RUN apk add --no-cache bash

COPY --from=builder /app/service-binary ./service

COPY config/config.docker.ini ./config.docker.ini

COPY wait-for-it.sh /wait-for-it.sh
RUN chmod +x /wait-for-it.sh

EXPOSE 4000

ENV CONFIG_PATH=/app/config.docker.ini

CMD ["bash", "/wait-for-it.sh", "mysql:3306", "--", "./service"]