FROM golang:1.13 AS builder

WORKDIR /build
ADD app.go .
ADD go.mod .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o simple .

FROM alpine:latest

WORKDIR /app

ADD simple .

CMD [ "/app/simple" ]
