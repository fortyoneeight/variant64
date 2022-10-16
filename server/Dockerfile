# syntax=docker/dockerfile:1

FROM golang:1.19.2-alpine

WORKDIR /app

COPY go.mod ./

COPY go.sum ./

RUN go mod download

COPY *.go ./

COPY . ./

RUN go build -o /server

EXPOSE 8000

CMD [ "/server"]