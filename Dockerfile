FROM golang:1.25 AS builder

WORKDIR /app
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o password-generator main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/password-generator .
CMD ["./password-generator"]
