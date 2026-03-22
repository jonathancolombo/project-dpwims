FROM golang:1.25-alpine AS builder
LABEL authors="jonathan"

WORKDIR /app
COPY . .

WORKDIR /app/auth-service
RUN go build -o auth-service ./cmd/server

FROM alpine:3.19
WORKDIR /app
COPY --from=builder /app/auth-service .
EXPOSE 8085
CMD ["./auth-service"]