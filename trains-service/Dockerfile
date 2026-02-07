FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY . .

RUN go build -o trains-service ./trainss-service/cmd/server

FROM alpine:3.19
WORKDIR /app

COPY --from=builder /app/trains-service .
COPY --from=builder /app/shared/database/scripts ./shared/database/scripts

EXPOSE 8082
CMD ["./trains-service"]
