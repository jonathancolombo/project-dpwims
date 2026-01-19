FROM golang:1.25 AS build

WORKDIR /app

COPY go.mod ./

RUN go mod download || true

COPY . .

RUN go build -o server ./cmd/api

FROM debian:stable-slim
WORKDIR /app
COPY --from=build /app/server .
EXPOSE 8080
CMD ["./server"]
