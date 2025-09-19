FROM golang:1.25.1-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod/cache \
  go mod download

COPY . .

RUN go build -o main ./cmd/app

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /var/www

COPY --from=builder /app/main .

RUN chmod +x main

EXPOSE 8080

CMD ["./main"]
