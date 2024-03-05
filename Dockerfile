FROM golang:1.17 AS builder
WORKDIR /app

COPY . .

RUN go build -o app .

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/app .

CMD ["./app"]
