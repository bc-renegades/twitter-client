FROM golang:1.14-alpine

WORKDIR /app

COPY ./ /app

RUN go mod vendor

RUN go get github.com/pilu/fresh