FROM golang:1.21.0 AS builder

WORKDIR /app

COPY . .

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux

RUN go build -ldflags "-s -w -extldflags '-static'" -o nodeJob

FROM alpine

COPY --from=builder /app/nodeJob /

CMD ["/nodeJob"]
