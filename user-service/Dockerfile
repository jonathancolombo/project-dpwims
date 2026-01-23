FROM golang:1.25 AS builder

WORKDIR /app
COPY . .
RUN go build -o user-service ./cmd/server

FROM debian:stable-slim
WORKDIR /app
COPY --from=builder /app/user-service .
EXPOSE 8081
CMD ["./user-service"]