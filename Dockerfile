FROM golang:1.22.5-alpine AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
ENV GIN_MODE release
ENV DB_HOST website-mysql
ENV DB_PORT 3306
ENV DB_USER root
ENV DB_PASSWORD ly@15984093508
ENV DB_NAME blog
RUN go build -o main ./cmd/main.go

FROM alpine:latest
WORKDIR /root/
COPY --from=build /app/main .
CMD ["./main"]