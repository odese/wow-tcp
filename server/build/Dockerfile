FROM golang:1.18.3-buster AS build-env

RUN mkdir -p /go/src/wow
WORKDIR /go/src/wow

COPY . .

WORKDIR /go/src/wow/server

RUN go mod init
RUN go mod tidy

RUN go build -o server ./cmd/main.go

FROM debian:buster

EXPOSE 6090
CMD ["./server"]