FROM golang:1.25-alpine AS builder
LABEL authors="jonathancolombo"

WORKDIR /app

COPY . .

WORKDIR /app/tickets-service
RUN go build -o tickets-services ./cmd/server

FROM alpine:3.19
WORKDIR /app

COPY --from=builder /app/tickets-service .

EXPOSE 8083
CMD ["./tickets-service"]
