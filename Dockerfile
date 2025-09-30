# Builder + runtime в одном образе
FROM golang:1.25.1

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy

COPY . ./
RUN go build -o bot ./cmd/bot

CMD ["./bot"]
