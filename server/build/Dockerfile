# syntax=docker/dockerfile:1

FROM golang:1.18.3-buster
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o server /app/server/cmd/main.go
EXPOSE 6090
CMD [ "server" ]