FROM golang:1.25

WORKDIR /app
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go install ./service.generator

FROM alpine:latest
WORKDIR /app
COPY --from=0 /go/bin/service.generator .
CMD ["./service.generator"]
