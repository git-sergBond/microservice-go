FROM golang:1.19 AS builder

WORKDIR /opt/time-service

COPY main.go .
COPY go.mod .

RUN CGO_ENABLED=0 GOOS=linux go build

FROM alpine:3.10

EXPOSE 8080

WORKDIR /opt/time-service

COPY --from=builder /opt/time-service/time-service .

CMD ["/opt/time-service/time-service"]