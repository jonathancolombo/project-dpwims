FROM golang:1.25-alpine AS builder

WORKDIR /app
COPY . .
RUN go build -o user-service ./cmd/server

FROM alpine:3.19
WORKDIR /app
COPY --from=builder /app/user-service .
EXPOSE 8081
CMD ["./user-service"]
