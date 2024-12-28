FROM golang:1.22 AS builder

ENV GOOS=linux \
    GOARCH=amd64 \
    CGO_ENABLED=0

WORKDIR /build

COPY go.mod ./
COPY go.sum ./
RUN go mod tidy

COPY . .

RUN go build -o /app

FROM alpine:3.18

RUN apk add --no-cache make bash

ENV GIN_MODE=release

RUN echo "ENVIRONMENT=production" > .env && \
    echo "DB_DRIVER=${DB_DRIVER}" >> .env && \
    echo "DB_NAME=${DB_NAME}" >> .env && \
    echo "DB_HOST=${DB_HOST}" >> .env && \
    echo "DB_PORT=${DB_PORT}" >> .env && \
    echo "DB_USER=${DB_USER}" >> .env && \
    echo "DB_PASS=${DB_PASS}" >> .env && \
    echo "DB_SSL_MODE=${DB_SSL_MODE}" >> .env

WORKDIR /

COPY --from=builder /build/makefile /makefile
COPY --from=builder /app /app

EXPOSE 8080

CMD [ "/app" ]

