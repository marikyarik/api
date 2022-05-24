# syntax=docker/dockerfile:1

FROM golang:latest

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . ./

RUN go build -o /api

EXPOSE 80

CMD [ "/api" ]