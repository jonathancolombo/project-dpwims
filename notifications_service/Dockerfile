FROM golang:1.25-alpine AS builder
LABEL authors="jonathancolombo"

WORKDIR /app

COPY . .

WORKDIR /app/notifications-service
RUN go build -o notifications-service ./cmd/server

FROM alpine:3.19
WORKDIR /app

COPY --from=builder /app/notifications-service .

EXPOSE 8084
CMD ["./notifications-service"]
