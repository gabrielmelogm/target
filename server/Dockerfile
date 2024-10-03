FROM golang:1.22 AS builder

ENV GOOS=linux \
    GOARCH=amd64 \
    CGO_ENABLED=0

WORKDIR /build

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o /app

FROM ghcr.io/surnet/alpine-node-wkhtmltopdf:latest

ENV GIN_MODE=release

WORKDIR /

COPY --from=builder /build/templates /templates
RUN mkdir /outputs
COPY --from=builder /app /app

EXPOSE 8080

CMD [ "/app" ]

