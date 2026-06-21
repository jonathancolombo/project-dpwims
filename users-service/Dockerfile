FROM golang:1.25-alpine AS builder
WORKDIR /app
COPY . .
WORKDIR /app/users-service
RUN go build -o users-service ./cmd/server

FROM alpine:3.19
WORKDIR /app

RUN apk add --no-cache curl

COPY --from=builder /app/users-service .
COPY users-service/scripts/seed-users.sh /app/seed-users.sh
RUN sed -i 's/\r$//' /app/seed-users.sh
RUN chmod +x /app/seed-users.sh

EXPOSE 8081

CMD ["sh", "-c", "/app/seed-users.sh & ./users-service"]