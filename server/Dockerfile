FROM golang:1.24.1 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o server ./main.go

# Финальный образ
FROM debian:bookworm

WORKDIR /app

COPY --from=builder /app/server /app/server

CMD ["/app/server"]