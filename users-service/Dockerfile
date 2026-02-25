FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY . .

WORKDIR /app/users-service
RUN go build -o users-services ./cmd/server

FROM alpine:3.19
WORKDIR /app

COPY --from=builder /app/users-service .

EXPOSE 8081
CMD ["./users-service"]
