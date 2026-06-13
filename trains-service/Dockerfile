FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY . .

WORKDIR /app/trains-service
RUN go build -o trains-service ./cmd/server

FROM alpine:3.19
WORKDIR /app

COPY --from=builder /app/trains-service .

EXPOSE 8082
CMD ["./trains-service"]
