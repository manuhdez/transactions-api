FROM golang:1.23-alpine as builder

RUN go install github.com/air-verse/air@latest
RUN go install github.com/pressly/goose/v3/cmd/goose@latest

WORKDIR /app

COPY ./internal/accounts /internal/accounts
COPY ./shared /shared
COPY ./internal/accounts/.air.toml /

ENTRYPOINT ["air"]
