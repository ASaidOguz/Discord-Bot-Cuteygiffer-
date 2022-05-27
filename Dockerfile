FROM golang:1.18-alpine

RUN mkdir /build

ADD go.mod go.sum main.go /build/

WORKDIR /build
RUN go build