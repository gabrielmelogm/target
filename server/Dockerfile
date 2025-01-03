FROM golang:1.22 AS builder

ENV GOOS=linux \
    GOARCH=amd64 \
    CGO_ENABLED=0

WORKDIR /build

COPY go.mod ./
COPY go.sum ./
RUN go mod tidy

COPY . .

RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

RUN go build -o /app

FROM alpine:3.18

RUN apk add --no-cache make bash curl

COPY --from=builder /go/bin/migrate /usr/local/bin/migrate

ENV GIN_MODE=release

ARG DB_URL

RUN echo "ENVIRONMENT=production" > .env && \
    echo "DB_URL=${DB_URL}" >> .env

WORKDIR /

COPY --from=builder /app /app
COPY --from=builder /build/db/migrations /db/migrations
COPY --from=builder /build/makefile /makefile

EXPOSE 8080

CMD [ "/app" ]

