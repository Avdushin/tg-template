FROM golang:1.21-alpine AS builder

WORKDIR /app

# Скопировать модули и исходники
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bot ./cmd/bot

# Финальный образ
FROM alpine:3.18
WORKDIR /app
COPY --from=builder /app/bot .
COPY scripts/wait-for-db.sh ./wait-for-db.sh
RUN chmod +x ./wait-for-db.sh

ENV DB_HOST=db
ENV DB_PORT=5432
ENV DB_USER=tguser
ENV DB_PASSWORD=tgpassword
ENV DB_NAME=tgdb

CMD ["sh", "-c", "./wait-for-db.sh ${DB_HOST}:${DB_PORT} ./bot"]
