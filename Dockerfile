FROM golang:1.23-alpine

WORKDIR /app

COPY . .

RUN go mod tidy

EXPOSE 8080
