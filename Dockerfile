FROM golang:1.21.0 AS builder
WORKDIR /app

COPY . .

RUN go build -o nodeJob .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/nodeJob /app/nodeJob

CMD ["/app/nodeJob"]
