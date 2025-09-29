FROM golang:1.25.1-alpine AS builder

WORKDIR /app

ENV GOPROXY=https://proxy.golang.org,direct

COPY go.mod go.sum ./
RUN go mod tidy

COPY . ./
RUN go build -o bot ./cmd/bot

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/bot .
COPY .env .env

CMD ["./bot"]
