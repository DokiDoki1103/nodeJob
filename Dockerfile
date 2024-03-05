FROM golang:1.21.0 AS builder
WORKDIR /app

COPY . .

RUN go build -o app .

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/app .

CMD ["./app"]
