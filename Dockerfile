FROM golang:1.19 AS builder

WORKDIR /app

COPY main.go .
COPY go.mod .

RUN CGO_ENABLED=0 GOOS=linux go build

FROM alpine:3.10

WORKDIR /app

COPY --from=builder /app/hello-go .

CMD ["./hello-go"]