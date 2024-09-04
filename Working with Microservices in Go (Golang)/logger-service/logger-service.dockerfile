# base go image
FROM golang:1.23-alpine as builder

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN CGO_ENABLED=0 go build -o brokerServiceApp ./cmd/api

RUN chmod +x /app/brokerServiceApp

# build a tiny docker image
FROM alpine:latest

RUN mkdir /app

COPY --from=builder /app/brokerServiceApp /app

CMD ["/app/brokerServiceApp"]

