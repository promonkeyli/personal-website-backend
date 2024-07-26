FROM golang:1.22.5-alpine AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main ./cmd/main.go
FROM alpine:latest
WORKDIR /root/
COPY --from=build /app/main .
ENV GIN_MODE=release \
    DB_HOST=website-mysql \
    DB_PORT=3306 \
    DB_USER=root \
    DB_PASSWORD=ly15984093508 \
    DB_NAME=blog
CMD ["./main"]