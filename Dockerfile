FROM golang:1.25.1-alpine AS builder

WORKDIR /app

# Устанавливаем swag для генерации документации
RUN go install github.com/swaggo/swag/cmd/swag@latest

COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod/cache \
  go mod download

COPY . .

# Генерируем Swagger документацию
RUN swag init -g internal/delivery/http/v1/handler.go -o docs

RUN go build -o main ./cmd/app

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /var/www

COPY --from=builder /app/main .
COPY --from=builder /app/docs ./docs

RUN chmod +x main

EXPOSE 8080

CMD ["./main"]
